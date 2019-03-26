package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
func main() {
	LoggingSettings("test.log")
	file, err := os.Open("fdfsdfsdf")
	if err != nil {
		log.Fatalln("exit!!", err)
	}
	log.Println(file)
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// この時点でプログラムが終了する
	log.Fatalln("error!!")
}
