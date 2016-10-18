package main

import "net/http"
import "fmt"
import "io/ioutil"

func main() {
	resp, err := http.Get("http://127.0.0.1:9999/witticism")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
