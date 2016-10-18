package main

import "net/http"
import "fmt"
import "io/ioutil"

func main() {
	count := 1
	messages := make(chan string, 2)
	for count < 5000 {
		go func(msg string) {
			resp, err := http.Get("http://127.0.0.1:9999/witticism")
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			messages <- string(body)
		}("getting")
		fmt.Println(count)
		msg := <-messages
		fmt.Println(msg)
		count++
	}
	close(messages)
}
