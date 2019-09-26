//Faça um programa que utilize closures para gerar
//sequências de 6 números pseudo-aleatórios não repetidos
//(a cada 6 chamadas a sequência é reinicializada).
package main

import (
    "fmt"
    "math/rand"
)

func main() {
	multiplica := multiplo()
	fmt.Println(multiplica())
	fmt.Println(multiplica())
	fmt.Println(multiplica())
   }
   func multiplo() func() int {  
	combinacao := []
   return func() int {
	for j := 0; j <= 6; j++ {
		for len(combinacao) <= j{
			aux := rand.Intn(60)
			if aux
		}
        
    }
	return n
	}
   }