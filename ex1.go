//Faça um programa que utilize closures para criar uma
//função que gera (generator) os múltiplos de um número.
package main

import "fmt"

func main() {
	multiplica := multiplo(2)
	fmt.Println(multiplica())
	fmt.Println(multiplica())
	fmt.Println(multiplica())
   }
   func multiplo(num int) func() int {  
	n := 1
   return func() int {
	n *= num
	return n
	}
   }