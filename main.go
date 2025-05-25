package main

import "fmt"

func main() {
	for {
		fmt.Println("\n==== Menu ====")
		fmt.Println("1. Tambah Self-Assessment")
		fmt.Println("2. Cari dengan Sequential Search")
		fmt.Println("3. Cari dengan Binary Search (butuh data terurut)")
		fmt.Println("4. Urutkan berdasarkan Skor (Selection Sort)")
		fmt.Println("5. Tampilkan 5 terakhir")
		fmt.Println("0. Keluar")
		pilihan := InputString("Pilih: ")

		switch pilihan {
		case "1":
			id := InputString("Masukkan ID: ")
			jawaban := InputJawaban()
			tanggal := InputTanggal()
			TambahAssessment(id, jawaban, tanggal)

		case "2":
			id := InputString("Masukkan ID untuk dicari: ")
			idx := SequentialSearch(id)
			if idx != -1 {
				fmt.Println("Ditemukan:", dataAssessment[idx])
			} else {
				fmt.Println("Tidak ditemukan.")
			}

		case "3":
			id := InputString("Masukkan ID (pastikan data sudah terurut): ")
			idx := BinarySearch(id)
			if idx != -1 {
				fmt.Println("Ditemukan:", dataAssessment[idx])
			} else {
				fmt.Println("Tidak ditemukan.")
			}

		case "4":
			SelectionSortBySkor(true)
			fmt.Println("Data diurutkan naik berdasarkan skor.")

		case "5":
			Tampilkan5Terakhir()

		case "0":
			return

		default:
			fmt.Println("Pilihan tidak dikenal.")
		}
	}
}
