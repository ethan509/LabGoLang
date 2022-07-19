package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array: ", array)                                 // [1 2 3 4 5]
	fmt.Println("slice: ", slice)                                 // [2]
	fmt.Printf("slice len(%d), cap(%d) ", len(slice), cap(slice)) // 1, 4
	fmt.Println()

	array[1] = 100

	fmt.Println()
	fmt.Println("array: ", array) // [1 100 3 4 5]
	fmt.Println("slice: ", slice) // [100]

	slice = append(slice, 500)

	fmt.Println()
	fmt.Println("array: ", array)                                 // [1 100 500 4 5]
	fmt.Println("slice: ", slice)                                 // [100, 500]
	fmt.Printf("slice len(%d), cap(%d) ", len(slice), cap(slice)) // 2, 4

	// 왜 array의 값이 500으로 바뀌었을까?
	// line 7(slice := array[1:2]) 에서 slicing할 때 cap 길이는 [1]부터 배열의 마지막 index까지 사용할 수 있게 되면서 cap이 4가 된다.
	// 그렇기 때문에 append를 사용하면 새로운 배열을 만들지 않고 index 2의 값을 변경하게 된다.

	// 만약 slice에 빈공간이 없었다면~?
	whatIf1()

	whatIf2()

	// slice ->(copy) -> append() 사용해서 slice 복사하는
	copySlice1()
	// slice ->(copy) -> copy() 사용해서 slice 복사하는
	copySlice2()

	copySlice3()

	deleteSlice()

	appendSlice()

	appendSlice2()

	sortSlice()
}

func whatIf1() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered whatIf1()")

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[4:5]

	fmt.Println("array: ", array)                                 // [1 2 3 4 5]
	fmt.Println("slice: ", slice)                                 // [4]
	fmt.Printf("slice len(%d), cap(%d) ", len(slice), cap(slice)) // len(1), cap(1)
	fmt.Println()

	array[4] = 100

	fmt.Println()
	fmt.Println("array: ", array) // [1 2 3 4 100]
	fmt.Println("slice: ", slice) // [100]

	slice = append(slice, 500)

	fmt.Println()
	fmt.Println("array: ", array)                                 // [1 2 3 4 100]
	fmt.Println("slice: ", slice)                                 // [100 500]
	fmt.Printf("slice len(%d), cap(%d) ", len(slice), cap(slice)) // len(2), cap(2)
}

func whatIf2() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered whatIf2()")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[0:3] // 이것과 동일한 선언 -> slice2 := slice1[:3]

	fmt.Println("slice1: ", slice1)                                   // [1 2 3 4 5]
	fmt.Println("slice2: ", slice2)                                   // [1 2 3]
	fmt.Printf("slice1 len(%d), cap(%d)\n", len(slice1), cap(slice1)) // len(5), cap(5)
	fmt.Printf("slice2 len(%d), cap(%d)\n", len(slice2), cap(slice2)) // len(3), cap(5)

}

func copySlice1() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered copySlice1()")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append([]int{}, slice1...) // 슬라이스에서 다른 슬라이스로 복사하는 방법.

	fmt.Println("slice1: ", slice1)                                   // [1 2 3 4 5]
	fmt.Println("slice2: ", slice2)                                   // [1 2 3 4 5]
	fmt.Printf("slice1 len(%d), cap(%d)\n", len(slice1), cap(slice1)) // len(5), cap(5)
	fmt.Printf("slice2 len(%d), cap(%d)\n", len(slice2), cap(slice2)) // len(5), cap(6)
	// 그런데 같지 않음! cap이 1개 더 많다.
	// 완전히 같게 하려면 copySlice3() 처럼 해야 함.
}

