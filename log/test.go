package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
