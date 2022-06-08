// go mod vendor の挙動を確認するためのサンプルです。
// 詳細については README.md を参照ください。
package main

import (
	"log"

	"github.com/devlights/gomy/logops"
)

func main() {
	// 本アプリは github.com/devlights/gomy/logops に依存している

	var (
		appLog, _, _ = logops.Default.Logger(false, func(appL, _, _ *log.Logger) {
			appL.SetPrefix(">>> ")
		})
	)

	appLog.Println("hello world")
}
