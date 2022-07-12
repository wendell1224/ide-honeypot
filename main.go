package main

import (
	"fmt"
	"ide-honeypot/flags"
	"ide-honeypot/model"
	"ide-honeypot/route"
	"ide-honeypot/write"
	"ide-honeypot/zip"
	"log"
)

func main() {
	banner()
	var info model.Info
	flags.Flag(&info)
	log.Println("嵌入的命令为" + info.Command)
	flags.Parse(&info)
	write.Write(info)
	zip.Zip("source", info.Zipname+".zip")
	write.WriteOld(info.Command)
	route.Init(info)

}

func banner() {
	fmt.Println("\n\n  ___ ___  ___    _  _                            _   \n |_ _|   \\| _____| || |___ _ _  ___ _  _ _ __ ___| |_ \n  | || |) | _|___| __ / _ | ' \\/ -_| || | '_ / _ |  _|\n |___|___/|___|  |_||_\\___|_||_\\___|\\_, | .__\\___/\\__|\n                                    |__/|_|           \n")
	fmt.Println("		IDE-Honeypot by WenD1l\n")
}
