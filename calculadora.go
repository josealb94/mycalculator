package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (c calc) operate(entrada, operador string) int {
	valores := strings.Split(entrada, operador)
	operador1 := parsear(valores[0])
	operador2 := parsear(valores[1])
	var result int
	switch operador {
	case "+":
		result = operador1 + operador2
	case "-":
		result = operador1 - operador2
	case "*":
		result = operador1 * operador2
	case "/":
		result = operador1 / operador2
	default:
		fmt.Println(operador, "No esta soportado")
		return 0
	}
	return result
}

func parsear(data string) int {
	operador, err := strconv.Atoi(data)
	if err != nil {
		fmt.Println(err)
	}
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
