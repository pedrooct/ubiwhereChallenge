package main

import(
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"strconv"
	"strings"
)
func openDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", "../bd_sqlite3/ubiwhereData.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS machine_data (id INTEGER PRIMARY KEY, cpu FLOAT, ram FLOAT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS simu_device (id INTEGER PRIMARY KEY, d1 INT, d2 INT,d3 INT, d4 INT)")
	statement.Exec()
	return database
}
/*
I didn't implemented json because of time. Response as string.
 */
func main(){
	r := gin.Default()
	r.GET("/get/last/:n", func (c* gin.Context){
		i, err := strconv.Atoi(c.Param("n"))
		if err !=nil || i < 0 && i > 10 {
			c.JSON(400,"Error: Invalid parameter")
			return
		}
		var result string
		var d1 int
		var d2 int
		var d3 int
		var d4 int
		db:=openDatabase()
		rows, _ := db.Query("SELECT d1,d2,d3,d4 FROM simu_device ORDER BY id DESC limit ?;",i)
		for rows.Next(){
			rows.Scan(&d1,&d2,&d3,&d4)
			result = result + fmt.Sprintf("[%s,%s,%s,%s]",strconv.Itoa(d1),strconv.Itoa(d2),strconv.Itoa(d3),strconv.Itoa(d4))
		}
		db.Close()
		c.JSON(200,result)
	})
	r.GET("/get/lastv/:n" , func (c* gin.Context){
		i, err := strconv.Atoi(c.Param("n"))
		if err !=nil || i < 0 && i > 10 {
			c.JSON(400,"Error: Invalid parameter")
			return
		}
		body, err := ioutil.ReadAll(c.Request.Body)
		args:= strings.Split(string(body),",")
		db:= openDatabase()
		var result string
		var val int
		for v :=range args{
			if args[v] == "d1" ||args[v] == "d2" ||args[v] == "d3" || args[v] == "d4" {
				q := "SELECT " + args[v] + " FROM simu_device ORDER BY id DESC limit ?;"
				rows, err := db.Query(q, i)
				if err != nil {
					c.JSON(400,"Error: Query malformed")
					return
				}
				result = result + fmt.Sprintf("[%s:", args[v])
				for rows.Next() {
					rows.Scan(&val)
					result = result + fmt.Sprintf("%d,", val)
				}
				result = result + fmt.Sprintf("] ")
			}
		}
		db.Close()
		c.JSON(200,result)
	})
	r.GET("/get/avg" , func (c* gin.Context){
		body, err := ioutil.ReadAll(c.Request.Body)
		if err!=nil {
			c.JSON(400, "Error: Body malformed")
			return
		}
		args:= strings.Split(string(body),",")
		db:= openDatabase()
		var result string
		var avg float64
		for v :=range args{
			if args[v] == "d1" ||args[v] == "d2" ||args[v] == "d3" || args[v] == "d4" {
				rows, err := db.Query("SELECT avg("+args[v]+") as avg_val FROM simu_device;")
				if err != nil {
					c.JSON(400,"Error: Query malformed")
					return
				}
				result = result + fmt.Sprintf("[Avg %s:",args[v])
				for rows.Next(){
					rows.Scan(&avg)
					result = result + fmt.Sprintf("%f] ",avg)
				}
			}
		}
		db.Close()
		c.JSON(200,result)
	})
	r.Run(":8008")
}
