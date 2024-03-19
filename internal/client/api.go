package main

// import (
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"time"

// 	"github.com/kirinlabs/HttpRequest"
// )

// var transport *http.Transport

// func init() {
// 	transport = &http.Transport{
// 		DialContext: (&net.Dialer{
// 			Timeout:   30 * time.Second,
// 			KeepAlive: 30 * time.Second,
// 			DualStack: true,
// 		}).DialContext,
// 		MaxIdleConns:          100,
// 		IdleConnTimeout:       90 * time.Second,
// 		TLSHandshakeTimeout:   5 * time.Second,
// 		ExpectContinueTimeout: 1 * time.Second,
// 	}
// }

// func main() {
// 	req := HttpRequest.NewRequest()
// 	resp, err := req.Get("http://43.142.28.232:8888/api/ping")
// 	defer resp.Close()
// 	body, err := resp.Body()
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(string(body))
// 	return
// }
