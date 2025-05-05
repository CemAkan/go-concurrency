/*
I got the idea from the ping-pong example in "Google I/O 2013 - Advanced Go Concurrency Patterns": https://www.youtube.com/watch?v=QDDwwePbDtw

I turned it into a 3-player bomb-passing game to show how unbuffered channels block until the receiver is ready.
The bomb goes around until it explodes. Simple and fun way to understand Go concurrency.
*/

package main

import (
	"fmt"
	"time"
	"math/rand"
)

const totalTime = 7

type Bomb struct{
	time int
	holder string
}

func Player(name string, in <-chan Bomb, out chan<- Bomb){
	for{
		bomb :=	 <- in //getting bomb from channel

		delay := rand.Intn(3) + 1 //it generate 0..2 but with +1 our outcome will be 1..3

		time.Sleep(time.Second * time.Duration(delay))

		bomb.time -= delay //decrease total time
		bomb.holder = name // change to holder name

		fmt.Printf("%s holded the bomb for %d seconds \n", name, delay ) //noticing who is holding

		if bomb.time <= 0 {
			fmt.Println("BOOOOMMMMMM -- Bomb is exploded on hands of " + name + " -- GAME OVER :)")
			return
		}
	
		out <- bomb // giving back the bomb
	}

}

func main(){

	rand.Seed(time.Now().UnixNano())
	
	// creating unbuffered channels
	a := make(chan Bomb)
	b := make(chan Bomb)
	c := make(chan Bomb)

	//starting Player gorutines
	go Player("Cem",a,b) // Cem -> Mete
	go Player("Mete",b,c) // Mete -> Melis
	go Player("Melis",c,a) // Melis -> Cem

	start := Bomb{time:totalTime, holder:"START"}

	a <- start //bomb is send to a channel for starting to cycle event

	time.Sleep(time.Second * time.Duration(totalTime + 2))

}


