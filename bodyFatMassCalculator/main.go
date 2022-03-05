package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const toInches = 0.393701
const toPounds = 2.20462

func main() {

	var fatPercentage float64
	fmt.Println("Fat mass Calculator !")
	fmt.Println("Uses US army formula")

	// are u men or woman
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you a Men (M) or a Female(F) ")
	rawGender, _ := reader.ReadString('\n')

	// validate gender
	gender := cleanInput(rawGender)
	if gender != "M" && gender != "F" {
		panic("invalid gender")
	}

	// ask for height
	fmt.Println("What is your height (m) ? (ex. 1.84)")
	rawHeight, _ := reader.ReadString('\n')
	heigth := convertQty(cleanInput(rawHeight), "heigth")
	heigth = convertToInches(heigth * 100)

	// fmt.Println("heigth ", heigth)

	// ask for neck measurment
	fmt.Println("What is your neck circunference (cm) ? ")
	rawNeck, _ := reader.ReadString('\n')
	neck := convertQty(cleanInput(rawNeck), "neck")
	neck = convertToInches(neck)

	// fmt.Println("neck", neck)

	// ask for navel measurement
	fmt.Println("What is your navel(stomach) circunference (cm) ? ")
	rawNavel, _ := reader.ReadString('\n')
	navel := convertQty(cleanInput(rawNavel), "navel")
	navel = convertToInches(navel)

	// fmt.Println("navel", navel)

	// ask for waist measurement if woman
	if gender == "F" {
		fmt.Println("What is your waist circunference (cm) ? ")
		rawWaist, _ := reader.ReadString('\n')
		waist := convertQty(cleanInput(rawWaist), "waist")
		// fmt.Println("waist", waist)
		waist = convertToInches(waist)
		fatPercentage = calculateFemale(heigth, neck, navel, waist)
	} else {
		fatPercentage = calculateMale(heigth, neck, navel)
	}
	fmt.Println("Yor approximate body fat mass is: ", fatPercentage)
}

func convertMassToEnglish(massInKilograms float64) float64 {
	return massInKilograms * toPounds
}

func convertToInches(lenghtInCentimeters float64) float64 {
	return lenghtInCentimeters * toInches
}

//men = [86.010 * Log10(Cintura – Cuello)] – [70.041 * Log10(Altura)] + 36.76
func calculateMale(heigth float64, neck float64, navel float64) float64 {
	return (86.010 * math.Log10(navel-neck)) - (70.041 * math.Log10(heigth)) + 36.76
}

// woman = [163.205 * Log10(Cintura + Cadera – Cuello)] – [97.684 * Log10(Altura)] – 78.387
func calculateFemale(heigth float64, neck float64, navel float64, waist float64) float64 {
	return (163.205 * math.Log10(navel+waist-neck)) - (97.684 * math.Log10(heigth)) - 78.387
}

func cleanInput(input string) string {
	return strings.TrimRight(input, "\n")
}

func convertQty(input string, paramName string) float64 {
	converted, _ := strconv.ParseFloat(input, 8)
	return converted
}
