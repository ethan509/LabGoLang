package exercise1

import "fmt"

func No1_1() {
	fmt.Println("No1_1")

	array := [3]int{1, 2, 3}
	slice := array[1:3]

	fmt.Printf("[BEFORE]array=%v, cap(%d)\n", array, cap(array)) // [BEFORE]array=[1 2 3], cap(3)
	fmt.Printf("[BEFORE]slice=%v, cap(%d)\n", slice, cap(slice)) // [BEFORE]slice=[2 3], cap(2)

	slice = append(slice, 100)

	fmt.Printf(" [AFTER]array=%v, cap(%d)\n", array, cap(array)) // [AFTER]array=[1 2 3], cap(3)
	fmt.Printf(" [AFTER]slice=%v, cap(%d)\n", slice, cap(slice)) // [AFTER]slice=[2 3 100], cap(4)
}

func No1_2() {
	fmt.Println("No1_2")

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]

	fmt.Printf("[BEFORE]array=%v, cap(%d)\n", array, cap(array)) // [BEFORE]array=[1 2 3 4 5], cap(5)
	fmt.Printf("[BEFORE]slice=%v, cap(%d)\n", slice, cap(slice)) // [BEFORE]slice=[2 3], cap(4)

	slice = append(slice, 100)

	fmt.Printf(" [AFTER]array=%v, cap(%d)\n", array, cap(array)) // [AFTER]array=[1 2 3 100 5], cap(5)
	fmt.Printf(" [AFTER]slice=%v, cap(%d)\n", slice, cap(slice)) // [AFTER]slice=[2 3 100], cap(4)
}
