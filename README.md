# ubiwhereChallenge  

#### This software uses Golang 1.13  

###Get started    
To use this software first clone it:  
```
git clone https://github.com/pedrooct/ubiwhereChallenge.git
```
  
##### Installing

```
- go get -u github.com/spf13/cobra/cobra
- go get -u github.com/gin-gonic/gin
- go get github.com/mattn/go-sqlite3
- go get github.com/mackerelio/go-osstat/memory 
- go get github.com/mackerelio/go-osstat/cpu
- go get -u google.golang.org/grpc
- go get github.com/pedrooct/ubiwhereChallenge/gRPC/pb
(opcional- IF compile proto is need, please install.)
- go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```
  
### Running
##### Main Process
Go to the root directory of the repo folder and run:

```
go run main.go
```
  
This will generate cpu/ram info,4 random integers and save it on the database(sqlite3).
The database is already populate for immediate use.

##### CLI  

Go to the cli directory, The main.go file allows the use of cli with the following commands:
```
- go run main.go test (Verifies that the cli is working)
- go run main.go getLast N (E.g: go run main.go getLast 3)
- go run main.go getLastV N D (E.g: go run main.go getLast 3 D1 D2)
- go run main.go getAvg D (E.g: go run main.go getLast D1)
```

##### RestApi
Go to the restApi directory and execute the following command:
```
go run main.go
```
This will open a rest API on the port :8008. To test the API you can try the following requests:
(The use of curl is to simplify the process, postman can be a good choice as well).
```
curl --location --request GET '127.0.0.1:8008/get/last/3' \
--data-raw '' (This will return the last N elements)

curl --location --request GET '127.0.0.1:8008/get/lastv/5' \
--header 'Content-Type: text/plain' \
--data-raw 'd1,d3,d4' (This will return the last N for d1,d3,d4)

curl --location --request GET '127.0.0.1:8008/get/avg' \
--header 'Content-Type: text/plain' \
--data-raw 'd1,d3,d4' (This will return the Avg for d1,d3,d4)
```
(The calls above return a string and not a json!)

##### gRPC

Go to the gRPC directory and run the following commands:
- First enter the server folder and execute: go run main.go
- Second go to the client folder and execute: go run main.go
With this process you will see that the client send a message(json) to the server with the generate values.  
P.S: The gRPC generator doesn't save to the database.

### Final considerations 

Since this was a challenge database security and transactions were not implemented. 