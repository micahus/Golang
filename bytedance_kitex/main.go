package main

import (
	"fmt"

	"micah/echo/service/echo"
)

func main() {
	e := echo.Echo{}
	fmt.Println(e.Echo())
}
