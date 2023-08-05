package utils

import (
	"log"
	"os"
)

var LOGFILENAME = "cliweather.log"
var F *os.File

// init function to start logger when program starts
func init() {
	InitLogger()
	log.Println("Initalized logger")
}

// inits log file with data_time.log
func InitLogger() {
	// curr_time := time.Now()
	// log_name := curr_time.Format("01-02-2006_15:04:05") + ".log"
	F, err := os.OpenFile(LOGFILENAME,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(F)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
