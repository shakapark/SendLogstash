package main

import (
	"fmt"
	"os"

	"github.com/heatxsink/go-logstash"
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
		l := logstash.New(server.Host, server.Port, 5)
		_, err := l.Connect()
		if err != nil {
			fmt.Println(err)
		}

		for _, entry := range server.Entries {
			err = l.Writeln(entry)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
