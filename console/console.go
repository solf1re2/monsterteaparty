package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func StartConsole() {
    prompt := "What would you like to do?"
    // Start a loop, prompting for input from the user for commands.
    // Have quit() as an escape command.
    endMessage := loopForCommands(prompt)
    fmt.Println(endMessage)

}

func loopForCommands(promptMessage string) string {
    var exit = false
    for !exit {
        response := PromptAndReturnInputFromUser(promptMessage)
        processResponse(response, &exit)
    }
    return "Exiting, Good Bye.\n"
}

// processResponse - this will have general commands such as quit and help.
// It will pass other input to a specialised command package designed for the program.
func processResponse(r string, exit *bool) {
    switch r {
    case "quit":
        fmt.Println("You have quit.\n")
        *exit = true
    case "help":
        fmt.Println("You'll find no help here.\n")
    default:
        fmt.Println("only commands available are: 'help', and 'quit'")
    }
}


/*
Function called with input string used to prompt user for a single line input.
@promptMessage message to prompt user for input.
@returns user input
*/
func PromptAndReturnInputFromUser(promptMessage string) string {
	var i string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(promptMessage)
	for scanner.Scan() {
		// assign the input to i
		i = scanner.Text()

		// Check that there is input
		if len(i) > 0 {
			break
		} else {
			fmt.Println(promptMessage)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Print(os.Stderr, "reading standard input:", err)
	}
    fmt.Printf("\n")

	return i
}
