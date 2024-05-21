package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.com",
	}

	// Creating a channel
	c := make(chan string)

	for _, link := range links {
		// go keyword is used to create a new go routine
		go checkLink(link, c)
	}

	// Receiving data from the channel
	// Receiving data from a channel is a blocking operation
	// fmt.Println(<-c)

	// Receving data from the channel for all the routines
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Receiving the link from the channel and checking the link again
	// This will run infinitely and constantly check the links
	// for {
	// 	go checkLink(<-c, c)
	// }

	// This also works the same as the above code
	// This waits for the channel to return a value and then calls the checkLink function
	for l := range c {
		go checkLink(l, c)
	
	}
	
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Sending data to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Sending data to the channel
	c <- link
}	