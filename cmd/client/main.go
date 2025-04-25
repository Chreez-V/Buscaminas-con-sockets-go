package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"buscaminas/internal/game"
)

func main() {
	board := game.NewBoard(5, 5, 5)
	reader := bufio.NewReader(os.Stdin)

	for {
		board.Print()
		fmt.Print("Ingresa coordenadas (x y): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Entrada inválida. Usa el formato: x y")
			continue
		}
		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Coordenadas inválidas")
			continue
		}
		success := board.Reveal(x, y)
		if !success {
			board.Print()
			fmt.Println("¡Boom! Pisaste una mina. Fin del juego.")
			break
		}
	}
}

