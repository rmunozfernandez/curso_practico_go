package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(operacion string, operador string) int {
	var resultado int

	valores := strings.Split(operacion, operador)

	valor1 := textToInt(valores[0])
	valor2 := textToInt(valores[1])

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

	return resultado
}

func entradaCLI() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func textToInt(texto string) int {
	valor, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println(err)
	}
	return valor
}

func main() {
	operacion := entradaCLI()
	operador := entradaCLI()

	c := calc{}
	fmt.Println(c.operate(operacion, operador))
}
