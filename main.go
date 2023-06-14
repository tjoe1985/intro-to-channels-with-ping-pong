package main

import (
	"fmt"
	"strings"
)

func main() {
	//create a channel that takes a type string
	ping := make(chan string)
	//create a channel that takes a type string
	pong := make(chan string)
	go shout(ping, pong)

	//provide instructions to user
	fmt.Println("Type somethign and press ENTER (ENter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input

		var userInput string
		//ignoring two returns because we don't care right now
		_, _ = fmt.Scanln(&userInput)

		// check if the user entered Q or q and break if needed
		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput
		//wait for response
		response := <-pong

		fmt.Println("Response: ", response)

	}
	//clsing the channels
	fmt.Println("Closing channels after done")
	close(ping)
	close(pong)
}
func shout(ping <-chan string, pong chan<- string) {
	for {
		// get values from channel ping
		s := <-ping
		// white info to channel pong
		pong <- fmt.Sprintf("%s!!!!!", strings.ToUpper(s))

	}

}
