package main
import "fmt"

func hitungSkor(soal, skor *int) {
    var waktu int
    *soal = 0
    *skor = 0

    for i := 0; i < 8; i++ {
        fmt.Scan(&waktu)
        if waktu != 301 {
            *soal++
            *skor += waktu
        }
    }
}

func main() {
    var nama, pemenang string
    var soal, skor int
    var maxSoal = -1
    var minSkor = 1000000

    for {
        fmt.Scan(&nama)
        if nama == "Selesai" {
            break
        }

        hitungSkor(&soal, &skor)

        if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
            maxSoal = soal
            minSkor = skor
            pemenang = nama
        }
    }

    fmt.Println(pemenang, maxSoal, minSkor)
}