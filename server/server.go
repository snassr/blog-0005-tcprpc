package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"reflect"
	"time"

	"github.com/snassr/blog-0005-tcprpc/egrpc"
)

var (
	// interrupt channel to receive interrupts from the operating system.
	interrupt = make(chan os.Signal, 1)
	// shutdown channel to receive shutdown signals from the application.
	shutdown = make(chan struct{})
)

func main() {
	// send interrupt signals on the interrupt channel
	signal.Notify(interrupt, os.Interrupt)

	// create person object
	person := &egrpc.Person{}

	// register person object as RPC (interface by the name `Person`)
	rpc.Register(person)

	// service address of server
	service := ":1200"

	// create tcp address
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Fatal(err)
	}

	// tcp network listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// run a shutdown countdown to notify the server it's time to end the process.
	go func() {
		shutdownTime := 10
		for i := 0; i < shutdownTime; i++ {
			time.Sleep(1 * time.Second)
			log.Printf("shutdown timer (%v)...%v", shutdownTime, i+1)
		}
		close(shutdown)
	}()

	// forever loop to listen for and accept client connections.
	for {
		// check if either of our channels have received a message
		// and act appropriately.
		select {
		case <-shutdown:
			// received shutdown signal
			log.Printf("shutdown request, rpc server ended.")
			// end server process successfully
			os.Exit(0)
		case <-interrupt:
			// received interrupt
			log.Printf("interrupted, rpc server ended.")
			// end server process successfully
			os.Exit(0)
		default:
			log.Printf("tcp listening...")
		}

		// wait a specific amount of time for a connection
		listener.SetDeadline(time.Now().Add(1 * time.Second))
		// handle connection if any, if not handle in error
		conn, err := listener.Accept()
		if err != nil {
			// check if timeout error
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			} else {
				// do something with bad errors
				log.Printf("connection error: %v", err)
				// end server process, unsucessfully
				os.Exit(1)
			}
		} else {
			// print connection info
			log.Printf("received rpc message %v %v", reflect.TypeOf(conn), conn)
			// handle listener (client) connections via rpc
			// using a goroutine (to handle more than one connection at a time)
			go rpc.ServeConn(conn)
		}
	}
}
