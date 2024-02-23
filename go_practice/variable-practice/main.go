package main

import "fmt"
var w = 99
func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

		var g int
		fmt.Println(g)
		g = 43
		fmt.Println(g)

		var h int = 44
		fmt.Println(h)

		adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	// fmt.Printf("%v \t %b \t %#X\n",a,a,a)
	// fmt.Printf("%v \t %b \t %#X\n",b,b,b)
	// fmt.Printf("%v \t %b \t %#X\n",c,c,c)
	// fmt.Printf("%v \t %b \t %#X\n",d,d,d)
	// fmt.Printf("%v \t %b \t %#X\n",e,e,e)
	// fmt.Printf("%v \t %b \t %#X\n",f,f,f)


	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
	
}