package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
	}

	//make a channel
	//channel monitors finish time of all routines
	//routines are spawned by go keyword to concurrently make http requests(no wait to resp)
	// chan string implies only string msgs can be sent across the channel

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//basically everytime a website is OK(returns resp), another ping is done, ping after ping after ping
	for l := range c {
		//ok so if u use this l and then another routine overwrites this l during exec u r screwed
		//so we pass that l(at the bottom) and store it seperately for the go func to use

		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Down")
		c <- link
		return
	}
	fmt.Println(link, "OK")
	c <- link
}
