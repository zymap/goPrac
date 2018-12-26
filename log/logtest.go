package main

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path"
	"time"
)
func SetLogLevel(level string) error {
	l, err := logrus.ParseLevel(level)
	if err == nil {
		logrus.SetLevel(l)
	}

	return err
}

func initLog() {
	// format log output
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05", DisableTimestamp: false})

	// set log level
	//conf := config.GetConf()
	err := SetLogLevel("DEBUG")
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	}


	currentPath, err := getCurrentPaht()
	if err != nil {
		panic("can not get path")
	}

	fmt.Println(currentPath)


	logPath := currentPath + "/log/logs/app.log"
	fmt.Println(logPath)
	fmt.Println("============")

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		panic("Can not open log file: " + err.Error())
	}

	baseLogPath := path.Join("/Users/zhangyong/GoProjects/goPrac/log/output/", "app.log")
	writer, err := rotatelogs.New(
		baseLogPath + ".%Y-%m-%d-%H-%M",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)

	fmt.Println(writer.CurrentFileName())
	fmt.Println("---------?")

	if err != nil {
		logrus.Errorf("config local file system logger error. %+v ", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel: writer,
		logrus.WarnLevel: writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})


	logrus.AddHook(lfHook)
	logrus.SetOutput(file)


}


func getCurrentPaht() (string, error) {
	cmd := exec.Command("pwd")
	out, err := cmd.Output()
	path := string(out[:len(out)-1])

	return path, err
}

func main()  {
	initLog()
	path, _ := getCurrentPaht()
	p := path + "/hello"
	print(p)
	fmt.Println(p)

	//fmt.Println(time.Now())
	for {
		time.Sleep(time.Second * 10)
		logrus.Println("hello ", time.Now())
	}
}