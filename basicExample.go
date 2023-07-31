package main

import (
	"context"
	"github.com/fatih/color"
	"time"
)


// This simple program starts with a context.Background() and derives a new context from it using context.WithCancel(). 
// This derived context is passed to a goroutine which sleeps for x seconds and then cancels the context. 
// The select block waits until either the context is done (which happens when cancel() is called), 
// or a time limit of y seconds is reached. The output will be "Context has been cancelled" as the context gets cancelled before the timeout.
func example1() {
	ctx := context.Background()         						// get a context
	ctx, cancel := context.WithCancel(ctx)						// derive a new context from it

	go func(){
		time.Sleep(3*time.Second)
		cancel()                                             	// cancel the context
	}()

	select{
	case <-ctx.Done():
		color.Red("Context has been cancelled!")
	case <-time.After(5 * time.Second):
		color.Green("Time's UP!!")
	}

}