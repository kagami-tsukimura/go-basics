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

func Logs() {
	mkDir("logger")
	file, err := os.Create("logger/log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
}
