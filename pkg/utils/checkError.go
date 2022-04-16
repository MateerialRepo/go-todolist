package utils

import "log"

//CheckErr is used to check for errors
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
