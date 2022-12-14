# log4go

log4go aims to enrich the go log library. Log events by environment with a single line.

# Example
```go
package main

import (
	"fmt"
	auto "github.com/erhankrygt/log4go"
)

// Main function
func main() {

  	var log4 log.ILog
	{
		log4 = log.NewLog(log.Config{FileName: "log4go"}, "prod")
		if log4 == nil {
			_ = fmt.Errorf("initializing log failed")
		}
	}

	log4.Debug("Debug Message", "Debug Message 2")
	log4.Warn("Warn Message")
	log4.Info("Info Message")
	log4.Error("Error Message")
	log4.Fatal("Fatal Message")

}
```
# output
in log4go_prod.log file 

<img width="599" alt="image" src="https://user-images.githubusercontent.com/6412354/196037697-928184d5-c67c-4d2d-bad0-bdc87afc9d69.png">

# Installation
```
go get -u github.com/erhankrygt/log4go
```
