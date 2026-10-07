## Why URL Shortner ?
As I have recently learned a bit golang I really wanted to build a project that could actually help me have understand backend development and can act as a foundation to expand my knowledge.<br>
This project might not be best or production ready as I will make mistakes solve them and then make more mistakes to learn more , I will force myself to avoid using AI rather will do the good old google search and reading articles and docmentation.<br>
This md file will contain articles I read , what I learn , and why I made a particular decision in this project.

## What is URL Shortner ?
So as the name suggest it just shortens your lengthy URLs by giving you a tiny version of them which act as a more polished links.
*But why would you need one ?**<br>
Ofcourse for the aesthetics ,aren't short polished clean urls are better and less cumbersome for online audience. Hence many business use them.You can read more online [here](https://buffer.com/social-media-terms/url-shortener).  

## Breaking down the URL Shortner 
Ok you got the idea what actually it is. So how will one actually design it.
Lets see **What we need to do ?** <br>
Hmm just the following -
1. User gives long url. -> *https://thisIsAVeryLongAndUglyURL.com/*
2. We give a short working url -> *https://tinyURL/a232* <br>
Sounds done we have to do this only conceptually. Lets get more technical now.

## Designing the URL Shortner
Ok we know what we want but lets actually think technically.<br>
So what all we need - 
1. A POST Request -> Here user provides long url and we convert it to a short url
2. A GET Request -> Whenever user hits that short url we redirect it 

## How to get a Short URL ?
So after all searching and stuff there are a few methods -
1. Hashing -> Using a hash function like MD5, SHA-256 etc that generate a fixed length hash of original url.
2. Base 62 encoding of sequential IDs -> Using auto incrementing integer ID and converting them to Base 62(0-9, a-z, A-Z) <br>
Both methods look good but have thir own problems

## Why I avoided Hashing ?
Ok i could have actually used hashing as it always gives a fixed length URL but it comes with need of collision resolution and with time it becomes expensive as number of collition increases. And I don't want to deal with this complexity at this small scale project.

## Base 62 encoding for sequential IDs
It is a common approach to use though it does not generate fixed length strings.<br>
Base 62 means using 62 different characters to represent a number . These include : *0-9, a-z, A-Z*<br>
Example -> 243534(base 10) = 11IY after getting remainder of each division

## Why Base 62 ?
I choose it because of simplicity on cost of not getting fixed length urls .

