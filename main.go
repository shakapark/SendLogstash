package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mickep76/go-logstash"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/shakapark/SendLogstash/config"
)

var (
	sc = &config.SafeConfig{C: &config.Config{},}
	configFile = kingpin.Flag("config.file", "Configuration file.").Default("conf.yml").String()
)

func main(){
	kingpin.Parse()
	if err := sc.ReloadConfig(*configFile); err != nil {
		fmt.Println("Error loading config", err)
		os.Exit(1)
	}

	for _, server := range sc.C.Servers {
		l := logstash.New(server.Host+":"+strconv.Itoa(server.Port), 5)
		go l.Start()
		defer l.Stop()

		for _, entry := range server.Entries {
			l.Info(entry)
		}
	}
}
