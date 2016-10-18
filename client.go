package client

import "net/http"
import "fmt"

func main() {
	resp, err := http.Get("http://goserver.dev/")
	if err != nil {
		// handle error
	}
	fmt.Println(resp)
}
