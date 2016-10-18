package main

import "net/http"
import "fmt"
import "math/rand"
import "log"

func witHandler(w http.ResponseWriter, r *http.Request) {
	var witticism [5]string
	witticism[0] = "He has all the virtues I dislike and none of the vices I admire."
	witticism[1] = "There is no such thing as public opinion. There is only published opinion."
	witticism[2] = "Without courage all virtues lose their meaning."
	witticism[3] = "A fanatic is one who can't change his mind and won't change the subject."
	witticism[4] = "The first quality that is needed is audacity."

	fmt.Fprintf(w, "%q", witticism[rand.Intn(5)])
}

func main() {
	http.HandleFunc("/witticism", witHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
