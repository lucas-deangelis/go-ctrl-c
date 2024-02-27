package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// We need a root context that we can cancel, to derive all other
	// contexts from, so that every context in the application is
	// cancelled when we press ctrl+c.
	ctx, cancel := context.WithCancel(context.Background())

	// This channel will allow us to transmit the signal to the
	// goroutine handling ctrl+c.
	c := make(chan os.Signal, 1)

	// We relay ctrl+c events (`os.Interrupt` in Go) to the channel.
	signal.Notify(c, os.Interrupt)

	// We start a goroutine that will cancel the root context when it
	// receives a signal from `c`.
	// In case it receives a second signal, it'll force exit.
	go func() {
		<-c
		fmt.Println("\nctrl+c pressed, canceling context")
		cancel()

		<-c
		fmt.Println("\nctrl+c pressed again, force exit")
		os.Exit(1)
	}()

	doStuff(ctx)
}

// doStuff simulates some work being done.
// The 1 second delay allows us to see that when receiving ctrl+c, it'll
// finish its current work (due to the `select` statement in the `for` loop),
// before cleaning up.
// If we press ctrl+c again, either during the work or during the cleanup, it'll force exit.
func doStuff(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("doStuff: context done, cleaning up")
			time.Sleep(3 * time.Second)
			fmt.Println("doStuff: cleanup done")
			return
		default:
			fmt.Println("starting to work")
			time.Sleep(1 * time.Second)
			fmt.Println("finishing work")
		}
	}
}
