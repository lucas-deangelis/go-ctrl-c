package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// With signal.NotifyContext, we can create a context that is cancelled
	// when we receive a signal. However, we can't handle the double ctrl+c
	// case as with the other example.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

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
