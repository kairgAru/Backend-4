package main
import "fmt"
		
func main() {

	runningExercises := [2]string{"Спринт", "Бег на длинную дистанцию"}
	walkingExercises := [2]string{"Быстрая ходьба", "Скандинавская ходьба"}
	fmt.Println("Количество упражнений для бега:", len(runningExercises))
	fmt.Println("Количество упражнений для ходьбы:", len(walkingExercises))


	subjectsList := [3]string{"Физика", "Химия", "География"}
	fmt.Println("Первый элемент:", subjectsList[0])
	fmt.Println("Последний элемент:", subjectsList[2])
	subjectsList[1] = "Биология"
	fmt.Println("Обновлённый массив:", subjectsList)


	data := [3]string{"Tom", "35", "New York"}
	name := data[0]
	age := data[1]
	home := data[2]
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(home)


	numbersList := [5]int{1, 2, 3, 4, 5}
	found := false
	for _, value := range numbersList {
		if value == 3 {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Число 3 есть в массиве")
	} else {
		fmt.Println("Число 3 отсутствует в массиве")
	}
		
	
	
	friendsList := [4]string{"Ali", "Dana", "Arman", "Aruzhan"}
	found = false
	for _, friend := range friendsList {
		if friend == "Bekbolat" {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Мне очень повезло")
	} else {
		fmt.Println("Мне не повезло")
	}


	firstList := [3]int{1, 2, 3}
	secondList := [3]int{1, 2, 4}
	if firstList == secondList {
		fmt.Println("Массивы равны")
	} else {
		fmt.Println("Массивы не равны")
	}


	myWishList := [3]string{"Книга", "Наушники", "Клавиатура"}
	friendsWishList := [3]string{"Часы", "Игровая мышь", "Рюкзак"}
	var registrationList [6]string
	for i := 0; i < len(myWishList); i++ {
		registrationList[i] = myWishList[i]
	}
	for i := 0; i < len(friendsWishList); i++ {
		registrationList[i+len(myWishList)] = friendsWishList[i]
	}
	fmt.Println(registrationList)


	toyList := [3]string{"Car", "Doll", "Ball"}
	testToyList := toyList
	testToyList[1] = "Boat"
	fmt.Println("toyList:", toyList)
	fmt.Println("testToyList:", testToyList)

// Урок 14


	var arr [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	max := arr[0][0]
	maxRow := 0
	maxCol := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if arr[i][j] > max {
				max = arr[i][j]
				maxRow = i
				maxCol = j
			}
		}
	}
	fmt.Println(maxRow, maxCol)



	var arr1 [11][11]string
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			arr1[i][j] = "."
		}
	}
	mid := 11 / 2 
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if i == mid || j == mid || i == j || i+j == 10 {
				arr1[i][j] = "*"
			}
		}
	}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Print(arr1[i][j], " ")
		}
		fmt.Println()
	}


	var board [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = "." 
			} else {
				board[i][j] = "*"
			}
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}



	var arr2 [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}
	var i, j int
	fmt.Scan(&i, &j)
	for col := 0; col < 4; col++ {
		arr2[i][col], arr2[j][col] = arr2[j][col], arr2[i][col]
	}
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Print(arr2[r][c], " ")
		}
		fmt.Println()
	}



	var arr3 [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&arr3[i][j])
		}
	}
	var col1, col2 int
	fmt.Scan(&col1, &col2)
	for row := 0; row < 4; row++ {
		arr3[row][col1], arr3[row][col2] = arr3[row][col2], arr3[row][col1]
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arr3[i][j], " ")
		}
		fmt.Println()
	}

}


