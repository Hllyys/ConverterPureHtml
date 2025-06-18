package scraper

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckFatal(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
