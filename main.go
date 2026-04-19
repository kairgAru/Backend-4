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




}




