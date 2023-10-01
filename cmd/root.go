package cmd

import (
	"fmt"
	"natas/levels"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "natas-cobra",
	Short: "Get natas passwords",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := cmd.Flags().GetInt("level")
		if level == 0 {
			println("User: natas0")
			println("Password: natas0")

		} else if level == 1 {
			var level1Password = levels.GetLevel1Password()
			println("User: natas1")
			fmt.Printf("Password: %s\n", level1Password)

		} else if level == 2 {
			var level1Password = levels.GetLevel1Password()
			var level2Password = levels.GetLevel2Password(level1Password)
			println("User: natas2")
			fmt.Printf("Password: %s\n", level2Password)

		} else if level == 3 {
			var level1Password = levels.GetLevel1Password()
			var level2Password = levels.GetLevel2Password(level1Password)
			var level3Password = levels.GetLevel3Password(level2Password)
			println("User: natas3")
			fmt.Printf("Password: %s\n", level3Password)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var level int
	rootCmd.Flags().IntVarP(&level, "level", "l", 1, "the number of the level to get the password")
}