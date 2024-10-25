// 2311102095
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah mahasiswa (N): ")
	fmt.Scanln(&n)

	var tiket []string
	fmt.Println("Masukkan nomor tiket:")
	for i := 0; i < n; i++ {
		var num string
		fmt.Scanln(&num)
		tiket = append(tiket, num)
	}

	validCount, invalidCount := 0, 0
	for _, t := range tiket {
		if isValidTicket(t) {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Printf("Tiket valid: %d\n", validCount)
	fmt.Printf("Tiket tidak valid: %d\n", invalidCount)
}

func isValidTicket(ticket string) bool {
	length := len(ticket)
	if length != 6 && length != 8 {
		return false
	}

	firstTwo, _ := strconv.Atoi(ticket[:2])
	lastTwo, _ := strconv.Atoi(ticket[length-2:])
	if firstTwo != lastTwo {
		return false
	}

	midStart := length/2 - 1
	midEnd := midStart + 2
	for i := midStart; i < midEnd; i++ {
		digit, _ := strconv.Atoi(string(ticket[i]))
		if digit <= 5 {
			return false
		}
	}

	return true
}
