package main

import (
	"context"
	"time"

	"github.com/fatih/color"
)

func deadline() {
	ctx := context.Background()
	deadline := time.Now().Add(20*time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)

	go func(){
		time.Sleep(10*time.Second)
		cancel()
	}()

	select{
	case <-ctx.Done():
		color.Red(ctx.Err().Error())
	case <-time.After(3*time.Second):
		color.Green("Time's UP!!")
	}
}