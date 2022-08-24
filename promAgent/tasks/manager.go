package tasks

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"promAgent/config"
)

type Task interface {
	Init(config *config.AgentConfig)
	Run()
}
type manager map[string]Task

var mgr = make(manager)

func Register(name string, task Task) {
	if _, ok := mgr[name]; ok {
		logrus.WithFields(logrus.Fields{"task": name}).Fatal("Task is exists")
	}
	mgr[name] = task
	logrus.WithFields(logrus.Fields{"task": name}).Info("Task is registered")
}

func Run(config *config.AgentConfig, errChan chan<- error) {
	for name, task := range mgr {
		task.Init(config)
		//一定要写变量,例程启动时,可能name和task已经变为别的了
		go func(name string, task Task) {
			logrus.WithFields(logrus.Fields{"task": name}).Info("Task is running")

			task.Run()

			logrus.WithFields(logrus.Fields{"task": name}).Error("Task is stopped")

			errChan <- fmt.Errorf("task %s is stopped\n", name)

		}(name, task)
	}
}
