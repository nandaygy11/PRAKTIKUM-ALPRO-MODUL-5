package main

import "fmt"

func main() {
	var hargaAwal, hargaDiskonPertama, hargaAkhir float64

	fmt.Scan(&hargaAwal)

	hargaDiskonPertama = hargaAwal - (hargaAwal * 0.10)
	hargaAkhir = hargaDiskonPertama - (hargaDiskonPertama * 0.20)

	fmt.Printf("%.2f\n", hargaAkhir)
}
