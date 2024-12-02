package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func ConvertBase(number string, fromBase int, toBase int) (string, error) {
	decimalNumber, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		return "", err
	}

	result := ""
	for decimalNumber > 0 {
		remainder := decimalNumber % int64(toBase)
		decimalNumber /= int64(toBase)
		result = strconv.FormatInt(remainder, toBase) + result
	}

	return strings.ToUpper(result), nil
}

func SolveQuadratic(a, b, c float64) (complex128, complex128) {
	discriminant := b*b - 4*a*c

	if discriminant >= 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	}

	realPart := -b / (2 * a)
	imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
	return complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)
}

func SortByAbsoluteValue(numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return math.Abs(float64(numbers[i])) < math.Abs(float64(numbers[j]))
	})
}

func MergeArrays(arr1, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2))

	for _, value := range arr1 {
		merged = append(merged, value)

	}
	for _, value := range arr2 {
		merged = append(merged, value)
	}

	return merged
}

func IndexOf(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func Calculator(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("деление на ноль")
		}
		return num1 / num2, nil
	case "^":
		result := 1.0
		for i := 0; i < int(num2); i++ {
			result *= num1
		}
		return result, nil
	case "%":
		if num2 == 0 {
			return 0, errors.New("деление на ноль в операции модуль")
		}
		return float64(int(num1) % int(num2)), nil
	default:
		return 0, errors.New("недопустимая операция")
	}
}

func isPalindrome(s string) bool {
	normalized := ""
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normalized += strings.ToLower(string(r))
		}
	}

	n := len(normalized)
	for i := 0; i < n/2; i++ {
		if normalized[i] != normalized[n-1-i] {
			return false
		}
	}
	return true
}

func isIntersection(segment1, segment2, segment3 [2]int) bool {
	left := max(segment1[0], segment2[0], segment3[0])
	right := min(segment1[1], segment2[1], segment3[1])

	return left <= right
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}
	return c
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func CleanAndSplit(sentence string) []string {
	var cleanedString strings.Builder
	for _, char := range sentence {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			cleanedString.WriteRune(char)
		}
	}
	return strings.Fields(cleanedString.String())
}

func FindLongestWord(sentence string) string {
	words := CleanAndSplit(sentence)
	longestWord := ""

	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

func IsArmstrongNumber(n int) bool {
	originalNum := n
	sum := 0
	numDigits := int(math.Log10(float64(n))) + 1

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		n /= 10
	}

	return sum == originalNum
}

func FindArmstrongNumbersInRange(start, end int) []int {
	var armstrongNumbers []int

	for i := start; i <= end; i++ {
		if IsArmstrongNumber(i) {
			armstrongNumbers = append(armstrongNumbers, i)
		}
	}

	return armstrongNumbers
}

func ReverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func TaskOneOne() {
	var number string
	var fromBase, toBase int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Print("Введите исходную систему счисления: ")
	fmt.Scan(&fromBase)
	fmt.Print("Введите конечную систему счисления: ")
	fmt.Scan(&toBase)

	convertedNumber, err := ConvertBase(number, fromBase, toBase)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Число в системе счисления %d: %s\n", toBase, convertedNumber)
}

func TaskOneTwo() {
	var a, b, c float64

	fmt.Print("Введите коэффициент a: ")
	fmt.Scan(&a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scan(&b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scan(&c)

	root1, root2 := SolveQuadratic(a, b, c)

	fmt.Printf("Корни уравнения: %v и %v\n", root1, root2)
}

func TaskOneThree() {
	numbers := []int{-10, 5, -3, 7, -2, 0, -8}
	fmt.Println("Исходный массив:", numbers)
	SortByAbsoluteValue(numbers)
	fmt.Println("Отсортированный массив по модулю:", numbers)
}

func TaskOneFour() {
	numbers1 := []int{-10, 5, -3, 7, -2, 0, -8}
	numbers2 := []int{-2, 0, -8}

	fmt.Println("Соединённый массив: ", MergeArrays(numbers1, numbers2))
}

func TaskOneFive() {
	var haystack, needle string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введините строку")
	if scanner.Scan() {
		haystack = scanner.Text()
	}
	fmt.Println("Введите подстроку для поиска")
	if scanner.Scan() {

		needle = scanner.Text()
	}
	result := IndexOf(haystack, needle)
	fmt.Printf("Первое вхождение подстроки \"%s\" в строке \"%s\": %d\n", needle, haystack, result)
}

func TaskTwoOne() {
	var num1, num2 float64
	var operator string

	// Ввод данных
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите оператор (+, -, *, /, ^, %): ")
	fmt.Scan(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	// Выполнение операции
	result, err := Calculator(num1, num2, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}
}

func TaskTwoTwo() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом")
	} else {
		fmt.Println("Строка не является палиндромом")
	}
}

func TaskTwoThree() {
	var s1, s2, s3 [2]int
	fmt.Println("Введите начальные и конечные точки первого отрезка:")
	fmt.Scan(&s1[0], &s1[1])
	fmt.Println("Введите начальные и конечные точки второго отрезка:")
	fmt.Scan(&s2[0], &s2[1])
	fmt.Println("Введите начальные и конечные точки третьего отрезка:")
	fmt.Scan(&s3[0], &s3[1])

	if isIntersection(s1, s2, s3) {
		fmt.Println("Пересечение существует")
	} else {
		fmt.Println("Пересечения нет")
	}
}

func TaskTwoFour() {
	var sentence string
	fmt.Println("Введите предложение:")
	Getline := func() string {
		var s string
		fmt.Scanln(&s)
		return s
	}
	sentence = Getline()
	longestWord := FindLongestWord(sentence)
	fmt.Printf("Самое длинное слово: %s\n", longestWord)
}

func TaskThreeThree() {
	var start, end int
	fmt.Println("Введите начальное число диапазона:")
	fmt.Scan(&start)
	fmt.Println("Введите конечное число диапазона:")
	fmt.Scan(&end)

	armstrongNumbers := FindArmstrongNumbersInRange(start, end)

	if len(armstrongNumbers) > 0 {
		fmt.Println("Числа Армстронга в заданном диапазоне:")
		for _, num := range armstrongNumbers {
			fmt.Println(num)
		}
	} else {
		fmt.Println("В данном диапазоне нет чисел Армстронга.")
	}
}

func TaskThreeFour() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scanln(&input)

	reversed := ReverseString(input)

	fmt.Println("Строка в обратном порядке:", reversed)
}

func TaskThreeFive() {
	var num1, num2 int
	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)

	result := Gcd(num1, num2)

	fmt.Printf("Наибольший общий делитель: %d\n", result)
}

func main() {
	TaskThreeFive()
}
