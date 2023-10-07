package cmd

import (
	"fmt"
	"natas/levels"
	"os"

	"github.com/spf13/cobra"
)

func getPasswordForLevel(level int) string {
	switch level {
	case 0:
		return "natas0"
	case 1:
		return levels.GetLevel1Password()
	case 2:
		level1Password := levels.GetLevel1Password()
		return levels.GetLevel2Password(level1Password)
	case 3:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		return levels.GetLevel3Password(level2Password)
	case 4:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		return levels.GetLevel4Password(level3Password)
	case 5:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		return levels.GetLevel5Password(level4Password)
	case 6:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		return levels.GetLevel6Password(level5Password)
	case 7:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		level6Password := levels.GetLevel6Password(level5Password)
		return levels.GetLevel7Password(level6Password)
	case 8:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		level6Password := levels.GetLevel6Password(level5Password)
		level7Password := levels.GetLevel7Password(level6Password)
		return levels.GetLevel8Password(level7Password)
	case 9:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		level6Password := levels.GetLevel6Password(level5Password)
		level7Password := levels.GetLevel7Password(level6Password)
		level8Password := levels.GetLevel8Password(level7Password)
		return levels.GetLevel9Password(level8Password)
	case 10:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		level6Password := levels.GetLevel6Password(level5Password)
		level7Password := levels.GetLevel7Password(level6Password)
		level8Password := levels.GetLevel8Password(level7Password)
		level9Password := levels.GetLevel9Password(level8Password)
		return levels.GetLevel10Password(level9Password)
	case 11:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		level6Password := levels.GetLevel6Password(level5Password)
		level7Password := levels.GetLevel7Password(level6Password)
		level8Password := levels.GetLevel8Password(level7Password)
		level9Password := levels.GetLevel9Password(level8Password)
		level10Password := levels.GetLevel10Password(level9Password)
		return levels.GetLevel11Password(level10Password)
	case 12:
		level1Password := levels.GetLevel1Password()
		level2Password := levels.GetLevel2Password(level1Password)
		level3Password := levels.GetLevel3Password(level2Password)
		level4Password := levels.GetLevel4Password(level3Password)
		level5Password := levels.GetLevel5Password(level4Password)
		level6Password := levels.GetLevel6Password(level5Password)
		level7Password := levels.GetLevel7Password(level6Password)
		level8Password := levels.GetLevel8Password(level7Password)
		level9Password := levels.GetLevel9Password(level8Password)
		level10Password := levels.GetLevel10Password(level9Password)
		level11Password := levels.GetLevel11Password(level10Password)
		return levels.GetLevel12Password(level11Password)
	default:
		return ""
	}
}

var rootCmd = &cobra.Command{
	Use:   "natas-cobra",
	Short: "Get natas passwords",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := cmd.Flags().GetInt("level")
		password := getPasswordForLevel(level)
		if password == "" {
			fmt.Println("Password not found")
			return
		}
		fmt.Printf("User: natas%d\n", level)
		fmt.Printf("Password: %s\n", password)
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
