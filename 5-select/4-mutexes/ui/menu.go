package ui

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	mainTitleColor   = color.New(color.BgHiCyan, color.Bold, color.FgBlack)
	selectTitleColor = color.New(color.FgHiMagenta)
	optionsColor     = color.New(color.FgYellow)
)

func MainMenu() string {
	clearScreen()                //clean the entail terminal screen
	mainMenuText()               // show the menu text
	return readMenuChoiceInput() //return the choice
}

func mainMenuText() {

	fmt.Println("======================================")
	mainTitleColor.Println("     Welcome to THE BOMBGAME v0.1     ")
	fmt.Println("======================================\n")
	selectTitleColor.Println("   Please select a option [1 or 2]:\n")
	optionsColor.Println("        [1]- Host a new game")
	optionsColor.Println("        [2]-  Join the game\n")
	fmt.Println("======================================")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func readMenuChoiceInput() string {
	//bufio reader etc.

	//return host|client (I did not prefer boolean vars because misunderstanding can occur over the references )
}
