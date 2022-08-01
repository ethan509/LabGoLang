package example08

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Ex08() {
	fmt.Println("Ex08")
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)

	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	fmt.Println("square")

	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d\n", n*n)
	}
	wg.Done()
}
