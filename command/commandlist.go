package command

import "fmt"

func ProcessCommandInput(input string) {
	switch input {
	case "north":
		//move player up one tile
		fmt.Println("Player moved North.\n")
	case "south":
		//move player down one tile
		fmt.Println("Player moved South.\n")
	case "east":
		//move player right one tile
		fmt.Println("Player moved East.\n")
	case "west":
		//move player left one tile
		fmt.Println("Player moved West.\n")
	default:
		fmt.Println("Action Not Recognised.\n")
	}
}
