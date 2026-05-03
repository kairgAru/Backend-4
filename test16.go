package main
import "fmt"
		
func main() {

	toolUsage := map[string]int{
	"Go":     3,
	"VSCode": 5,
	"Git":    2,
}

for k, v := range toolUsage {
	fmt.Println(k, v)
}

//2

buildStatus := map[string]bool{
	"build": true,
	"run":   false,
}

if buildStatus["build"] {
	fmt.Println("Сборка прошла успешно")
}


//3

var name string
println("Введите имя:")
fmt.Scanln(&name)

userInfo := map[string]interface{}{
	"name":       name,
	"isLoggedIn": true,
}

fmt.Println("Пользователь", userInfo["name"], "вошёл в систему")
	

	//4 

	cpuLoad := map[int]int{
	1: 40,
	2: 65,
	3: 30,
}

maxCore := 0
maxLoad := 0

for k, v := range cpuLoad {
	if v > maxLoad {
		maxLoad = v
		maxCore = k
	}
}

fmt.Println("Ядро:", maxCore)


//5


examResults := map[string]int{
	"Aruzhan": 85,
	"Dias":    92,
	"Alina":   78,
}

for name, grade := range examResults {
	if grade >= 80 {
		fmt.Println(name, "сдал")
	} else {
		fmt.Println(name, "не сдал")
	}
}


//6

words := []string{"go", "is", "fast"}
wordLength := map[string]int{}

for _, w := range words {
	wordLength[w] = len(w)
}

fmt.Println(wordLength)

//7

menu := map[string]int{
	"Burger": 2500,
	"Pizza":  3200,
	"Tea":    500,
}

var dish string
fmt.Scanln(&dish)

if price, ok := menu[dish]; ok {
	fmt.Println(price)
} else {
	fmt.Println("Блюдо не найдено")
}


//8
loginAttempts := map[string]int{
	"admin": 2,
	"guest": 0,
}

loginAttempts["admin"]++

if loginAttempts["admin"] > 2 {
	fmt.Println("Доступ заблокирован")
}

//9

sales := [2][3]int{
	{10, 20, 30},
	{15, 25, 35},
}

total := map[int]int{}

for i := 0; i < 2; i++ {
	sum := 0
	for j := 0; j < 3; j++ {
		sum += sales[i][j]
	}
	total[i] = sum
}

fmt.Println(total)

//10

numbers := []int{4, 7, 2, 9, 5}
numberStatus := map[int]string{}

for _, n := range numbers {
	if n%2 == 0 {
		numberStatus[n] = "even"
	} else {
		numberStatus[n] = "odd"
	}
}

fmt.Println(numberStatus)

//11

defaultConfig := map[string]string{
	"host": "localhost",
	"port": "8080",
	"mode": "production",
}

currentConfig := map[string]string{
	"host": "localhost",
	"port": "8080",
	"mode": "production",
}

equal := true
for k, v := range defaultConfig {
	if currentConfig[k] != v {
		equal = false
	}
}

if equal {
	fmt.Println("Совпадают")
} else {
	fmt.Println("Отличаются")
}

currentConfig["mode"] = "debug"

equal = true
for k, v := range defaultConfig {
	if currentConfig[k] != v {
		equal = false
	}
}

if equal {
	fmt.Println("Совпадают")
} else {
	fmt.Println("Отличаются")
}

}

