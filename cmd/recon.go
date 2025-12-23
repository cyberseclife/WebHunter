package cmd

import (
	"webhunter/internal/scanner"
	"webhunter/internal/utils"

	"github.com/spf13/cobra"
)

var (
	reconSv     bool
	reconSs     bool
	reconPorts  string
)

var reconCmd = &cobra.Command{
	Use:   "recon [targets]",
	Short: "Perform reconnaissance and port scanning",
	Long: `Run network reconnaissance, including port scanning and service version detection.
Example: webhunter recon example.com -p 80,443 -sV`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := scanner.Config{
			RateLimit: globalRate,
			Header:    globalHeader,
		}
		sc := scanner.New(cfg)
		
		utils.LogInfo("Starting Recon on %d targets", len(args))
		sc.Recon(args, reconPorts, reconSv, reconSs)
	},
}

func init() {
	rootCmd.AddCommand(reconCmd)
	reconCmd.Flags().BoolVarP(&reconSv, "service-version", "s", false, "Service Version Detection (-sV)")
	reconCmd.Flags().BoolVarP(&reconSs, "stealth", "S", false, "Stealth Scan (-sS)")
	reconCmd.Flags().StringVarP(&reconPorts, "ports", "p", "", "Ports to scan (e.g. 80,443 or - for all)")
}
