package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {
	resp, body, errs := gorequest.New().Get("http://43.142.28.232:8888/api/ping").End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(resp)
	fmt.Println(body)
	// fmt.Println(resp, body)
}
