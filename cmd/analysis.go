package cmd

import (
	"webhunter/internal/scanner"
	"webhunter/internal/utils"

	"github.com/spf13/cobra"
)

var (
	analVuln   bool
	analFuzz   bool
	analHeader []string
)

var analysisCmd = &cobra.Command{
	Use:   "analysis [targets]",
	Short: "Analyze targets for vulnerabilities",
	Long: `Perform vulnerability scanning and directory fuzzing on targets.
Example: webhunter analysis example.com --vuln-scan --fuzz`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := scanner.Config{
			RateLimit: globalRate,
			Header:    globalHeader,
		}
		sc := scanner.New(cfg)
		
		utils.LogInfo("Starting Analysis on %d targets", len(args))
		sc.Analysis(args, analVuln, analFuzz, analHeader)
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)
	analysisCmd.Flags().BoolVar(&analVuln, "vuln-scan", false, "Enable vulnerability scanning")
	analysisCmd.Flags().BoolVar(&analFuzz, "fuzz", false, "Enable directory brute-force/fuzzing")
	analysisCmd.Flags().StringSliceVar(&analHeader, "headers", []string{}, "Custom headers for analysis")
}
