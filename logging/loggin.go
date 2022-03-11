package logging

import (
	"io"
	"os"
	"time"

	"github.com/cheolgyu/sbm-base/c"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var var_log *log.Logger
var Log *log.Logger

func init() {
	set_config_log()
	Log = var_log
	Log.Debug("logging init 실행")
}
func set_config_log() {
	var_log = log.New()

	if c.DEBUG {
		var_log.SetFormatter(&log.TextFormatter{})
		var_log.SetLevel(log.DebugLevel)
	} else {
		var_log.SetLevel(log.InfoLevel)
		var_log.SetFormatter(&log.JSONFormatter{})
	}
	var_log.SetReportCaller(true)
	saveLocalFile()
}
func ElapsedTime() func() {
	start := time.Now()
	return func() { log.Info("걸린시간:", time.Since(start)) }
}
func saveLocalFile() {
	t := time.Now()
	dirname := t.Format("2006-01-02")
	filename := t.Format("2006-01-02_15_04_05_000Z")

	os.MkdirAll("./logs/"+dirname, 0755)
	fnm := "./logs/" + dirname + "/" + filename

	log_file := &lumberjack.Logger{
		Filename:   fnm,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}

	multi := io.MultiWriter(log_file, os.Stdout)
	var_log.SetOutput(multi)
}
