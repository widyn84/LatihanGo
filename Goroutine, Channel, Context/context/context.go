package main

import (
	"context"
	"fmt"
)

func main() {
	contextParent := context.Background()
	ctx1 := context.WithValue(contextParent, "key1", "Hello World")
	ctx2 := context.WithValue(ctx1, "key2", "Hello Girls")
	ctx3 := context.WithValue(ctx2, "key3", "Hello Boys")
	ctx4 := context.WithValue(contextParent, "key4", "Hello Children")
	ctx5 := context.WithValue(contextParent, "key5", "Hello Later")
	fmt.Println(ctx5.Value("key5"))
	fmt.Println(ctx5.Value("key4"))
	fmt.Println(ctx5.Value("key3"))
	fmt.Println(ctx5.Value("key2"))
	fmt.Println(ctx5.Value("key1"))
	fmt.Println(ctx3.Value("key1"))
	fmt.Println(ctx4.Value("key1"))
}
