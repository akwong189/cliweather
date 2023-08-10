package utils

import (
	"log"
	"regexp"
)

func ValidIPAddress(addr string) bool {
	exp, err := regexp.Compile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	if err != nil {
		log.Fatalln(err)
	}

	return exp.MatchString(addr)
}
