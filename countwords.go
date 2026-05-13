package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountWords(text string) map[string]int {

	words := strings.Fields(text)
	count:=make(map[string]int)

	for _, word:= range words{
		word=strings.ToLower(word)
		word= strings.TrimFunc(word, func (r rune)bool{
			return unicode.IsPunct(r)
		})
		count[word]++
	}
	return count
}

func main(){

	text:="go is good to learn and go is fast...."

	result:=CountWords(text)

	for word, count:= range result{

		fmt.Printf("%s:%d\n",word, count)
	}
}
