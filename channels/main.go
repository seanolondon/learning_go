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
		"http://amazon.com",
	}
	//c can be passed around the application
	//but only inside this function
	c := make(chan string)

	//ignore the index, i
	//create the intial websites to check, it creates then overwrites the place in memory
	for _, link := range links {
		// go... can only be used infront of function calls
		go checkLink(link, c)
	}

	//print the reult of the channel
	//google is the fastest website so it will be the first to respond
	//waiting on the first response if only a fmt.Print, recieving messages from a channel
	//is a BLOCKER.
	//will print just google
	//fmt.Println(<-c)
	// will print stack overflow
	//fmt.Println(<-c)

	// use a C style for looop

	//for i := 0; i < len(links); i++ {
	//for {
	//fmt.Println(<-c)
	//	go checkLink(<-c, c)
	//}
	// wait for this to return a channel

	//don't share a variable to a child, use a channel to do that
	for l := range c {
		//replace checkLink with a FUNCTION LITERAL (ANNONYMOUS FUNCTION)
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l) //extra parentheses are for a function literal () means call the function
	} //pass l into the literal function
	//don't reference the same variable in two separate routines, rely on pass by value
}

//channels are the go betweens between the different routinges, the main routine
// and all the other child routines. The channel must all be of the same type, eg. string
//int, etc.

//note from lecture: concurrency is not parallelism:
// concurrent if it has the ability to load up multiple go routines at the same time
//----could still be doing this on one core. Can schedule work to be done throughout each other
//Parallelism - only exists when there are multiple CPU cores
//can do the exact same thing and the same nanosecond.
//other core working at the same exact time.
//Concurrent - schedule work and change between them at the same time i.e. if one is blocked
//send the next thing through.
//Main reoutine vs child routines - main routine is the big/main/respected routine
//Child routines do not have respect yet

func checkLink(link string, c chan string) {
	//don't care about a response
	//only care about an error since this is an error test program
	// _ means an object that you don't care about the response
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		//send to channel
		// c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up")
	//send to channel
	// c <- "Yup, it's up"
	c <- link
}
