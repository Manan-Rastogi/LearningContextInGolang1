package main

import (
	"context"
	"time"

	"github.com/fatih/color"
)

func timeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)

	go func(){
		time.Sleep(12 * time.Second)
		cancel()
	}()

	select{
	case <-ctx.Done():
		color.Red(ctx.Err().Error())
	case <-time.After(10*time.Second):
		color.Green("Time's UP!!")
	}
}