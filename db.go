package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"runtime"
)

var Store = ".onepass.save"
var DbHistory = cache.New(cache.NoExpiration, cache.NoExpiration)

func init() {
	if runtime.GOOS == "windows" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		Store = home + "/" + Store
	} else {
		Store = "~/" + Store
	}

	if FileExists(Store) {
		err := DbHistory.LoadFile(Store)
		if err != nil {
			fmt.Println("Error", err)
		}
	}
}

func HistoryContains(key string) (value string, ok bool) {
	defer DeferRecover()
	ok = false
	v, ok := DbHistory.Get(key)
	value = fmt.Sprint(v)
	return
}

func HistoryAdd(key string, value string) {
	defer DeferRecover()
	DbHistory.SetDefault(key, value)
	err := DbHistory.SaveFile(Store)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func HistoryList() {
	i := 0
	for k := range DbHistory.Items() {
		i++
		fmt.Println(i, k)
	}
}
