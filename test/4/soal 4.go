package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var attempts []int

func main() {
	//Generate the random integer

	min := 1
	max := 100
	secret := rand.Intn(max-min) + min

	//Create the reader

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter an integer between 1 and 100: ")

	for {
		entry, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		line := strings.Fields(entry)[0]
		num, err := strconv.Atoi(line)

		attempts = append(attempts, num)

		if err != nil {
			fmt.Println("masukan numer antara 1 dan 100.")
		} else if num < 1 || num > 100 {
			fmt.Println("masukan numer antara 1 dan 100.")
		} else {
			if secret-5 <= num && num < secret {
				fmt.Println("Anda sangat dekat lebih ke kiri.")
			} else if secret-10 <= num && num < secret-5 {
				fmt.Println("Anda sangat dekat lebih ke kiri.")
			} else if secret-20 <= num && num < secret-10 {
				fmt.Println("Anda lumayan dekat lebih kiri.")
			} else if num < secret-20 {
				fmt.Println("kau sangat jauh dari kiri")
			} else if secret+5 >= num && num > secret {
				fmt.Println("Anda sangat dekat lebih ke kanan.")
			} else if secret+10 >= num && num > secret+5 {
				fmt.Println("Anda sangat dekat lebih ke kanan.")
			} else if secret+20 >= num && num > secret+10 {
				fmt.Println("Anda lumayan dekat lebih kiri.")
			} else if num > secret+20 {
				fmt.Println("kau sangat jauh dari kiri")
			} else if secret == num {
				message := fmt.Sprintf("selamat!!!, kau menemukan numer nya setelah %d !!!\n", len(attempts))
				fmt.Println(message)
				break
			}
		}

		if secret != num {
			fmt.Println("Try again: ")
		}
	}

}