package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/*
Test command for cli
 */
var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "This is just a test for the cli commands",
	Long: `A Fast and Flexible CLI built with Complete documentation available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("It's working!\n")
	},
}

func openDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", "../bd_sqlite3/ubiwhereData.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS machine_data (id INTEGER PRIMARY KEY, cpu FLOAT, ram FLOAT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS simu_device (id INTEGER PRIMARY KEY, d1 INT, d2 INT,d3 INT, d4 INT)")
	statement.Exec()
	return database
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}