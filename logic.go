package main

import (
	"fmt"
	"time"
)

func TambahAssessment(id string, jawaban [10]int, tanggal time.Time) {
	if jumlahData >= MaxData {
		fmt.Println("Data penuh.")
		return
	}

	skor := HitungSkor(jawaban)

	dataAssessment[jumlahData] = Assessment{
		IDUser:    id,
		Tanggal:   tanggal,
		Jawaban:   jawaban,
		SkorTotal: skor,
	}

	jumlahData++
	fmt.Println("Skor Total Anda:", skor)
	fmt.Println("Interpretasi Skor Anda:", InterpretasiSkor(skor))
	fmt.Println("Data berhasil ditambahkan.")
}

func SequentialSearch(id string) int {
	for i := 0; i < jumlahData; i++ {
		if dataAssessment[i].IDUser == id {
			dataAssessment[i].Print("Data ditemukan (Sequential Search):")
			return i
		}
	}
	return -1
}


func BinarySearch(id string) int {
	low, high := 0, jumlahData-1

	for low <= high {
		mid := (low + high) / 2
		if dataAssessment[mid].IDUser == id {
			dataAssessment[mid].Print("Data ditemukan (Binary Search):")
			return mid
		} else if dataAssessment[mid].IDUser < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Data tidak ditemukan.")
	return -1
}


func SelectionSortBySkor(ascending bool) {
	for i := 0; i < jumlahData-1; i++ {
		idx := i
		for j := i + 1; j < jumlahData; j++ {
			if (ascending && dataAssessment[j].SkorTotal < dataAssessment[idx].SkorTotal) || 
			   (!ascending && dataAssessment[j].SkorTotal > dataAssessment[idx].SkorTotal) {
				idx = j
			}
		}
		dataAssessment[i], dataAssessment[idx] = dataAssessment[idx], dataAssessment[i]
	}
}

func Tampilkan5Terakhir() {
	fmt.Println("5 Self-Assessment Terakhir:")
	start := jumlahData - 5
	if start < 0 {
		start = 0
	}
	for i := jumlahData - 1; i >= start; i-- {
		dataAssessment[i].Print(fmt.Sprintf("\n[%s]", dataAssessment[i].IDUser))
		fmt.Println("----------------------------")
	}
}


func UbahAssessment(id string) {
    idx := SequentialSearch(id)
    if idx == -1 {
        fmt.Println("Data tidak ditemukan.")
        return
    }

    fmt.Println("Data ditemukan. Silakan masukkan data baru.")
    jawabanBaru := InputJawaban()

    dataAssessment[idx].Jawaban = jawabanBaru
    dataAssessment[idx].Tanggal = InputTanggal()
    dataAssessment[idx].SkorTotal = HitungSkor(jawabanBaru)

    fmt.Println("Data berhasil diubah.")
}

func HapusAssessment(id string) {
    idx := SequentialSearch(id)
    if idx == -1 {
        fmt.Println("Data tidak ditemukan.")
        return
    }

    for i := idx; i < jumlahData-1; i++ {
        dataAssessment[i] = dataAssessment[i+1]
    }

    jumlahData--
    fmt.Println("Data berhasil dihapus.")
}

func InsertionSortByTanggal(ascending bool) {
    for i := 1; i < jumlahData; i++ {
        temp := dataAssessment[i]
        j := i - 1

        for j >= 0 && ((ascending && dataAssessment[j].Tanggal.After(temp.Tanggal)) || 
                       (!ascending && dataAssessment[j].Tanggal.Before(temp.Tanggal))) {
            dataAssessment[j+1] = dataAssessment[j]
            j--
        }
        dataAssessment[j+1] = temp
    }
}

func RataRataSkorSebulan() {
	totalSkor := 0
	jumlah := 0
	sekarang := time.Now()

	for i := 0; i < jumlahData; i++ {
		selisih := sekarang.Sub(dataAssessment[i].Tanggal)
		if selisih.Hours() <= 720 {
			totalSkor += dataAssessment[i].SkorTotal
			jumlah++
		}
	}

	if jumlah == 0 {
		fmt.Println("Tidak ada data dalam 1 bulan terakhir.")
		return
	}

	rata := float64(totalSkor) / float64(jumlah)
	fmt.Printf("Rata-rata skor semua assessment dalam 1 bulan terakhir: %.2f\n", rata)
}

