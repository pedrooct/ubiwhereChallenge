package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	_ "github.com/mattn/go-sqlite3"

)


func init() {
	rootCmd.AddCommand(getlast)
}

/*
This command allows to fetch of the last n metrics for all variables.
*/
var getlast = &cobra.Command{
	Use:   "getLast",
	Short: "Get last n metrics for all variables",
	Long: `A Fast and Flexible CLI built with Complete documentation available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("requires at least one value 0-4\n")
			return
		}
		i, err := strconv.Atoi(args[0])
		if err !=nil || i < 0 && i > 10 {
			fmt.Printf("Invalid value\n")
			return
		}
		var d1 int
		var d2 int
		var d3 int
		var d4 int
		db:= openDatabase()
		rows, _ := db.Query("SELECT d1,d2,d3,d4 FROM simu_device ORDER BY id DESC limit ?;",i) // limit to N values
		for rows.Next(){
			rows.Scan(&d1,&d2,&d3,&d4)
			fmt.Printf("%d,%d,%d,%d\n",d1,d2,d3,d4)
		}
		db.Close()
	},
}

