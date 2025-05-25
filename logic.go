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

	skor := 0
	for _, j := range jawaban {
		skor += j
	}

	dataAssessment[jumlahData] = Assessment{
		IDUser:    id,
		Tanggal:   tanggal,
		Jawaban:   jawaban,
		SkorTotal: skor,
	}

	jumlahData++
	fmt.Println("Data berhasil ditambahkan.")
}

func SequentialSearch(id string) int {
	for i := 0; i < jumlahData; i++ {
		if dataAssessment[i].IDUser == id {
			return i
		}
	}
	return -1
}

func BinarySearch(id string) int {
	low := 0
	high := jumlahData - 1

	for low <= high {
		mid := (low + high) / 2
		if dataAssessment[mid].IDUser == id {
			return mid
		} else if dataAssessment[mid].IDUser < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func SelectionSortBySkor(ascending bool) {
	for i := 0; i < jumlahData-1; i++ {
		idx := i
		for j := i + 1; j < jumlahData; j++ {
			if ascending {
				if dataAssessment[j].SkorTotal < dataAssessment[idx].SkorTotal {
					idx = j
				}
			} else {
				if dataAssessment[j].SkorTotal > dataAssessment[idx].SkorTotal {
					idx = j
				}
			}
		}
		// Tukar posisi
		dataAssessment[i], dataAssessment[idx] = dataAssessment[idx], dataAssessment[i]
	}
}

func Tampilkan5Terakhir() {
	// Menampilkan 5 data terakhir
	fmt.Println("5 Self-Assessment Terakhir:")
	start := jumlahData - 5
	if start < 0 {
		start = 0
	}
	for i := jumlahData - 1; i >= start; i-- {
		fmt.Printf("[%s] Skor: %d\n", dataAssessment[i].IDUser, dataAssessment[i].SkorTotal)
	}
}
