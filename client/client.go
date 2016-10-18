package main

import "net/http"
import "fmt"

func main() {
	resp, err := http.Get("http://127.0.0.1:9999/witticism")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}
