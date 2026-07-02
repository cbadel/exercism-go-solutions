package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	kelime:=strings.Fields(s)
	m:=make(map [string]int)
	for i:= 0 ;i<len(kelime);i++{
		_,ok:=m[kelime[i]]
		if ok== true {
			m[kelime[i]]+=1
		}else{
			m[kelime[i]]=1
		}
		
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
