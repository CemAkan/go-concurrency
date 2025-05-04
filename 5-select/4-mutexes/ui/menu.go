package ui

import (
	"bombgame/conf"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"
)

var (
	mainTitleColor   = color.New(color.BgHiCyan, color.Bold, color.FgBlack)
	selectTitleColor = color.New(color.FgHiMagenta)
	optionsColor     = color.New(color.FgYellow)
	choiceWarning    = color.New(color.BgRed, color.Bold)
)

func MainMenu() {
	for {
		clearScreen()                        //clean the entail terminal screen
		mainMenuText()                       // show the menu text
		inputChoice := readMenuChoiceInput() //return the choice

		if inputChoice != "invalid" {
			conf.PlayerStatus = inputChoice // assorting input data to config's project based global var
			return
		}
	}
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

	switch reader() {
	case "1":
		return "host"
	case "2":
		return "client"
	default:
		clearScreen()
		inputWarningMessage()
		time.Sleep(time.Second * 3)
		return "invalid"
	}
	//return host|client (I did not prefer boolean vars because misunderstanding can occur over the references )

}

func inputWarningMessage() {
	fmt.Println("======================================")
	choiceWarning.Println("    Please select from only 1 or 2    ")
	fmt.Println("======================================")
}

func reader() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	return strings.TrimSpace(line)
}
