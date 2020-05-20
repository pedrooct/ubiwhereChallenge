package cmd

import (

	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	_ "github.com/mattn/go-sqlite3"

)


func init() {
	rootCmd.AddCommand(getlastv)
}
/*
This command allows to fetch of the last n metrics for one or more variables.
*/
var getlastv = &cobra.Command{
	Use:   "getLastV",
	Short: "Get last n metrics for one or more variables",
	Long: `Use d as define variable: EX getlastv 4 d1 d2 d3 ... `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("Requires at least one value 0-4\n")
			return
		}
		i, err := strconv.Atoi(args[0])
		if err !=nil || i < 0 && i > 10 {
			fmt.Printf("Invalid value\n")
			return
		}
		db:= openDatabase()
		var val int
		for v :=range args{
			if v!=0{
				q:="SELECT "+args[v]+" FROM simu_device ORDER BY id DESC limit ?;" // For some reason i had to join strings
				rows, err := db.Query(q,i)
				if err != nil{
					fmt.Printf("Error: query malformed\n")
				}
				fmt.Printf("%s:\n",args[v])
				for rows.Next(){
					rows.Scan(&val)
					fmt.Printf("%d|",val)
				}
				fmt.Printf("\n")
			}
		}

		db.Close()
	},
}
