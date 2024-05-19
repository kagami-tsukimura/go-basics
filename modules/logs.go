package modules

import (
	"log"
	"os"
)

func mkDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		// ディレクトリの存在をチェック
		if os.IsNotExist(err) {
			// 存在しないので作る
			err = os.Mkdir(dir, 0755)
			if err != nil {
				// ディレクトリの作成に失敗
				panic(err)
			}
		}
	}
}

func createLogFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func Logs() {
	logDir := "logger"
	filePath := logDir + "/log.txt"
	mkDir(logDir)
	createLogFile(filePath)
}
