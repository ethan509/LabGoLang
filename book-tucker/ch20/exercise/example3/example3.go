package example3

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	//r := Stringer.(Reader)
	if r, ok := stringer.(Reader); ok {
		r.Read()
	}
}

func Ex3() {

}
