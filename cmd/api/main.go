package main

import (
	"flag"
	"fmt"

	"github.com/raspibuddy/rpi/pkg/api"
	"github.com/raspibuddy/rpi/utl/config"
)

func main() {
	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	fmt.Println("No error when parsing the config file")
	checkErr(api.Start(cfg))

	// fmt.Println(cpu.Times(false))
	// fmt.Println(cpu.Percent(1, true))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
