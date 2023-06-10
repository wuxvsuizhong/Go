package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5)
	<-ctx.Done()
	fmt.Println("time out", ctx)
}
