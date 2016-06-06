package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	s := "/v1/packages?com=papa-http&ver=1.2.3"
	u, _ := url.Parse(s)
	fmt.Println(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)

	fmt.Println(m["com"])

	str := m["ver"][0]

	fmt.Println(str)

	t := m["test"]
	if t != nil {
		fmt.Println(t)
	} else {
		fmt.Println("test is not exist")
	}

	var st []string
	if st == nil {
		fmt.Println("st == nil")
	} else {
		fmt.Println("st != nil")
	}
}
