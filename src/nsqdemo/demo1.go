package main

import (
	"log"
	"nsqdemo/models"
)

func addMsg() error {

	logmodel := models.Log{}

	log.Println(logmodel)

	logmsg := models.Log{FAddtime: "tttm", FContent: "xiaoxineirong"}

	log.Println(logmsg)
	logmodel.Add(logmsg)
	return nil
}

func main() {

	for i := 0; i < 20; i++ {
		addMsg()
	}
}
