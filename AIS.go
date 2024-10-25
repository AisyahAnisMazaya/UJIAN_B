// 2311102095
package main

import "fmt"

func countEvenNumbers(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if numbers[0] < 0 {
		return 0
	}
	if numbers[0]%2 == 0 {
		return 1 + countEvenNumbers(numbers[1:])
	}
	return countEvenNumbers(numbers[1:])
}

func main() {
	var numbers []int
	var input int
	fmt.Println("Masukkan bilangan bulat (bilangan negatif sebagai sentinel):")
	for {
		fmt.Scanln(&input)
		if input < 0 {
			break
		}
		numbers = append(numbers, input)
	}
	fmt.Printf("Jumlah angka genap (metode rekursif): %d\n", countEvenNumbers(numbers))
}
