package main

import (
	"earth/pkg/stream"
	"log"
	"os"
	"os/signal"
	"syscall"

	gops "github.com/google/gops/agent"
)

func updateCallback(rev string) {
	log.Println("New revision detected:", rev)
}

func main() {
	// watcher, _ := watcheretcd.NewWatcher([]string{"http://127.0.0.1:49155"}, "/agent/node/127.0.0.1")
	// _ = watcher.SetUpdateCallback(updateCallback)

	// select {}

	client()

}

func client() {
	// Gops (for detecting goroutine leaks)
	if err := gops.Listen(gops.Options{}); err != nil {
		log.Fatal(err)
	}

	// Init the client
	client := &stream.RPCClient{}
	client.Init()
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Continue indefinitely
	//select {}

	// Listen to the USR1 signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGUSR1,
	)

	for {
		// Wait for the signal to disconnect
		<-sigCh
		err = client.Disconnect()
		if err != nil {
			panic(err)
		}

		// On the next USR1 signal, re-connect
		<-sigCh
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}
}
