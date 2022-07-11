package main

import (
	"fmt"
	"ide-honeypot/flags"
	"ide-honeypot/model"
	"ide-honeypot/route"
	"ide-honeypot/write"
	"ide-honeypot/zip"
)

func main() {
	var info model.Info
	flags.Flag(&info)
	fmt.Println(info.Command)
	flags.Parse(&info)
	write.Write(info)
	zip.Zip("source", info.Zipname+".zip")
	write.WriteOld(info.Command)
	route.Init(info)

}
