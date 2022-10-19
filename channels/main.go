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
		"http://golang.org",
		"http://amazon.in",
	}

	c := make(chan string)
	
	for _, link := range links {
		go checkLink(link, c)
	}
	// for i:= 1; i <= len(links); i++ {
	// 	fmt.Println(<-c)
	// }	

	// for looping infinitely
	// for{
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		// go checkLink(l, c)
		go func(link string) {                   //function literal
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		// c <- "might be down"
		c <- link
		return
	}

	fmt.Println(link, "is up")
	// c <- "yep it is"
	c <- link
}
