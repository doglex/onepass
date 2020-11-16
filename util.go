package main

import (
	"log"
	"os"
	"runtime"
)

func DeferRecover() {
	if err := recover(); err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Println("Recover", file, line, err)
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
