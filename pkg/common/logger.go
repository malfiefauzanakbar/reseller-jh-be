package common

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Setup() {
	Log.SetFormatter(&logrus.JSONFormatter{})
}
