package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Эспрессо": 800,
		"Латте":    1200,
		"Капучино": 1100,
		"Сэндвич":  1500,
		"Круассан": 900,
	}

	order := []string{}

	fmt.Println("=== МЕНЮ ===")
	for name, price := range menu {
		fmt.Printf("%s — %.0f тг\n", name, price)
	}

	for {
		var item string
		fmt.Print("\nВведите блюдо (или 'exit'): ")
		fmt.Scanln(&item)

		if item == "exit" {
			break
		}

		if _, ok := menu[item]; ok {
			order = append(order, item)
			fmt.Println("Добавлено!")
		} else {
			fmt.Println("К сожалению, этого блюда нет в меню")
		}
	}

	if len(order) == 0 {
		fmt.Println("Заказ пуст")
		return
	}

	countMap := map[string]int{}
	for _, item := range order {
		countMap[item]++
	}

	var total float64

	fmt.Println("\n=== ЧЕК ===")
	for item, count := range countMap {
		price := menu[item]
		sum := price * float64(count)
		total += sum

		fmt.Printf("%s x%d — %.0f тг\n", item, count, sum)
	}

	fmt.Printf("\nСумма: %.0f тг\n", total)

	var discount float64
	if total > 5000 {
		discount = total * 0.10
		fmt.Printf("Скидка: %.0f тг\n", discount)
	}

	afterDiscount := total - discount

	nds := afterDiscount * 0.12
	fmt.Printf("НДС: %.0f тг\n", nds)

	final := afterDiscount + nds
	fmt.Printf("Итого к оплате: %.0f тг\n", final)
}