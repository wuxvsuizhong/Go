package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 0)
	<-ctx.Done()
	fmt.Println("time out")
}
