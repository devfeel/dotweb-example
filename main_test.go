package main

import (
	"testing"
	"strings"
	"fmt"
)

func TestHostParse(t *testing.T){
	host := "www.dev.com:2"
	if strings.Index(host, ":") >0{
		fmt.Println(strings.Split(host, ":")[0])
		fmt.Println(host[0:strings.Index(host, ":")])
	}else{
		fmt.Println(host)
	}
}

func BenchmarkTestHostParse1(b *testing.B) {
	host := "www.dev.com:2"
	for i := 0; i < b.N; i++ {
		_ = strings.Index(host, ":")
	}
}

func BenchmarkTestHostParse2(b *testing.B) {
	host := "www.dev.com:2"
	for i := 0; i < b.N; i++ {
		_ = strings.Split(host, ":")[0]
	}
}

func BenchmarkTestHostParse3(b *testing.B) {
	host := "www.dev.com:2"
	for i := 0; i < b.N; i++ {
		s := host[0:strings.Index(host, ":")]
		_ = s
	}
}
