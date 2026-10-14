package main

import "fmt"

type Buku struct {
	Judul      string
	Genre      string
	TahunRilis int
	Penilaian  float64
}

func main() {
	var buku1, buku2 Buku

	fmt.Scan(&buku1.Judul, &buku1.Genre, &buku1.TahunRilis, &buku1.Penilaian)
	fmt.Scan(&buku2.Judul, &buku2.Genre, &buku2.TahunRilis, &buku2.Penilaian)

	fmt.Println("--- Book 1 ---")
	fmt.Println("Book Title:", buku1.Judul)
	fmt.Println("Book genre:", buku1.Genre)
	fmt.Println("Release year:", buku1.TahunRilis)
	fmt.Println("Rating:", buku1.Penilaian)

	fmt.Println("--- Book 2 ---")
	fmt.Println("Book Title:", buku2.Judul)
	fmt.Println("Book genre:", buku2.Genre)
	fmt.Println("Release year:", buku2.TahunRilis)
	fmt.Println("Rating:", buku2.Penilaian)
}
