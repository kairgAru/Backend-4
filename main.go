package main
import ("fmt"
		"math")
func main() {
	// урок 4
	fmt.Println("Я сегодня изучил:")
	fmt.Println("Основы языка Go")

	//урок 5
	schooling:=11
	fmt.Println(schooling)
	schooling=12
	fmt.Println(schooling)

	name:="Vladislav"
	fmt.Println(name)
	name="Aru"
	fmt.Println(name)

	steps:=0
	fmt.Println(steps)
	steps=200
	fmt.Println(steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели")

	largeNumber:= 6000000
	fmt.Println(largeNumber)

	// const breakTime = 15
	// fmt.Println(breakTime)
	// const breakTime= 20
	// fmt.Println(breakTime)
	// конст нельзя поменять


	// Урок 6
	 age := 23
	 fmt.Println(age)
	 age= age + 1
	 fmt.Println(age)


	 height := 170
	 height_in_meters := 1.7
	 fmt.Println(height, height_in_meters)

	 isStudent := true
	 fmt.Println(isStudent)

	 temperature := 25
	 fmt.Println("Погода теплая:", temperature)


	 favoriteQuote :="Отличный солнесный день!"
	 fmt.Println(favoriteQuote)

	  PI := 3.14
	  fmt.Println(PI)
	//   PI = "3.1415"
	//   fmt.Println(PI)
	//  нельзя присвоить строку переменной типа float64

	// Урок 7
	bannerWidth:= 12
	bannerHeight :=8
	bannerArea := bannerHeight*bannerWidth
	halfBannerArea := bannerArea/2
	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println(bannerArea, halfBannerArea,bannerBorderLength)

	boxCount :=29
	leftoverBoxes:= boxCount % 5
	fmt.Println(leftoverBoxes)

	tempMorning:= 16
	tempAfternoon := 20
	tempEvening := 28
	totalTemp:= tempAfternoon + tempEvening + tempMorning
	averageTemp:= totalTemp/3
	fmt.Println(totalTemp, averageTemp)

	knownWords :=47
	wordsGoal :=120
	progressPercent:= float64(knownWords) / float64(wordsGoal) * 100
	fmt.Println(progressPercent)

	coins := 0
	fmt.Println(coins)
	coins+=500
	fmt.Println(coins)
	coins+=1200
	fmt.Println(coins)
	coins/=2
	fmt.Println(coins)
	coins*=2
	fmt.Println(coins)
	coins-=300
	fmt.Println(coins)

	participants := 42
	groupCount :=8
	participantsPerGroup:= participants/groupCount
	fmt.Println(participantsPerGroup)

	fmt.Println(20 - 4 * 3)
	fmt.Println((20 - 4) * 3)


	squareValue := 81
	squareRoot := math.Sqrt(float64(squareValue))
	multiplier := 5
	exponent := 2
	powerResult := math.Pow(float64(multiplier), float64(exponent))
	fmt.Println(squareRoot,powerResult)
	
	
	// Урок 9

	// tr
	fmt.Println(5 == 5)
	// tr
	fmt.Println(10 != 3)
	// fl
	fmt.Println(7 > 12)
	// tr
	fmt.Println(15 < 20)
	fmt.Println(8 >= 8)
	fmt.Println(6 <= 4)
	fmt.Println((10 > 5) && (3 < 1))
	fmt.Println((10 > 5) || (3 < 1))
	fmt.Println(!(5 == 5))
	fmt.Println(true && false)
	fmt.Println(false || false)
	fmt.Println(true || false)
	fmt.Println((4 + 6 == 10) && (9 > 2))
	fmt.Println((12 / 3 == 4) || (8 < 5))


	age1 := 16
	hasTicket :=true
	canEnter := age1 >= 18 && hasTicket
	fmt.Println(canEnter)

	isLoggedIn := true
	isAdmin := true
	hasAccess:= isAdmin && isLoggedIn || isLoggedIn && !isAdmin
	fmt.Println(hasAccess)

	// Урок 10
	temperature2 :=21
	if temperature2 < 0 {
		fmt.Println("Холодно")
	} else if temperature2 >= 0 && temperature2 <= 20{
			fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}

	score2 := 66
	if score2>=90 {
		fmt.Println("Отлично")
	} else if score2 >= 70 && score2<= 89{
		fmt.Println("Хорошо")
	}else if score2 >= 50 && score2<= 69 {
		fmt.Println("Удовлетворительно")
	}else if score2 <= 50{
		fmt.Println("Не сдал")
	}


	hour := 14
	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Ночь")
	case hour >= 6 && hour <= 11:
		fmt.Println("Утро")
	case hour >= 12 && hour <= 17:
		fmt.Println("День")
	case hour >= 18 && hour <= 23:
		fmt.Println("Вечер")
	default:
		fmt.Println("Некорректное значение времени")
	}
	

	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("Чётное число")
	} else {
		fmt.Println("Нечётное число")
	}


	day := "Monday" 
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	default:
		fmt.Println("Некорректный день")
	}


	balance := 100 
	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}


	var age3 int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age3)
	if age3 < 13 {
		fmt.Println("Ребёнок")
	} else if age3 >= 13 && age3 <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}


	var command string
	fmt.Print("Введите команду: ")
	fmt.Scan(&command)
	switch command {
	case "start":
		fmt.Println("Запуск")
	case "stop":
		fmt.Println("Остановка")
	case "restart":
		fmt.Println("Перезапуск")
	default:
		fmt.Println("Неизвестная команда")
	}


	grade := 4 
	switch grade {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 2:
		fmt.Println("D")
	case 1:
		fmt.Println("F")
	default:
		fmt.Println("Некорректная оценка")
	}


	// Урок 11

	for i:=1; i<=20 ; i++ {
		fmt.Println( i)
	}


	sum :=0
	for i:=1 ; i<=100; i++ {
		sum +=i
	}
	fmt.Println(sum)


	number3 := 5 

	for i := 1; i <= 10; i++ {
		fmt.Println(number3, "x", i, "=", number3*i)
	}


	var t int

	fmt.Print("Введите число: ")
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}


	var num int
	count := 0
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if num == 0 {
		count = 1
	} else {
		if num < 0 {
			num = -num
		}
		for num > 0 {
			num = num / 10
			count++
		}
	}
	fmt.Println("Количество цифр:", count)


	text := "Developer"
	for i := 0; i < len(text); i++ {
		fmt.Println(string(text[i]))
	}


	balance1 := 3000

	for {
		var choice int

		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - Показать баланс")
		fmt.Println("2 - Пополнить (+500)")
		fmt.Println("3 - Снять (-200)")
		fmt.Println("0 - Выход")
		fmt.Print("Ваш выбор: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Текущий баланс:", balance1)
		case 2:
			balance1 += 500
			fmt.Println("Баланс пополнен. Новый баланс:", balance1)
		case 3:
			balance1 -= 200
			fmt.Println("Списано 200. Новый баланс:", balance1)
		case 0:
			fmt.Println("Выход из программы")
			break
		default:
			fmt.Println("Неверный ввод")
		}

		if choice == 0 {
			break
		}
	}






}




