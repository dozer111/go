package rectangle

import "fmt"

func init() {
	fmt.Println("Init package Rectangle")
}

type Rectangle struct {
	A, B int
}

func New(newA, newB int) *Rectangle {
	return &Rectangle{newA, newB}
}

func main() {
	fmt.Println("Rectangle MAIN")
}
