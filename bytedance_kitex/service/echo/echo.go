package echo

import "fmt"

type Echo struct {
	Req  string
	Resp string
}

func (e *Echo) Echo() string {
	fmt.Println("a")
	return "a"
}
func (e *Echo) ToString() string {
	return "b"
}
