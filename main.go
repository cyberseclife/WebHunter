package main

import (
	"webhunter/cmd"
	"webhunter/internal/utils"
)

func main() {
	utils.PrintBanner()
	cmd.Execute()
}