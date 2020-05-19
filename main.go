package main

import (
	"database/sql"
	"math/rand"
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"os"
	"time"
	_ "github.com/mattn/go-sqlite3"

)

func openDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", "./bd_sqlite3/ubiwhereData.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS machine_data (id INTEGER PRIMARY KEY, cpu FLOAT, ram FLOAT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS simu_device (id INTEGER PRIMARY KEY, d1 INT, d2 INT,d3 INT, d4 INT)")
	statement.Exec()
	return database
}


func simulateDevice(db *sql.DB){
	var arr  [4]int
	for {
		for i:=0 ; i< 4;i++{
			arr[i]= rand.Intn(99)
			//fmt.Printf("%d|",arr[i])
		}
		statement, _ := db.Prepare("INSERT INTO simu_device (d1,d2,d3,d4) VALUES (?,?,?,?)")
		statement.Exec(arr[0],arr[1],arr[2],arr[3])
		fmt.Printf("%d,%d,%d,%d\n",arr[0],arr[1],arr[2],arr[3])
		time.Sleep(1*time.Second)
	}


}

func getCpuRam(db *sql.DB ) {
	for true{
		before, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}

		total := float64(after.Total - before.Total)
		mem, err := memory.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		statement, _ := db.Prepare("INSERT INTO machine_data (cpu, ram) VALUES (?, ?)")
		statement.Exec(100-float64(after.Idle-before.Idle)/total*100,float64(mem.Used)/1000000000)
		fmt.Printf("CPU: %0.2f %% , RAM: %0.2f Gb \n",100-float64(after.Idle-before.Idle)/total*100,float64(mem.Used)/1000000000)

	}
}



func main() {
	db:=openDatabase()
	go getCpuRam(db)
	go simulateDevice(db)
	for true{
		time.Sleep(1* time.Second)
		//cpu,ram := <-c1, <-c1
		//fmt.Printf("cpu: %f ,ram: %f Gbytes\n",cpu,ram)
	}
}
