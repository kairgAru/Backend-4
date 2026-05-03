package main

import "fmt"

// Задание 1
func square(n int) int {
	return n * n
}

// Задание 2
func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Задание 3
func isEven(n int) bool {
	return n%2 == 0
}

// Задание 4
func greetUser(name string) {
	fmt.Println("Привет,", name)
}

// Задание 5
func sumSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Задание 6
func checkLogin(login, password string) bool {
	users := map[string]string{
		"admin": "1234",
	}
	return users[login] == password
}

// Задание 7
func increaseBalance(balance *int, amount int) {
	*balance += amount
}

// Задание 8
func resetAttempts(attempts *int) {
	if *attempts > 3 {
		*attempts = 0
	}
}

func main() {

	// 1
	fmt.Println(square(2))
	fmt.Println(square(5))

	// 2
	fmt.Println(maxNumber(10, 5))

	// 3
	fmt.Println(isEven(4))
	fmt.Println(isEven(7))

	// 4
	greetUser("Aruzhan")
	greetUser("Dias")

	// 5
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(nums))

	// 6
	fmt.Println(checkLogin("admin", "1234"))
	fmt.Println(checkLogin("admin", "0000"))

	// 7
	balance := 100
	increaseBalance(&balance, 50)
	fmt.Println(balance)

	// 8
	attempts := 5
	resetAttempts(&attempts)
	fmt.Println(attempts)
}