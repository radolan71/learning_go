package main

import (
	"fmt"
	"sort"
)

func main() {
	// pointers()
	// arrays()
	// slices()
	// maps()
	structs()
}

func pointers() {

	//initialized with nil
	var pointer *int
	randomInt := 71

	//assign pointer to the randomInt adress in memory
	pointer = &randomInt
	fmt.Println("randomInt", *pointer)

	//increase pointer
	*pointer++

	//original value and pointer value shoud have been incrementer
	fmt.Println("pointer", *pointer)
	fmt.Println("randomInt", randomInt)

}

func arrays() {
	var names [3]string
	names[0] = "Kim"
	names[1] = "Xavi"
	names[2] = "Pam"

	var familyNames = [3]string{"Macdaniels", "Alonso", "Wright"}

	fmt.Println(names)
	fmt.Println(len(names))

	fmt.Println(familyNames)
	fmt.Println(len(familyNames))

	for index, element := range names {
		fmt.Println("Full Name index", index, "is", element, familyNames[index])
	}
}

// abstraction layer of arrays , rezisible and can be sorted quite easily
func slices() {

	//another option is
	// nflTeams := make([]string)
	var nflTeams = []string{"Jaguars", "Dolphins", "Packers", "Patriots", "Cowboys"}
	fmt.Println(nflTeams)

	//add item
	nflTeams = append(nflTeams, "Falcons")
	fmt.Println(nflTeams)

	//delete first item
	nflTeams = append(nflTeams[1:len(nflTeams)])
	fmt.Println(nflTeams)

	//delete last item
	nflTeams = append(nflTeams[:len(nflTeams)-1])
	fmt.Println(nflTeams)

	//sort
	sort.Strings(nflTeams)
	fmt.Println(nflTeams)
}

//unordered collections
func maps() {
	currencies := make(map[string]string)
	fmt.Println(currencies)
	currencies["USD"] = "$"
	currencies["EUR"] = "€"
	currencies["POUND"] = "£"
	currencies["CORONA"] = "kr"
	currencies["RUPEE"] = "₹"

	fmt.Println(currencies)
	delete(currencies, "CORONA")

	fmt.Println(currencies)

	fmt.Println("unsorted")
	for name, symbol := range currencies {
		fmt.Printf("%v: %v \n", name, symbol)
	}

	//warranty order
	keys := make([]string, len(currencies))
	i := 0
	for l := range currencies {
		keys[i] = l
		i++
	}
	sort.Strings(keys)

	fmt.Println("sorted")
	for k := range keys {
		fmt.Printf("%v: %v \n", keys[k], currencies[keys[k]])
	}
}

// there is no inheritance
func structs() {
	player := Player{"Zack Thomas", 54, "Texas Tech", false}
	fmt.Println(player)

	fmt.Printf("%+v\n", player)
	fmt.Printf("Name %v Number %v University %v \n ", player.Name, player.Number, player.University)

	//change number
	player.Number = 55
	fmt.Printf("%+v\n", player)

	//retire
	player.Retire()
	fmt.Printf("%+v\n", player)
}

type Player struct {
	Name       string
	Number     int
	University string
	Retired    bool
}

func (p *Player) Retire() {
	p.Retired = true
}
