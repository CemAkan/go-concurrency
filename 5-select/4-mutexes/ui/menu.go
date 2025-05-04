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
	mainTitleColor        = color.New(color.BgHiCyan, color.Bold, color.FgBlack)
	selectTitleColor      = color.New(color.FgHiMagenta)
	optionsColor          = color.New(color.FgYellow)
	choiceWarningColor    = color.New(color.BgRed, color.Bold)
	askSettingsTitleColor = color.New(color.BgWhite, color.FgBlack)
	infoTitleColor        = color.New(color.BgWhite, color.FgHiYellow)
)

func UserInputStart() {
	mainMenu()
	askName()

	if conf.PlayerStatus == "client" {
		askIP()
	}
}

func mainMenu() {
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
		showWarningMessage("Please select from only 1 or 2")
		return "invalid"
	}
	//return host|client (I did not prefer boolean vars because misunderstanding can occur over the references )

}

func showWarningMessage(msg string) {
	clearScreen()
	fmt.Println("======================================")
	choiceWarningColor.Println(msg)
	fmt.Println("======================================")

	time.Sleep(time.Second * 3)
}

func reader() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	return strings.TrimSpace(line)
}

func askName() {
	for {
		clearScreen()
		askSettingsTitleColor.Print("Please, write a name: ")
		nameInput := reader()

		if len(nameInput) < 2 {
			showWarningMessage("Name too short. Try again please")
		}
		conf.PlayerName = nameInput
		return
	}
}

func askIP() {
	clearScreen()
	askSettingsTitleColor.Print("Please write host [ IP:PORT ] to connect: ")

	conf.GameAddress = reader()
}

func HostInfoShowMenu(ip string, port int) {
	clearScreen()
	fmt.Println("======================================")
	fmt.Print("Game address is: ")
	infoTitleColor.Println(ip, ":", port)
	fmt.Println("======================================")
}
