package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mvdan/xurls"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	baseURL := "https://en.wikipedia.org/wiki/National_Basketball_Association"
	for i := 0; i < 100; i++ {
		response, err := http.Get(baseURL)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		//fmt.Println(response.Body)
		body, err := ioutil.ReadAll(response.Body)
		// fmt.Println(string(body))
		links := xurls.Strict.FindAllString(string(body), -1)
		fmt.Println(len(links))
		fmt.Println(links[0])
		response.Body.Close()
		baseURL = links[0]
	}
}
