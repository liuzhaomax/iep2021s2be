/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/30 23:56
 * @version     v1.0
 * @filename    logger.go
 * @description
 ***************************************************************************/
package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel) // Trace << Debug << Info << Warning << Error << Fatal << Panic
	InitializeLogging("golog.txt")
}

func InitializeLogging(logFile string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Cannot open log file: " + err.Error())
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	// log.SetFormatter(&log.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func LoggerToFile() gin.HandlerFunc {
	fileName := "golog.txt"
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Cannot open log file: " + err.Error())
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %8s | %s ",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	fileName := "golog.txt"
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Cannot open log file: " + err.Error())
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(io.MultiWriter(src, os.Stdout))
	if v[0] == "sql" {
		logger.WithFields(logrus.Fields{
			"module":  "gorm",
			"type":    "sql",
			"rows":    v[5],
			"src_ref": v[1],
			"values":  v[4],
		}).Print(v[3])
	}
	if v[0] == "log" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

//func GormLogger() *logrus.Logger {
//	fileName := "bin/golog.txt"
//	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//	if err != nil {
//		fmt.Println("Cannot open log file: " + err.Error())
//	}
//	logger := logrus.New()
//	logger.Out = src
//	logger.SetLevel(logrus.InfoLevel)
//	logger.SetFormatter(&logrus.TextFormatter{})
//	logger.SetOutput(io.MultiWriter(src, os.Stdout))
//	return logger
//}

//logger.WithFields(logger.Fields{
//	"animal": "walrus",
//	"size":   10,
//}).Info("A group of walrus emerges from the ocean")
//
//logger.WithFields(logger.Fields{
//	"omg":    true,
//	"number": 122,
//}).Warn("The group's number increased tremendously!")
//
//logger.WithFields(logger.Fields{
//	"omg":    true,
//	"number": 100,
//}).Fatal("The ice breaks!")
