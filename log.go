package main

import (
	"github.com/wonderivan/logger"
	_ "github.com/wonderivan/logger"
)

//var logger Logger

//type Logger struct {
//	Debugf func(format string, v ...any)
//	Infof  func(format string, v ...any)
//	Errorf func(format string, v ...any)
//}

func InitLog() {
	logger.SetLogger("./log.json")
	logger.Debug("debug init : %v ", logger.LevelMap)

	//// 创建一个日志文件
	//file, err := os.OpenFile("log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatal("Cannot create log file", err)
	//}
	//defer file.Close()
	//
	//debugLogger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	//debugLogger.SetPrefix("[DEBUG] ")
	//
	//infoLogger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	//infoLogger.SetPrefix("[Info] ")
	//
	//errorLogger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	//errorLogger.SetPrefix("[error] ")
	//
	//logger.Debugf = debugLogger.Printf
	//logger.Infof = infoLogger.Printf
	//logger.Errorf = errorLogger.Printf
	//
	//logger.Debugf("debugLogger %s", "init")
	//logger.Infof("InfoLogger %s", "init")
	//logger.Errorf("ErrorLogger %s", "init")
}
