package main

import (
	"fmt"
	"net/http"
	"time"
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

		// function literal syntax (Anonymous function, lambda, etc...)
		// When creating a child goroutine, you need to pass in the
		// variable to the function literal otherwise you will be
		// operating on an unexpected memory location.
		go func(link string) {
			// Putting a call to time.Sleep in a function literal ensures
			// that the main goroutine does not block while waiting for
			// the timer to expire.
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)

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
