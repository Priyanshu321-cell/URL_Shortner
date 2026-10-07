package main

import (
	"fmt"
	"math/big"
	"slices"
)

func main() {
	num, _ := Id_to_base_62(11157)
	fmt.Println(num)
}

const base_62_map string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Id_to_base_62(val int64) (string, error) {
	var res []byte
	if val <= 0 {
		return string(res), nil
	}

	rem := big.NewInt(val)
	bigVal := big.NewInt(val)
	div := big.NewInt(62)
	zero := big.NewInt(0)
	for bigVal.Cmp(zero) != 0 {
		rem.Mod(bigVal, div)
		bigVal.Div(bigVal, div)
		res = append(res, base_62_map[rem.Int64()])
	}

	slices.Reverse(res)
	return string(res), nil
}

func Base_62_to_Id(val string) (int64, error) {
	if val == "" {
		return -1, nil
	}

	var b_val []byte
	b_val = []byte(val)

	slices.Reverse(b_val)

}
