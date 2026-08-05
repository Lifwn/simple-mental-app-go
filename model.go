package main

import (
	"fmt"
	"time"
)

type Assessment struct {
	IDUser    string
	Tanggal   time.Time
	Jawaban   [10]int
	SkorTotal int
}

func (a *Assessment) Print(title string) {
	if title != "" {
		fmt.Println(title)
	}
	fmt.Printf("ID          : %s\n", a.IDUser)
	fmt.Printf("Tanggal     : %s\n", a.Tanggal.Format("2006-01-02"))
	fmt.Printf("Skor        : %d\n", a.SkorTotal)
	fmt.Printf("Interpretasi: %s\n", InterpretasiSkor(a.SkorTotal))
}

const MaxData = 100
var dataAssessment [MaxData]Assessment
var jumlahData int = 0
