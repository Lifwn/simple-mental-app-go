package main

import "fmt"


func main() {
	for {
		TampilkanMenu()
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
				fmt.Println("Ditemukan")
			} else {
				fmt.Println("Tidak ditemukan.")
			}

		case "3":
			id := InputString("Masukkan ID (pastikan data sudah terurut): ")
			idx := BinarySearch(id)
			if idx != -1 {
				fmt.Println("Ditemukan")
			} else {
				fmt.Println("Tidak ditemukan.")
			}

		case "4":
			SelectionSortBySkor(true)
			fmt.Println("Data diurutkan naik berdasarkan skor.")

		case "5":
			Tampilkan5Terakhir()
		
		case "6":
			id := InputString("Masukkan ID yang ingin diubah: ")
			UbahAssessment(id)

		case "7":
			id := InputString("Masukkan ID yang ingin dihapus: ")
			HapusAssessment(id)

		case "8":
			urutan := InputString("Ascending (a) atau Descending (d)? ")
			if urutan == "a" {
				InsertionSortByTanggal(true)
				fmt.Println("Data diurutkan berdasarkan tanggal naik.")
			} else if urutan == "d" {
				InsertionSortByTanggal(false)
				fmt.Println("Data diurutkan berdasarkan tanggal turun.")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		case "9":
			RataRataSkorSebulan()
				
		case "0":
			return

		default:
			fmt.Println("Pilihan tidak dikenal.")
		}
	}
}
