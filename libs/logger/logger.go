package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	logPath := "logs/app.log"
	logfile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln("open logfile failed")
	}
	//defer logfile.Close()
	log.SetOutput(logfile)
	log.SetLevel(log.WarnLevel)
}
