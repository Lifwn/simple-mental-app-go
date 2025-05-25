package main

import "time"


type Assessment struct {
	IDUser      string
	Tanggal     time.Time
	Jawaban     [10]int 
	SkorTotal   int
}


const MaxData = 100
var dataAssessment [MaxData]Assessment
var jumlahData int = 0
