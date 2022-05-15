package main

import "fmt"

var x int = 10 //Quando declarar uma variavel, o tipo da vaiavel será estático, ou seja, no caso, será int sempre.

var y = 10.5 //atribuindo automaticamente que a variavel Y é do tipo float, mostrado no block do codigo

var z int //indicando que a variavel Z é do tipo INT

func main() {

	z = 20 //tipo int declarado no package

	fmt.Printf("%v, %T\n\n", x, x)
	fmt.Printf("%v, %T\n\n", y, y)
	fmt.Printf("%v, %T", z, z)
}
