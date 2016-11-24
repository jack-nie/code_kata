// Description
//
// To disarm the bomb you have to cut some wires. These wires are either white, black, purple, red, green or orange.
// The rules for disarming are simple:
// If you cut a white cable you can't cut white or black cable.
// If you cut a red cable you have to cut a green one
// If you cut a black cable it is not allowed to cut a white, green or orange one
// If you cut a orange cable you should cut a red or black one
// If you cut a green one you have to cut a orange or white one
// If you cut a purple cable you can't cut a purple, green, orange or white cable
// If you have anything wrong in the wrong order, the bomb will explode.
// There can be multiple wires with the same colour and these instructions are for one wire at a time. Once you cut a wire you can forget about the previous ones.
// Formal Inputs & Outputs
//
// Input description
//
// You will recieve a sequence of wires that where cut in that order and you have to determine if the person was succesfull in disarming the bomb or that it blew up.
// Input 1
//
// white
// red
// green
// white
// Input 2
//
// white
// orange
// green
// white
// Output description
//
// Wheter or not the bomb exploded
// Output 1
//
// "Bomb defused"
// Output 2
//
// "Boom"

package main

import "fmt"

func disarm(bomb []string) {
	arr := bomb
	var lastColor string

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			lastColor = arr[i]
			continue
		}

		if lastColor == "white" && (arr[i] == "black" || arr[i] == "white") {
			fmt.Println("Boom")
			return
		}

		if lastColor == "red" && arr[i] != "green" {
			fmt.Println("Boom")
			return
		}

		if lastColor == "black" && (arr[i] == "white" || arr[i] == "green" || arr[i] == "orange") {
			fmt.Println("Boom")
			return
		}

		if lastColor == "orange" && (arr[i] != "red" && arr[i] != "black") {
			fmt.Println("Boom")
			return
		}

		if lastColor == "green" && (arr[i] != "orange" && arr[i] != "white") {
			fmt.Println("Boom")
			return
		}

		if lastColor == "purple" && (arr[i] == "purple" || arr[i] == "green" || arr[i] == "orange" || arr[i] == "white") {
			fmt.Println("Boom")
			return
		}

		lastColor = arr[i]
	}

	fmt.Println("Bomb defused")

}

func main() {
	arr := []string{"white", "red", "green", "white"}

	disarm(arr)
	bomb := []string{"white", "orange", "green", "white"}
	disarm(bomb)
}
