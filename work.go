package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	Verbose bool
	Force   bool
	Size    int
	ListApp bool
)

func init() {
	flag.BoolVar(&Verbose, "v", false, "verbose password")
	flag.BoolVar(&Force, "f", false, "force update")
	flag.BoolVar(&ListApp, "l", false, "list apps")
	flag.IntVar(&Size, "s", 10, "password size")
	flag.Parse()
	//fmt.Println("verbose:", Verbose)
	//fmt.Println("force:", Force)
	//fmt.Println("size:", Size)
}

func Work(app string) {
	v, ok := HistoryContains(app)
	if ok {
		fmt.Println("->Found in Cache:", app)
	} else {
		v = GenPassword(Size)
		HistoryAdd(app, v)
		fmt.Println("->Generated password for:", app, "\tsize:", Size)
	}
	if Force {
		v = GenPassword(Size)
		HistoryAdd(app, v)
		fmt.Println("->Force renewed password for:", app, "\tsize:", Size)
	}
	err := clipboard.WriteAll(v)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("->Already Copied. You can use it by Ctrl + V or Win + V")
	if Verbose {
		fmt.Println(v)
	}
}

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
		collector = append(collector, fmt.Sprint(i))
		collector = append(collector, fmt.Sprint(i))
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
