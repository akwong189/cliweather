package utils

import (
    "log"
    "os"
    "time"
)

var Log *log.Logger
var F *os.File
var err error

// init function to start logger when program starts
func init() {
    Log = InitLogger() 
    Log.Println("Initalized logger")
}

// inits log file with data_time.log
func InitLogger() *log.Logger {
    curr_time := time.Now()
    log_name := curr_time.Format("01-02-2006_15:04:05") + ".log"
    F, err = os.OpenFile("pkg/logs/" + log_name,
    	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
	    log.Println(err)
    }

    return log.New(F, "", log.LstdFlags | log.Lshortfile)
}
