package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var flag bool
	var isAwesome = true
	fmt.Println("HelloWorld!")
	fmt.Println(flag && isAwesome)

	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

	var x1 int = 10
	var b byte = 100
	var y1 int = int(b) + x1
	fmt.Println(y1)

	var x2, y2 int // both x2 and y2 are set to 0
	fmt.Println(x2, y2)

	// declaring many variable together
	var (
		x3   int
		y3       = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)
	fmt.Println(x3, y3, z, d, e, f, g)

	x4 := 30
	x5, y5 := 10, 20
	fmt.Println(x4, x5, y5)

	// re-assigment using var
	var x6 = 20
	x6 = 60

	// re-assignment without using var
	x7 := 20
	x7 = 30
	fmt.Println(x6, x7)

	// Not go compiler doesnt complain if the const are not utilized
	// since they are marked as immutable and go compiler wont have to worry
	// about them
	const (
		idKey   = "id"
		nameKey = "name"
	)

	const z1 = 20 * 10

	//fmt.Println(z1)
	// go doesnt provide a way to calcuate at runtime the value of an immutable

	/*
		constants can be typed or untyped and and can only to assigned to a varible of that type
	*/

	const x8 = 10

	var y8 int = x8

	fmt.Println(y8) // since y8 is int type it wont through error for const x8 int type

	y8 = 100 // note: since y8 is read already once, the compiler should not complain

	// constants in go is calculated at compile time, and hence its easy to eliminate if unused and will not b
	// part of the compiled binary and hence no errors

	// Exercies

	// #1
	var i = 20
	var f1 = float64(i) // implicit type casting of f1
	fmt.Println(f1)

	// # 2
	const value = 20
	var intVal = value
	var floatVal float32 = value

	// Note we need implicit or explicit casting for floatVal else it will default to int
	fmt.Println(intVal, floatVal, reflect.TypeOf(value), reflect.TypeOf(intVal), reflect.TypeOf(floatVal))

	var (
		bb     byte   = math.MaxUint8
		smallI int32  = math.MaxInt32
		bigI   uint64 = math.MaxUint64
	)

	fmt.Println(bb, smallI, bigI)

}
