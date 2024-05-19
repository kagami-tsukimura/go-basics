package threadings

import (
	"log"
	"os"
)

func Tracers() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error: ", err)
		}
	}()

}
