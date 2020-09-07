package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://twitter.com",
		"http://instagram.com",
	}

	// 1) Makes a channel
	c := make(chan string)

	for _, link := range links {
		// 2) Send links into the channel
		go checkLink(link, c)
	}

	// infinite loop syntax
	// for {}
	// Alternale syntax for infinite loop with a channel
	for l := range c {
		// 3) Wait until something is received from the channel
		// and send it back into the channel
		go checkLink(l, c)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "- might be down!")
		// Send message into the channel
		c <- link
		return
	}
	fmt.Println(link, "- is up!")
	// Send message into the channel
	c <- link
}
