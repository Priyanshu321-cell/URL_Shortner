package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

// in memory store
type Store struct {
	mu      sync.Mutex
	memory  map[int64]string
	counter int64
}

func NewStore(counter int64) *Store {
	return &Store{counter: 10000, memory: make(map[int64]string)}
}

func (s *Store) Add(longUrl string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	s.memory[s.counter] = longUrl
	return s.counter
}

func (s *Store) Get(id int64) (string, error) {
	if s.memory[id] == "" {
		return "", fmt.Errorf("Key not found")
	}

	return s.memory[id], nil
}

func main() {
	srv := &Server{store: NewStore(0)}

	addrPtr := flag.String("addr", ":4000", "a string address default=\"4000\"")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", srv.home)
	mux.HandleFunc("POST /longURL", srv.urlPost)
	mux.HandleFunc("GET /shortURL", srv.urlGet)

	log.Println("Starting server at address", *addrPtr)
	err := http.ListenAndServe(*addrPtr, mux)
	log.Println(err)
}

type Server struct {
	store *Store
}

func (srv *Server) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It works\n"))
}

type UrlData struct {
	LongUrl string `json:"LongUrl"`
}

const baseURL string = "http://tinyURL/"

func (srv *Server) urlPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var data UrlData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if data.LongUrl == "" {
		http.Error(w, "invalid JSON: ", http.StatusBadRequest)
		return
	}

	id := srv.store.Add(data.LongUrl)
	shortUrl, err := Id_to_base_62(id)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	shortUrl = baseURL + shortUrl
	w.Header().Add("Short-URL", shortUrl)
}

type ShortUrlData struct {
	ShortUrl string `json:"ShortUrl"`
}

func (srv *Server) urlGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data ShortUrlData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if data.ShortUrl == "" {
		http.Error(w, "invalid JSON: ", http.StatusBadRequest)
		return
	}

	str := data.ShortUrl
	str = strings.TrimPrefix(str, baseURL)

	if str == "" {
		http.Error(w, "invalid JSON: ", http.StatusBadRequest)
		return
	}

	id, err := Base_62_to_Id(str)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	longUrl, err := srv.store.Get(id)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Long-URL", longUrl)
}
