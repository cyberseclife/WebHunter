package cmd

import (
	"os"

	"webhunter/internal/utils"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration and scope files",
	Long: `Initialize or check the status of include.txt and exclude.txt files.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.LogInfo("Checking scope files...")
		
		checkFile("include.txt", "# Add allowed targets here (one per line)")
		checkFile("exclude.txt", "# Add excluded targets regex here (one per line)")
		
		utils.LogOut("Configuration check complete. Use global flags --include and --exclude to specify paths if different from current directory.")
	},
}

func checkFile(name, defaultContent string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		utils.LogWarn("File %s not found. Creating default...", name)
		err := os.WriteFile(name, []byte(defaultContent+"\n"), 0644)
		if err != nil {
			utils.LogError("Error creating %s: %v", name, err)
		} else {
			utils.LogInfo("Created %s", name)
		}
	} else {
		utils.LogInfo("File %s exists.", name)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
