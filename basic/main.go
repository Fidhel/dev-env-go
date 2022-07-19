package main

import (
	"fmt"
)

func mult_return(first, last string) (string, int) {
	return first, 35
}

func say(s string) {

}

func special_sun(value int, v chan int) {
	sum := 0
	sum += value
	v <- sum
}

//funcao com letra minuscula é privada, com letra maiscula e publica.
func main() {
	//cria um canal para receber o que está sendo processado pro go-routines
	v := make(chan int)

	//processamento paralelo
	go special_sun(5, v)
	go special_sun(10, v)

	x, y := <-v, <-v

	fmt.Printf("Vuvu Hello %d", x+y)
}
