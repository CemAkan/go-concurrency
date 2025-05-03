package ui

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	titleColor = color.New(color.BgHiCyan, color.Bold, color.FgBlack)
)

func ShowMainMenu() {
	fmt.Println("======================================")
	titleColor.Println("      Welcome to THE BOMBGAME v0.1    ")
	fmt.Println("======================================\n")
	fmt.Println("    Please select a option [1 or 2]:\n")
	fmt.Println("          1- Host a new game")
	fmt.Println("           2- Join the game")
	fmt.Println(" ======================================")
}
