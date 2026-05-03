package main
import "fmt"
		
func main() {
	

	
	// урок 15
	var numbers []int
for {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		break
	}
	numbers = append(numbers, n)
} 
sum := 0
for _, v := range numbers {
	sum += v
}
fmt.Println(sum)

// задание 2
var values []int

for {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		break
	}
	values = append(values, n)
}

var even []int
for _, v := range values {
	if v%2 == 0 {
		even = append(even, v)
	}
}

fmt.Println(values)
fmt.Println(even)

// 3
var data []int

for {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		break
	}
	data = append(data, n)
}

if len(data) > 2 {
	data = append(data[:2], data[3:]...)
}

fmt.Println(data)
	

// 4
var temps []int

for {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		break
	}
	temps = append(temps, n)
}

min := temps[0]
max := temps[0]

for _, t := range temps {
	if t < min {
		min = t
	}
	if t > max {
		max = t
	}
}

fmt.Println(min, max)

// 5
var words []string

for {
	var w string
	fmt.Scan(&w)
	if w == "stop" {
		break
	}
	words = append(words, w)
}

var reversed []string
for i := len(words) - 1; i >= 0; i-- {
	reversed = append(reversed, words[i])
}

fmt.Println(reversed)

// 6

var nums []int

for {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		break
	}
	nums = append(nums, n)
}

sorted := true

for i := 1; i < len(nums); i++ {
	if nums[i] < nums[i-1] {
		sorted = false
		break
	}
}

fmt.Println(sorted)

// 7 
myWishList := [3]string{"Phone", "Book", "Laptop"}
friendsWishList := [3]string{"Game", "Watch", "Headphones"}

var result []string

for _, v := range myWishList {
	result = append(result, v)
}

for _, v := range friendsWishList {
	result = append(result, v)
}

fmt.Println(result)
// 8

toyList := [3]string{"Car", "Doll", "Ball"}
testToyList := toyList

testToyList[1] = "Boat"

fmt.Println(toyList)
fmt.Println(testToyList)





















}
