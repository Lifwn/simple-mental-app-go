package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func InputString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func InputJawaban() [10]int {
	var j [10]int
	for i := 0; i < 10; i++ {
		for {
			input := InputString(fmt.Sprintf("Pertanyaan %d (1-5): ", i+1))
			n, err := strconv.Atoi(input)
			if err == nil && n >= 1 && n <= 5 {
				j[i] = n
				break
			}
			fmt.Println("Input tidak valid.")
		}
	}
	return j
}

func InputTanggal() time.Time {
	for {
		input := InputString("Tanggal (YYYY-MM-DD): ")
		t, err := time.Parse("2006-01-02", input)
		if err == nil {
			return t
		}
		fmt.Println("Format salah.")
	}
}
