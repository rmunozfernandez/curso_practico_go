package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var resultado int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	scanner.Scan()
	operador := scanner.Text()
	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	valor1, err1 := strconv.Atoi(valores[0])

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(valor1)
	}
	valor2, _ := strconv.Atoi(valores[1])

	switch operador {
	case "+":
		resultado = valor1 + valor2
	case "-":
		resultado = valor1 - valor2
	case "*":
		resultado = valor1 * valor2
	case "/":
		resultado = valor1 / valor2
	default:
		resultado = 0
	}

	fmt.Println(resultado)
}
