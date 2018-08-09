package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}
