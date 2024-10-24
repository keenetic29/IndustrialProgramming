package main

import (
	"fmt"
	"math"
)

func task_1_1(){
	var num int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	result := sumDigits(num)
	fmt.Printf("Сумма цифр числа %d равна %d\n", num, result)
}

func sumDigits(n int) int {
	if n == 0 {
		return 0
	}
	lastDigit := n % 10
	remainingNumber := n / 10
	return lastDigit + sumDigits(remainingNumber)
}

func task_1_2(){
	var gradus float64
	fmt.Print("Введите число градусов: ")
	fmt.Scan(&gradus)
	result := Farengeite(gradus)
	fmt.Print("Градусы в фаренгейтах: ", result)
}

func Farengeite(gradus float64) float64 {
	farengeite := gradus * 1.8 + 32
	return farengeite
}


func task_1_3(){
	var count int
	fmt.Print("Введите кол-во элементов массива: ")
	fmt.Scan(&count)
	var slice []int = make([]int, count)
	fmt.Print("Введите элементы массива: ")
	for i := 0; i < count; i++ {
		fmt.Scan(&slice[i])
	}
	doubleElement(slice)
	fmt.Print(slice)
}

func doubleElement(slice []int){
	slice[0] = slice[0] * 2
	if len(slice) > 1 {
		doubleElement(slice[1:])
	}
}

func task_1_4(){
	var str1, str2, result string
	fmt.Print("Введите 1ую строку: ")
	fmt.Scan(&str1)
	fmt.Print("Введите 2ую строку: ")
	fmt.Scan(&str2)
	result = str1 + str2
	fmt.Print(result)
}

func task_1_5(){
	fmt.Println("Введите координаты")
	var x1, y1, x2, y2 int
	_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Print(distance(x1, y1, x2, y2))

}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)))
}

func task_2_1(){
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number % 2 == 0 {
		fmt.Print("Четное")
	} else {
		fmt.Print("Нечетное")
	}
}

func task_2_2(){
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)
	if year % 400 == 0 || year % 4 == 0 && year % 100 != 0 {
		fmt.Print("Високосный")
	} else {
		fmt.Print("Не високосный")
	}
}

func task_2_3(){
	var number1, number2, number3 int
	fmt.Print("Введите 3 числа: ")
	fmt.Scan(&number1, &number2, &number3)
	fmt.Print("Максимальное число = ", max(number1, number2, number3))
}

func task_2_4(){
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	if age < 12 {
		fmt.Print("Ребенок")
	} else if age < 18 {
		fmt.Print("Подросток")
	} else if age < 66 {
		fmt.Print("Взрослый")
	} else {
		fmt.Print("Пожилой")
	}
}

func task_2_5() {
	var number int
	fmt.Println("Введите число")
	fmt.Scan(&number)
	if number % 3 == 0 && number % 5 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

func task_3_1() {
	var number int
	fmt.Println("Введите число")
	fmt.Scan(&number)
	fmt.Print(factorial(number))
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func task_3_2() {
	var number int
	fmt.Println("Введите число")
	fmt.Scan(&number)
	fib(number)
}

func fib(n int) {
	var fib1, fib2 int = 0, 1
	if n == 1 {
		fmt.Print(fib1)
	} else if n == 2 {
		fmt.Print(fib1, ", ", fib2)
	} else {
		fmt.Print(fib1, ", ", fib2)
		for i := 3; i <= n; i++ {
			fib1, fib2 = fib2, fib1 + fib2
			fmt.Print(", ", fib2)
		}
	}
}


func task_3_3() {
	var count int
	fmt.Println("Введите количество элементов в массиве")
	_, _ = fmt.Scan(&count)
	slice := make([]int, count)
	fmt.Println("Введите элементы массива")
	for i := 0; i < count; i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	reverseArray(slice)
	fmt.Print(slice)
}

func reverseArray(slice []int) {
	count := len(slice)
	for i := 0; i < count / 2; i++ {
		slice[i], slice[count - i - 1] = slice[count - i - 1], slice[i]
	}
}

func task_3_4() {
	fmt.Println("Введите число")
	var number int
	_, _ = fmt.Scan(&number)
	for i := 2; i <= number; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
}

func isPrime(n int) bool {
	result := true
	for i := 2; i < n; i++ {
		if n % i == 0 {
			result = false
			break
		}
	}
	return result
}


func task_3_5() {
	fmt.Println("5. Сумма чисел в массиве")
	var n int
	fmt.Println("Введите количество элементов в массиве")
	_, _ = fmt.Scan(&n)
	slice := make([]int, n)
	fmt.Println("Введите элементы массива")
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	fmt.Print(sumOfArr(slice))
}

func sumOfArr(slice []int) int {
	sum := 0
	l := len(slice)
	for i := 0; i < l; i++ {
		sum += slice[i]
	}
	return sum
}


func main() {
	//task_1_1()
	//task_1_2()
	//task_1_3()
	//task_1_4()
	//task_1_5()
	//task_2_1()
	//task_2_2()
	//task_2_3()
	//task_2_4()
	//task_2_5()
	//task_3_1()
	//task_3_2()
	//task_3_3()
	//task_3_4()
	task_3_5()
}
