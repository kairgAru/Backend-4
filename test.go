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



}


