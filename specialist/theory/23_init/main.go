package main

import (
	"23_init/rectangle"
	"fmt"
)

// init() - функция, срабатывающая единожды припервом импортировании пакета
// в пакета может быть несколько init()

func init() {
	fmt.Println("Init package MAIN")
}

func main() {
	r := rectangle.New(12, 15)
	fmt.Println(r)
}
