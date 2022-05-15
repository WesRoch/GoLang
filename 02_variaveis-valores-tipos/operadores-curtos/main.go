package main

import "fmt"

func main() {

	x := 10 // operador curto de declaração - irá declarar pelo menos uma variável
	z := 10.0
	y := "Olá, bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20 //operador de atribuição (atribuindo um valor a uma variável já existente)

	fmt.Printf("x: %v, %T\n", x, x)
}

// keywords -> palavras reservadas = nao podem ser usadas para atribuir, definir algo, como por ex: func, package, import

// x := 10 + 10 -> 10 é o operando e + é o operador
