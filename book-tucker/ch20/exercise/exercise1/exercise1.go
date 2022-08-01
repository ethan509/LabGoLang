package exercise1

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

// 위 둘을 합친다면 아마 이런 모습이었을 것.
// type Student struct implement Stringer {
// 	Name string
// 	Age  int
// }

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해.", s.Age, s.Name)
}

func Ex1() {
	fmt.Println("No1_1")

	student := Student{"철수", 12}

	var stringer Stringer = student

	fmt.Printf("%s\n", stringer.String())
}
