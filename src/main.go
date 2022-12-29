package main

import "fmt"

func main() {
	// Declarative variables
	const itsForever int = 8
	const PI = 3.14

	fmt.Println("PI + twelve", PI+12)
	fmt.Println(itsForever, PI)

	base := 12 // no ha sido declarada anteriormente sino que se crea ahora :=
	fmt.Println(base)

	var height int = 14
	var area int = 23

	fmt.Println(height, area)

	// Zero value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// area al cuadrado
	const baseCuadrada = 10
	areaCuadrada := baseCuadrada * baseCuadrada
	fmt.Println("Area:", areaCuadrada)
}
