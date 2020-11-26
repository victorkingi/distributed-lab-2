package main

import (
	"bufio"
	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"
	"secretstrings/stubs"

	//	"bufio"
	//	"os"
	//	"secretstrings/stubs"
	"fmt"
)

func main() {
	server := flag.String("server", "18.212.75.40:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	f, _ := os.Open("wordlist")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		request := stubs.Request{Message: line}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded: " + response.Message)
	}
}
