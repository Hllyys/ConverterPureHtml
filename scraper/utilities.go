package scraper

import (
	"fmt"
	"log"
)

func CheckFatal(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}

}

func CheckErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CheckStatusCode(statusCode int, statusText string) error {
	if statusCode != 200 {
		return fmt.Errorf("status code error: %d %s", statusCode, statusText)
	}
	return nil
}
