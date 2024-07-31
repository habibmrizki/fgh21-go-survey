package main

import "fmt"

type survey struct{
	name string
	age int
	gender string
	isCigar bool
	// cigarVariant []string
}

func main() {

	var name string
	fmt.Println("Siapa nama anda?: ")
	fmt.Scanln(&name)
	fmt.Println("Berapa umur anda?: ")
	var age int
	fmt.Scanln(&age) 
	fmt.Println("Apa jenis kelamin anda?: 1. Laki-Laki 2. Perempuan")
	var genderChoice int
	fmt.Scanln(&genderChoice)

	var gender string
	if genderChoice == 1 {
		gender = "Laki-Laki"
	} else if genderChoice == 2 {
		gender = "Perempuan"
	} else {
		fmt.Println("Pilihan tidak valid. Harap masukkan 1 atau 2.")
		return
	}

	fmt.Println("Apakah anda perokok? (y/n): ")
	var isCigarChoice string
	fmt.Scanln(&isCigarChoice)

	var isCigar bool
	if isCigarChoice == "y" {
		isCigar = true
	} else if isCigarChoice == "n" {
		isCigar = false
	} else {
		fmt.Println("Pilihan tidak valid. Harap masukkan y atau n.")
		return
	}


	surveyResult := survey{
		name:         name,
		age:          age,
		gender:       gender,
		isCigar:      isCigar,
	}

	fmt.Printf("Hasil Survey: %+v\n", surveyResult)
}


