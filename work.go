package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenPassword(size int) (password string) {
	var choices = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	collector := make([]string, 0)
	for _, c := range choices {
		collector = append(collector, strings.ToLower(c))
		collector = append(collector, c)
	}
	for i := 0; i < 10; i++ {
		collector = append(collector, fmt.Sprint(i))
	}
	rand.Seed(time.Now().Unix())
	passwordCollector := make([]string, 0)
	for i := 0; i < size; i++ {
		c := fmt.Sprint(collector[rand.Intn(len(collector))])
		passwordCollector = append(passwordCollector, c)
	}
	password = strings.Join(passwordCollector, "")
	return
}
