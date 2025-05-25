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
		fmt.Println("6. Ubah Data Assessment")
		fmt.Println("7. Hapus Data Assessment")
		fmt.Println("8. Urutkan berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("9. Rata-rata skor 1 bulan terakhir berdasarkan ID")
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
				id := InputString("Masukkan ID pengguna: ")
				RataRataSkorSebulan(id)
				
		case "0":
			return

		default:
			fmt.Println("Pilihan tidak dikenal.")
		}
	}
}
