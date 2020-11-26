package main

import (
	//	"net/rpc"
	//	"fmt"
	//	"time"
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"time"
)

var nextAddr string
var initialised = false
var nextround *rpc.Client

type BottlesOfBeer struct{}
type Token struct {
	Bottles int
}

func Beers(i int) {
	time.Sleep(2 * time.Second)
	if i > 1 {
		fmt.Printf("%v bottles of beer on the wall, %v bottles of beer\nTake one down, pass it around...", i, i)
	} else if i == 1 {
		fmt.Printf("%v bottle of beer on the wall, %v bottle of beer\nTake one down, pass it around...", i, i)
	} else {
		fmt.Printf("No more bottles of beer on the wall, no more bottles of beer\nGo to the store and buy some more!\n")
	}
}

func PassItAround(bottles int) {
	reqs := Token{Bottles: bottles}
	resp := new(Token)

	if initialised == false {
		nextround, _ = rpc.Dial("tcp", nextAddr)
		initialised = true
	}
	nextround.Go("BottlesOfBeer.Round", reqs, resp, nil)
}

func (b *BottlesOfBeer) Round(inToken Token) (err error) {
	bottles := inToken.Bottles
	Beers(bottles)

	if bottles > 0 {
		PassItAround(bottles - 1)
	}

	return
}

func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
	rpc.Register(&BottlesOfBeer{})
	listener, _ := net.Listen("tcp", ":"+*thisPort)
	defer listener.Close()

	if *bottles > 0 {
		Beers(*bottles)
		go PassItAround(*bottles - 1)
	}
	rpc.Accept(listener)
}
