package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	value := ctx.Value("key")
	fmt.Println("Value:", value)
}

func main() {
	ctx := context.WithValue(context.Background(), "key", "thread-specific data")
	go worker(ctx)
}
