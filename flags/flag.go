package flags

import (
	"flag"
	"ide-honeypot/model"
)

func Flag(info *model.Info) {
	flag.StringVar(&info.Host, "h", "", "Set host: -h 127.0.0.1")
	flag.StringVar(&info.Command, "c", "", "Set Command: -c whoami")
	flag.StringVar(&info.Zipname, "f", "", "Set Filename: -f source")
	flag.StringVar(&info.Port, "p", "", "Set Port: -p 8080")
	flag.Parse()
}