func copySlice2() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered copySlice2()")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)

	fmt.Println("slice1: ", slice1) // [1 2 3 4 5]
	fmt.Println("slice2: ", slice2) // [1 2 3 4 5]
	fmt.Println("slice3: ", slice3) // [1 2 3 4 5]
	fmt.Printf("slice1:%v len(%d), cap(%d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2:%v len(%d), cap(%d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3:%v len(%d), cap(%d)\n", slice3, len(slice3), cap(slice3))

	// The copy built-in function copies elements from a source slice into a
	// destination slice. (As a special case, it also will copy bytes from a
	// string to a slice of bytes.) The source and destination may overlap. Copy
	// returns the number of elements copied, which will be the minimum of
	// len(src) and len(dst).
	// PROTO-TYPE
	//		func copy(dst, src []Type) int
	cnt1 := copy(slice2, slice1) // !
	cnt2 := copy(slice3, slice1) // !

	fmt.Println(cnt1)
	fmt.Println(cnt2)

	fmt.Printf("slice1:%v len(%d), cap(%d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2:%v len(%d), cap(%d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3:%v len(%d), cap(%d)\n", slice3, len(slice3), cap(slice3))
}

func copySlice3() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered copySlice3()")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	fmt.Printf("slice1:%v len(%d), cap(%d)\n", slice1, len(slice1), cap(slice1)) // slice1:[1 2 3 4 5] len(5), cap(5)
	fmt.Printf("slice2:%v len(%d), cap(%d)\n", slice2, len(slice2), cap(slice2)) // slice2:[1 2 3 4 5] len(5), cap(5)
}

func deleteSlice() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered deleteSlice()")

	idx := 2
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("[BEFORE] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) // [BEFORE] slice:[1 2 3 4 5 6 7 8 9 10] len(10), cap(10)
	slice = append(slice[:idx], slice[idx+1:]...)

	// 그런데.... cap은 여전히 10.
	fmt.Printf(" [AFTER] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) //  [AFTER] slice:[1 2 4 5 6 7 8 9 10] len(9), cap(10)
}

func appendSlice() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered appendSlice()")

	idx := 2
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("[BEFORE] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) // [BEFORE] slice:[1 2 3 4 5 6 7 8 9 10] len(10), cap(10)
	slice = append(slice[:idx+1], slice[idx:]...)

	slice[idx] = 100
	// 그렇지만 cap이 10 에서 20으로 두배가 되었다.
	fmt.Printf(" [AFTER] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) //   [AFTER] slice:[1 2 100 3 4 5 6 7 8 9 10] len(11), cap(20)
}

func appendSlice2() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered appendSlice2()")

	idx := 2
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	slice := make([]int, 10)

	fmt.Println(unsafe.Sizeof(idx))
	fmt.Println(unsafe.Sizeof(slice))

	//fmt.Printf("Page Size: %d\n", os.Getpagesize())

	fmt.Printf("[BEFORE] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) // [BEFORE] slice:[1 2 3 4 5 6 7 8 9 10] len(10), cap(10)
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])

	slice[idx] = 100
	// 그렇지만 cap이 10 에서 20으로 두배가 되었다.
	fmt.Printf(" [AFTER] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) //   [AFTER] slice:[1 2 100 3 4 5 6 7 8 9 10] len(11), cap(20)

	fmt.Println(unsafe.Sizeof(slice))

	//fmt.Printf("Page Size: %d\n", os.Getpagesize())
}

func sortSlice() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Entered sortSlice()")

	slice := []int{12, 2, 9, 4, 5, 8, 7, 6, 3, 10, 11, 1, 13}

	fmt.Printf("[BEFORE] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) // [BEFORE] slice:[12 2 9 4 5 8 7 6 3 10 11 1 13] len(13), cap(13)

	sort.Ints(slice)
	fmt.Printf(" [AFTER] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) //  [AFTER] slice:[1 2 3 4 5 6 7 8 9 10 11 12 13] len(13), cap(13)

	n := sort.SearchInts(slice, 11)
	fmt.Printf(" [AFTER] slice:%v len(%d), cap(%d)\n", slice, len(slice), cap(slice)) //  [AFTER] slice:[1 2 3 4 5 6 7 8 9 10 11 12 13] len(13), cap(13)
	fmt.Println(n)
}
