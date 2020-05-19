package cmd

import (

	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/mattn/go-sqlite3"

)


func init() {
	rootCmd.AddCommand(getavg)
}


var getavg = &cobra.Command{
	Use:   "getAvg",
	Short: "Get avg of a value for one or more variables",
	Long: `Use d as define variable: EX getlastv d1 d2 d3 ... `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("requires at least one value 0-4\n")
			return
		}
		db:= openDatabase()
		var avg float64
		for v:=range args{
			rows, err := db.Query("SELECT avg("+args[v]+") as avg_val FROM simu_device;")
			if err != nil{
				fmt.Printf("Error: query malformed !\n")
				return
			}
			fmt.Printf("Avg %s:",args[v])
			for rows.Next(){
				rows.Scan(&avg)
				fmt.Printf("%f\n",avg)
			}
		}
		db.Close()
	},
}
