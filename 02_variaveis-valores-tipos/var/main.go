package main

import "fmt"

var y = 10 //declarar a variavel no package e fora do codeblock usando o var

func main() {

	qualquercoisa(y)

}

func qualquercoisa(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
