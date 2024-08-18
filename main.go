package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silahkan masukkan nama teman yang ingin dicari..")
		return
	}

	nama := os.Args[1]
	cariTeman(nama)
}

var temanKelas = []Biodata{
	{"1", "Jeff", "Medan", "Frontend Developer", "Iseng"},
	{"2", "Josh", "Bandung", "Software Engineer", "Gabuts"},
	{"3", "Reza", "Surabaya", "Mobile Dev Engineer", "Uang"},
	{"4", "Gasba", "Jakarta", "Data Analyst", "Buat pamer skill"},
}

type Biodata struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func cariTeman(nama string) {
	nama = strings.ToLower(nama)

	for _, teman := range temanKelas {
		if strings.HasPrefix(strings.ToLower(teman.Nama), nama) {
			fmt.Printf("Id: %s\n", teman.Id)
			fmt.Printf("Nama: %s\n", teman.Nama)
			fmt.Printf("Alamat: %s\n", teman.Alamat)
			fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
			fmt.Printf("Alasan: %s\n", teman.Alasan)
			return
		}
	}
	fmt.Println("Teman yang anda cari sudah resign 😢")
}
