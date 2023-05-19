package logger

import (
	"os"
	"runtime"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

var singleton *log.Logger
var once sync.Once

func Init() {
	once.Do(func() {
		singleton = log.New()
		singleton.SetFormatter(&log.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (
				function string, file string) {
				return f.File, strconv.Itoa(f.Line)
			},
			TimestampFormat: "15:04:05 2006-01-02",
			FullTimestamp:   true,
		})
		singleton.SetOutput(os.Stdout)
		singleton.SetLevel(log.InfoLevel)
		singleton.SetReportCaller(true)
	})
}

func GetLogger() *log.Logger {
	return singleton
}
