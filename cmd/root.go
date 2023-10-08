package cmd

import (
	"fmt"
	"log"
	"natas/levels"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

type NatasPassword struct {
	Username string
	Password string
}

func getDatabaseEnvironmentVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func savePasswordInDatabase(levelUsername string, levelPassword string) error {
	host := getDatabaseEnvironmentVariable("HOST")
	dbUser := getDatabaseEnvironmentVariable("USER")
	dbPassword := getDatabaseEnvironmentVariable("PASSWORD")
	dbName := getDatabaseEnvironmentVariable("DBNAME")
	port := getDatabaseEnvironmentVariable("PORT")
	sslMode := getDatabaseEnvironmentVariable("SSLMODE")
	timezone := getDatabaseEnvironmentVariable("TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, dbUser, dbPassword, dbName, port, sslMode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	level := NatasPassword{Username: levelUsername, Password: levelPassword}

	result := db.Create(&level)
	println("[+] Credentials correctly saved in the database.")
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

func getPasswordFromDatabase(username string) string {
	host := getDatabaseEnvironmentVariable("HOST")
	dbUser := getDatabaseEnvironmentVariable("USER")
	dbPassword := getDatabaseEnvironmentVariable("PASSWORD")
	dbName := getDatabaseEnvironmentVariable("DBNAME")
	port := getDatabaseEnvironmentVariable("PORT")
	sslMode := getDatabaseEnvironmentVariable("SSLMODE")
	timezone := getDatabaseEnvironmentVariable("TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, dbUser, dbPassword, dbName, port, sslMode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	credential := NatasPassword{}

	result := db.Where("username = ?", username).First(&credential)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return credential.Password
}

var rootCmd = &cobra.Command{
	Use:   "natas-cobra",
	Short: "Get natas passwords",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := cmd.Flags().GetInt("level")
		save, _ := cmd.Flags().GetBool("save")
		db, _ := cmd.Flags().GetBool("db")

		var password string
		username := fmt.Sprintf("natas%d", level)

		if db {
			password = getPasswordFromDatabase(username)
		} else {
			password = getPasswordForLevel(level)
		}

		if password == "" {
			return
		}

		fmt.Printf("User: %s\n", username)
		fmt.Printf("Password: %s\n", password)

		if save {
			err := savePasswordInDatabase(username, password)
			if err != nil {
				log.Fatal(err)
			}
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
	rootCmd.Flags().Bool("save", false, "save the password of the level in the database")
	rootCmd.Flags().Bool("db", false, "obtain the passwords from the database")
}
