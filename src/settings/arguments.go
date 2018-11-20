package settings

import (
	"flag"
	"os"
)

func Init() {
	configurationFile := flag.String("config", "", "configuration file")
	flag.Parse()

	if *configurationFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	C = LoadConfiguration(*configurationFile)
}
