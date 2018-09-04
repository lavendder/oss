package log

import (
	"github.com/astaxie/beego"
	"oss_dfs/library/utils"
	"github.com/astaxie/beego/logs"
	"runtime"
	"strings"
	"fmt"
	"oss_dfs/library/server"
)

type OssLog struct {
	logs.BeeLogger
}

// init project logger instance
var OssLogger = new(OssLog)
var beeLogger = logs.NewLogger()

// int log
func init() {
	fileConfig := getLogFilePath()
	beeLogger.SetLogger("file", fileConfig)
	// set config log level
	setLogLevel()
}

// format log message
func Format(message string) string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("- requestId: %s - filePath: %s - onLine: %d - %s ",
			server.GetRequestId(),
			parseFilePath(file),
			line,
			message)
	}
	return fmt.Sprintf("- requestId: %s - %s ", server.GetRequestId(), message)
}

func (ol *OssLog) Info(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Info(message, v...)
}

func (ol *OssLog) Warning(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Warning(message, v...)
}

func (ol *OssLog) Emergency(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Emergency(message, v...)
}

func (ol *OssLog) Alert(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Alert(message, v...)
}

func (ol *OssLog) Critical(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Critical(message, v...)
}

func (ol *OssLog) Error(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Error(message, v...)
}

func (ol *OssLog) Notice(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Notice(message, v...)
}

func (ol *OssLog) Informational(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Informational(message, v...)
}

func (ol *OssLog) Debug(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Debug(message, v...)
}

func (ol *OssLog) Warn(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Warn(message, v...)
}

func (ol *OssLog) Trace(message string, v ...interface{}) {
	message = Format(message)
	beeLogger.Trace(message, v...)
}

/**
 *	LevelEmergency = iota
 *	LevelAlert
 *	LevelCritical
 *	LevelError
 *	LevelWarning
 *	LevelNotice
 *	LevelInformational
 *	LevelDebug
 */
func setLogLevel() {
	level := beego.AppConfig.String("log::log_level")
	switch level {
	case "LevelEmergency":
		beeLogger.SetLevel(beego.LevelEmergency)
	case "LevelAlert":
		beeLogger.SetLevel(beego.LevelAlert)
	case "LevelCritical":
		beeLogger.SetLevel(beego.LevelCritical)
	case "LevelError":
		beeLogger.SetLevel(beego.LevelError)
	case "LevelWarning":
		beeLogger.SetLevel(beego.LevelWarning)
	case "LevelNotice":
		beeLogger.SetLevel(beego.LevelNotice)
	case "LevelInformational":
		beeLogger.SetLevel(beego.LevelInformational)
	case "LevelDebug":
		beeLogger.SetLevel(beego.LevelDebug)
	default:
		beeLogger.SetLevel(beego.LevelDebug)
	}
}

// get log file path config info
func getLogFilePath() string {
	fileName := beego.AppConfig.String("log::log_file")
	fileName = fileName + "-" + utils.GetCurrentDate()
	fileConfig := `{"filename":"` + fileName + `.log"}`
	return fileConfig
}

// get just file in project path not full path
func parseFilePath(fullPath string) string {
	appname := beego.AppConfig.String("appname")
	paths := strings.Split(fullPath, appname)
	if len(paths) >= 2 {
		return paths[1]
	}
	return ""
}
