package cmd

import (
	"fmt"
	"os"

	"webhunter/internal/utils"

	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	globalRate  int
	globalHeader string
	includeFile string
	excludeFile string
)

var rootCmd = &cobra.Command{
	Use:   "webhunter",
	Short: "WebHunter is a high-performance security assessment tool",
	Long: `A modular CLI application for Recon, Analysis, and Exploitation.
Automates the security assessment pipeline.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Default to local files if flags are empty and files exist
		if includeFile == "" {
			if _, err := os.Stat("include.txt"); err == nil {
				includeFile = "include.txt"
			}
		}
		if excludeFile == "" {
			if _, err := os.Stat("exclude.txt"); err == nil {
				excludeFile = "exclude.txt"
			}
		}

		err := utils.InitScope(includeFile, excludeFile)
		if err != nil {
			utils.LogError("Failed to initialize scope: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVar(&globalRate, "rate", 5, "Rate limit in requests per second")
	rootCmd.PersistentFlags().StringVar(&globalHeader, "header", "", "Custom header to include in requests")
	rootCmd.PersistentFlags().StringVar(&includeFile, "include", "", "File containing allowed targets (Scope)")
	rootCmd.PersistentFlags().StringVar(&excludeFile, "exclude", "", "File containing excluded targets (Out of Scope)")
}
