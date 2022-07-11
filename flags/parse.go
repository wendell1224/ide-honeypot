package flags

import (
	"fmt"
	"ide-honeypot/model"
	"os"
)

func Parse(info *model.Info) {
	if info.Command == "" {
		fmt.Println("You should set command \n Example：\n ide-honeypot -h 0.0.0.0 -c whoami -f zipname -p 8080")
		os.Exit(0)
	}
	if info.Host == "" {
		fmt.Println("You should set host \n Example：\n ide-honeypot -h 0.0.0.0 -c whoami -f zipname -p 8080")
		os.Exit(0)
	}
	if info.Port == "" {
		info.Port = "8080"
	}
	if info.Zipname == "" {
		info.Zipname = "source.zip"
	}

}
