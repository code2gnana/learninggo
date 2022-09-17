package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomething(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("Doing domething cool")
		}
		time.Sleep((500 * time.Millisecond))
	}
}

func main() {
	fmt.Println("Go Context Tutorials")
	// Declaring a default context
	//ctx := context.Background()
	// Declaring a context with Deadline
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Oh no, I have exceeded the dead line")
		fmt.Println(ctx.Err())
	}
	doSomething(ctx)
}
