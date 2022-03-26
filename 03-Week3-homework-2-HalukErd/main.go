package main

import (
	"fmt"
	"homework-2-week-3-HalukErd/command"
	"homework-2-week-3-HalukErd/service"
)

var Commands = make(map[string]command.Cmd)

func main() {
	PopulateCommands()
	command.PopulateParamParser()
	cmdReq := command.HandleReadArgs()
	cmd, ok := Commands[cmdReq.CmdKey]
	if commandNotFound(ok, cmdReq.CmdKey) {
		return
	}
	cmdReq.Params = command.HandleReadParams(cmd)
	cmd.Task(cmdReq.Params)
}

func commandNotFound(ok bool, cmdReq string) bool {
	if !ok {
		fmt.Printf("Could not run '%s' command.\nAvailable Commands are below.\n", cmdReq)
		for k, _ := range Commands {
			fmt.Println(k)
		}
		return true
	}
	return false
}

func PopulateCommands() {
	Commands["list"] = command.Cmd{Key: "list", Task: service.ListAllBooks}
	Commands["search"] = command.Cmd{Key: "search", Task: service.SearchBooksAndPrintResult}
	Commands["get"] = command.Cmd{Key: "get", Task: service.GetBookByIdAndPrintResult}
	Commands["delete"] = command.Cmd{Key: "delete", Task: service.DeleteBookByIdAndPrintResponse}
	Commands["buy"] = command.Cmd{Key: "buy", Task: service.BuyBookByIdAndQty}
}
