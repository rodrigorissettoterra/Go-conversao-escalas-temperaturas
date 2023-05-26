package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	var F = ebulicaoF
	var C = (F - 32) * 5 / 9
	var K = C + 273
	fmt.Println("O ponto de ebulição da água, em fahrenheit, é:", F, "ºF.")
	fmt.Println("O ponto de ebulição da água, em celsius, é:", C, "ºC.")
	fmt.Println("O ponto de ebulição da água, em kelvin, é:", K, "K.")
}
