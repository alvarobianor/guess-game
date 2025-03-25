package games

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Guess() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println(
		"Um número aletório será sorteado. Tente acertar. O número é um inteiro entre 0 e 30")

	randomNumber := rand.Intn(31)
	leftTentatives := 10
	var myTentatives, input = []int{}, 0

	scanner := bufio.NewScanner(os.Stdin)
	for leftTentatives > 0 {

		fmt.Println("Digite um número: ")
		scanner.Scan()
		valueInt, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		input = valueInt

		if err != nil || valueInt < 0 || valueInt > 30 {
			fmt.Println("Por favor, digite um número válido. Você ainda tem", leftTentatives, "tentativas")
			continue
		}

		var hasTriedThisNumber bool = false

		for _, myTentative := range myTentatives {
			if valueInt == myTentative {
				fmt.Println("Você já tentou esse número. Tente outro.")
				hasTriedThisNumber = true
				continue
			}
		}

		if valueInt == randomNumber {
			fmt.Println("Parabéns! Você acertou!")
			break
		}

		if !hasTriedThisNumber {
			leftTentatives--
		}

		myTentatives = append(myTentatives, valueInt)
		if leftTentatives == 0 {
			break
		}
		fmt.Println("Você errou! Você ainda tem", leftTentatives, "tentativas")
	}

	defer func() {
		if input == randomNumber {
			fmt.Println("O número era:", randomNumber)
		} else {
			fmt.Println("Você perdeu! O número era:", randomNumber)
		}
		fmt.Println("Essas foram as suas tentativas:", myTentatives)
	}()
}
