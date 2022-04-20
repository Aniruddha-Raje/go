package main

import "fmt"

func main() {
	var a = 20
	var b = "Abc"
	var c = 34.80

	fmt.Printf("a = %d of type = %T \n", a, a)
    fmt.Printf("b = %s of type = %T \n", b, b)
	fmt.Printf("c = %f of type = %T \n", c, c)

	var num int
    var str string
    var flt float64

	fmt.Printf("\na = %d \nb = %s \nc = %f\n", num, str, flt)

	var x, y, z int = 1, 2, 3
	fmt.Printf("\nx = %d \ny = %d \nz = %d\n", x, y, z)

	p := 10
	q := "PQR"
	r := 10.50
	fmt.Printf("\np = %d of type = %T \n", p, p)
    fmt.Printf("q = %s of type = %T \n", q, q)
	fmt.Printf("r = %f of type = %T \n", r, r)
}