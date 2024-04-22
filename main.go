package main

import "fmt"

// HitungLuasPersegi menghitung luas persegi dengan sisi yang diberikan
func HitungLuasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func main() {
	sisi := 4.0
	luas := HitungLuasPersegi(sisi)
	fmt.Printf("Luas persegi dengan sisi %.2f adalah %.2f\n", sisi, luas)
}
