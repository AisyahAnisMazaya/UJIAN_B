// 2311102095
package main

import (
	"fmt"
)

func main() {
	var M int
	fmt.Print("Masukkan jumlah kelompok: ")
	fmt.Scanln(&M)

	for i := 1; i <= M; i++ {
		var jumlahJenisMinuman, banyakOrang int
		var sisa bool
		fmt.Printf("Masukkan jumlah jenis minuman, banyak orang, dan apakah sisa (true/false) untuk kelompok %d: ", i)
		fmt.Scanln(&jumlahJenisMinuman, &banyakOrang, &sisa)

		totalBiaya := hitungTotalBiaya(jumlahJenisMinuman, banyakOrang, sisa)
		fmt.Printf("Total biaya untuk kelompok %d: Rp %d\n", i, totalBiaya)
	}
}

func hitungTotalBiaya(jumlahJenisMinuman, banyakOrang int, sisa bool) int {
	biayaPerMinuman := 10000
	totalBiaya := jumlahJenisMinuman * biayaPerMinuman * banyakOrang

	if sisa {
		biayaTambahan := (biayaPerMinuman * banyakOrang * 20) / 100
		totalBiaya += biayaTambahan
	}

	return totalBiaya
}
