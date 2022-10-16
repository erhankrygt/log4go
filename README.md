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

  var lgo log.ILog
	{
		lgo = log.NewLog(log.Config{FileName: "log4net"}, "prod")
		if lgo == nil {
			_ = fmt.Errorf("initializing log failed")
		}
	}

	lgo.Debug("Debug Message")
	lgo.Warn("Warn Message")
	lgo.Info("Info Message")
	lgo.Error("Error Message")
	lgo.Fatal("Fatal Message")
    

  //output
  in log4net_prod.log file 
  
}
```
  <img width="423" alt="image" src="https://user-images.githubusercontent.com/6412354/196023617-e183e152-6918-4131-b555-ace50ce8823d.png">

# Installation
```
go get -u github.com/erhankrygt/log4go
```
