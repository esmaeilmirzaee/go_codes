package main

import (
	"fmt"
	"net/http"
)

type link string

func main() {
	links := []link{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan link)

	for _, link := range links {
		go link.checkLink(c)
	}

	occr := [4]int{0}

	for l := range c {

		switch l {
		case "http://google.com":
			occr[0]++
			fmt.Println("ggl-", occr[0])
			break
		case "http://golang.org":
			occr[1]++
			fmt.Println("go-", occr[1])
		case "http://facebook.com":
			occr[2]++
			fmt.Println("fb-", occr[2])
		case "http://stackoverflow.com":
			occr[3]++
			fmt.Println("so-", occr[3])
		default:
			break
		}
		go l.checkLink(c)
	}
}

func (l link) checkLink(c chan link) {
	_, err := http.Get(string(l))
	if err != nil {
		fmt.Println("Problem--", l)
		c <- l
		return
	}
	fmt.Println(l, " is OK.")
	c <- l
}
