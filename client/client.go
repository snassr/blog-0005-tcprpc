package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/snassr/blog-0005-tcprpc/egrpc"
)

func main() {
	// get sample person object array data
	ppl := egrpc.SampleData

	// connect to server via rpc tcp
	client, err := rpc.Dial("tcp", ":1200")
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	// reply is used to hold responses
	var reply string

	// RPC call (1)
	// remote call procedure Person.FirstLast (FirstLast method on Person object)
	err = client.Call("Person.FirstLast", ppl[0], &reply)
	if err != nil {
		log.Fatal("Person.FirstLast error: ", err)
	}
	fmt.Printf("Person.FirstLast: %v\n", reply)
	// sleep for 2 seconds before next RPC (not required, just for demo)
	time.Sleep(2 * time.Second)

	// RPC call (2)
	// remote call procedure Person.LastFirst (LastFirst method on Person object)
	var bPerson egrpc.PersonData
	err = client.Call("Person.LastFirst", ppl[1], &bPerson)
	if err != nil {
		log.Fatal("Person.LastFirst error: ", err)
	}
	fmt.Printf("Person.LastFirst: %v\n", bPerson.Data)
	// sleep for 2 seconds before next RPC (not required, just for demo)
	time.Sleep(2 * time.Second)

	// RPC call (3)
	// remote call procedure Person.Bio (Bio method on Person object)
	err = client.Call("Person.Bio", ppl[2], &reply)
	if err != nil {
		log.Fatal("Person.Bio error: ", err)
	}
	fmt.Printf("Person.Bio:\n%v", reply)
	// sleep for 2 seconds before next RPC (not required, just for demo)
	time.Sleep(2 * time.Second)
}
