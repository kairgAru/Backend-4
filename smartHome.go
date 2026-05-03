package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const BaseTariff = 0.45
	const HighLoadTax = 0.15
	const NightDiscount = 0.30

	reader := bufio.NewReader(os.Stdin)

	for {
		var power float64
		var hours float64
		var isNight bool

		fmt.Print("Введите название прибора (или 'done'): ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "done" {
			fmt.Println("Расчет завершен.")
			break
		}

		fmt.Print("Мощность (Вт): ")
		fmt.Scanln(&power)

		fmt.Print("Время работы (ч): ")
		fmt.Scanln(&hours)

		fmt.Print("Ночной режим (true/false): ")
		fmt.Scanln(&isNight)

		energy := (power * hours) / 1000
		cost := energy * BaseTariff

		if isNight {
			cost = cost * (1 - NightDiscount)
		}

		if energy > 10 {
			cost = cost * (1 + HighLoadTax)
		}

		var category string
		switch {
		case power < 100:
			category = "Экономный"
		case power <= 1000:
			category = "Стандартный"
		default:
			category = "Мощный"
		}

		fmt.Println("\n--- Отчет по прибору ---")
		fmt.Printf("Прибор: %s [%s]\n", name, category)
		fmt.Printf("Расход: %.2f кВт·ч\n", energy)
		fmt.Printf("Итоговая стоимость: %.2f\n", cost)
		fmt.Println("------------------------")
	}
}