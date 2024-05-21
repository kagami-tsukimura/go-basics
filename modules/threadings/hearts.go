package threadings

import (
	"log"
	"os"
)

func Hearts() {
	file, err := os.Create("heart.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
}
