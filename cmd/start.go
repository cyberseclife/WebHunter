package cmd

import (
	"webhunter/internal/scanner"
	"webhunter/internal/utils"

	"github.com/spf13/cobra"
)

var (
	autoMode bool
)

var startCmd = &cobra.Command{
	Use:   "start [targets]",
	Short: "Start the master assessment workflow",
	Long: `Run the full pipeline: Recon -> Analysis -> Exploitation.
Use --auto to enable all stages with default settings.`, 
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !autoMode {
			utils.LogWarn("Please use --auto to confirm automatic full-chain execution.")
			return
		}

		cfg := scanner.Config{
			RateLimit: globalRate,
			Header:    globalHeader,
		}
		sc := scanner.New(cfg)
		
		utils.LogInfo(">>> AUTOMATIC MODE ENGAGED <<<")
		
		// 1. Recon
		utils.LogOut("\n--- Stage 1: Recon ---")
		// Defaulting to -p- logic or standard ports for auto mode
		activeTargets := sc.Recon(args, "", true, true)
		
		if len(activeTargets) == 0 {
			utils.LogWarn("No active targets found. Stopping.")
			return
		}

		// 2. Analysis
		utils.LogOut("\n--- Stage 2: Analysis ---")
		sc.Analysis(activeTargets, true, true, nil)

		// 3. Exploitation
		utils.LogOut("\n--- Stage 3: Exploitation ---")
		// Using a default set of test payloads for auto mode
		sc.Exploitation(activeTargets, []string{"test_payload_1", "sqli_test"})

		utils.LogInfo("WebHunter run complete.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVar(&autoMode, "auto", false, "Run all stages sequentially")
}
