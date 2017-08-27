package main

import "strconv"

func main() {
	for i := 99; i > 0; i-- {
		b := bottles(i) + " of beer"

		println(b + " on the wall, " + b + ".")
		println("Take one down and pass it around, " + bottles(i-1) + " of beer on the wall.\n")
	}

	println("No more bottles of beer on the wall, no more bottles of beer.")
	println("Go to the store and buy some more, 99 bottles of beer on the wall.")
}

func bottles(beer int) string {
	switch beer {
	case 0:
		return "no bottles"
	case 1:
		return "1 bottle"
	}

	return strconv.Itoa(beer) + " bottles"
}
