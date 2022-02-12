package main

import "fmt"

//https://golang.gr/2022/02/11/buffered-channels/
func main() {
	//make channel with 6 buffers
	ch := make(chan int, 6)
	//spin a goroutine by which you send
	//numbers 0-5 into the channel
	go func() {
		//when go func finishes close channel
		//in order to break the range
		defer close(ch)

		//send numbers 0-5 into the channel
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending... %d\n", i)
			ch <- i
		}
		fmt.Println("--------------")
	}()

	//receive the contents of the buffered channel
	for v := range ch {
		fmt.Printf("Receiving...%d\n", v)
	}
}

//next:channel-direction
