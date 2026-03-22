package main

import (
	"fmt"
	"time"
	"context"
)

func main () {
	fmt.Println ("Understanding context in golang \n")

	// Here we are telling the context to become unavailable after x second
	// this means, any task that runs withing this context should finish before x second if not, the context closes
	// and the task does not finish.
	ctx, cancel:= context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	go Task(context.WithValue(ctx, "hello", "world"))
	go Task(context.WithValue(ctx, "hello", "john"))

	<- ctx.Done()

	// go func () {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Task finished")
	// 	cancel ()
	// } ()

	// select {
	// case <- ctx.Done () :
	// 	fmt.Println("Context is done.")
	// 	err := ctx.Err()
	// 	if err != nil {
	// 		fmt.Printf("err : %s\n", err)
	// 	}

	// }
}


func Task (ctx context.Context) {
	var i = 0
	for {
		select {
			case <- ctx.Done():
				fmt.Println("Context is done.")
				return
			default:
				i++
				fmt.Printf("Running [%s] ...%d\n", ctx.Value("hello"), i)
				time.Sleep(500 * time.Millisecond)
		}
	}
}