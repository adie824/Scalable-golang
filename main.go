package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (biodata *Biodata) LihatAlasan() string {
	return biodata.Nama + " memilih kelas Golang karena " + biodata.Alasan
}

func main() {
	biodata := []Biodata{
		{
			Nama:      "Adi",
			Alamat:    "Surabaya",
			Pekerjaan: "IT Suport",
			Alasan:    "Belajar Bahasa Pemograman Baru",
		},
		{
			Nama:      "Agil",
			Alamat:    "Kediri",
			Pekerjaan: "Web Development",
			Alasan:    "ingin meningkatkan bahasa pemograman yang baru",
		},
	}

	arg, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Nama : ", biodata[arg].Nama)
	fmt.Println("Alamat : ", biodata[arg].Alamat)
	fmt.Println("Pekerjaan : ", biodata[arg].Pekerjaan)
	fmt.Println("Alasan : ", biodata[arg].LihatAlasan())
}
