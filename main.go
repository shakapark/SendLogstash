package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

        "github.com/bshuster-repo/logrus-logstash-hook"
        "github.com/sirupsen/logrus"
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

	fmt.Println("Starting Sender")

	log := logrus.New()
	for _, server := range sc.C.Servers {
		conn, err := net.Dial("tcp", server.Host+":"+strconv.Itoa(server.Port))
		if err != nil {
                	log.Warnln(err)
			break
        	}
		fmt.Println("New Server : ", server.Host, ":", strconv.Itoa(server.Port))

                for entryName, entry := range server.Entries {
	                hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"entry": entryName}))
                	log.Hooks.Add(hook)

			mapStr := make(map[string]interface{})
			for _, logs := range entry {
				t := strings.Split(logs, ":")
				mapStr[t[0]] = t[1]
			}
			ctx := log.WithFields(logrus.Fields(mapStr))
			ctx.Info("Hello World!")
		}
		fmt.Println("Stop Connection")
	}
}
