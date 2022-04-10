package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 20
	fmt.Println(v)

	// Pointers to struct
	pv := &v
	pv.X = 30 // in C this would be pv->X = 30, i.e  (*pv).X
	fmt.Println(v)

	// Struct literals
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X: 0 and Y: 0 is implicit
	p := &Vertex{1, 2} // has type *Vertex

	fmt.Println(v2, v3, p)

}
