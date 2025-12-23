package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

// PrintBanner displays the WEBHUNTER ascii art in rainbow colors
func PrintBanner() {
	banner := `
 _       __     __    __  __            __
| |     / /__  / /_  / / / /_  ______  / /____  _____ 
| | /| / / _ \/ __ \/ /_/ / / / / __ \/ __/ _ \/ ___/
| |/ |/ /  __/ /_/ / __  / /_/ / / / / /_/  __/ /
|__/|__/_\___/_.___/_/ /_/\__,_/_/ /_/\__/ _\___/_/
`
	colors := []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
	}

	rand.Seed(time.Now().UnixNano())
	
	lines := strings.Split(banner, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		// Cycle through colors
		c := color.New(colors[i%len(colors)], color.Bold)
		c.Println(line)
	}

	fmt.Println("    High-Performance Security Assessment CLI")
	fmt.Println()
}

// LogInfo logs an informational message
func LogInfo(format string, args ...interface{}) {
	color.Green("[INFO] "+format, args...)
}

// LogWarn logs a warning message
func LogWarn(format string, args ...interface{}) {
	color.Yellow("[WARN] "+format, args...)
}

// LogError logs an error message
func LogError(format string, args ...interface{}) {
	color.Red("[ERR] "+format, args...)
}

// LogOut logs a generic output message
func LogOut(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}