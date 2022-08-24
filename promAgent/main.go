package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"promAgent/cmds"
	_ "promAgent/init"
	"promAgent/tasks"
	"promAgent/utils"
	"syscall"
	"time"
)

func main() {
	cmds.Execute()
	if cmds.Help {
		fmt.Println("Help")
		os.Exit(0)
	}
	if cmds.Verbose {
		fmt.Println("Verbose")
	}
	config := utils.InitConfig(cmds.Path)
	utils.InitLog(cmds.Verbose, config)
	logrus.WithField("k", "v").Debug("This is logrus debug log")
	logrus.WithField("k", "v").Info("This is logrus info log")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT)
	reload := make(chan os.Signal, 1)
	signal.Notify(reload, syscall.SIGHUP)
	errChan := make(chan error, 1)

	go func() {
		for {
			logrus.Debug("doing")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			<-reload
			fmt.Println("Go channel receive the reload signal")
		}
	}()

	logrus.WithFields(logrus.Fields{"pid": os.Getpid()}).Info("Start promAgent...")

	go func() {
		tasks.Run(config, errChan)
	}()

	select {
	case <-stop:
		logrus.Info("promAgent stopped")
	case err := <-errChan:
		logrus.Error("promAgent Error :", err)
		os.Exit(1)
	}

}
