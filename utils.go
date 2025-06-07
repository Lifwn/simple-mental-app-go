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
    var jawaban [10]int
    pertanyaan := [10]string{
        "1. Menjadi marah karena hal-hal kecil/sepele",
        "2. Merasa sepertinya tidak kuat lagi untuk melakukan suatu kegiatan",
        "3. Merasa banyak menghabiskan energi karena cemas",
        "4. Perubahan kegiatan jantung dan denyut nadi tanpa stimulasi oleh latihan fisik",
        "5. Takut diri terhambat oleh tugas-tugas yang tidak biasa dilakukan",
        "6. Sulit mentoleransi gangguan-gangguan terhadap hal yang sedang dilakukan",
        "7. Mudah gelisah atau ketakutan",
        "8. Sulit untuk beristirahat",
        "9. Cenderung bereaksi berlebihan pada situasi",
        "10. Kesulitan untuk tenang setelah sesuatu yang mengganggu",
    }

    fmt.Println("\nSilakan jawab pertanyaan berikut dengan skala 1–5:")
    fmt.Println("1 = Tidak Pernah, 2 = Pernah Sekali, 3 = Kadang-kadang, 4 = Sering, 5 = Selalu\n")

    for i := 0; i < 10; i++ {
        for {
            input := InputString(pertanyaan[i] + " (1-5): ")
            nilai, err := strconv.Atoi(input)
            if err == nil && nilai >= 1 && nilai <= 5 {
                jawaban[i] = nilai
                break
            }
            fmt.Println("Input harus angka 1 sampai 5.")
        }
    }

    return jawaban
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

func InterpretasiSkor(skor int) string {
    switch {
    case skor < 20:
        return "Hampir tidak ada gejala; teruskan pola hidup sehat."
    case skor <= 29:
        return "Gejala ringan; coba praktikkan teknik coping seperti relaksasi atau bercerita."
    case skor <= 39:
        return "Gangguan emosional ringan; disarankan konsultasi dengan psikolog."
    case skor <= 45:
        return "Gejala cukup intens; perlu bantuan profesional secepatnya."
    case skor <= 50:
        return "Gangguan mental sangat mungkin; segera konsultasi ke psikolog/psikiater."
    default:
        return "Skor tidak valid."
    }
}