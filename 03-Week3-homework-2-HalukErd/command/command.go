package command

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Cmd struct {
	Key  string
	Task func(params Params)
}

type Params map[string]string

type cmdReq struct {
	CmdKey string
	Params Params
}

var paramParser = make(map[string]func() Params)

func PopulateParamParser() {
	paramParser["search"] = ReadNameParam
	paramParser["get"] = ReadBookIdParam
	paramParser["delete"] = ReadBookIdParam
	paramParser["buy"] = ReadIdAndQtyParam
}

func HandleReadParams(c Cmd) Params {
	paramParserFunc, ok := paramParser[c.Key]
	if !ok {
		return Params{}
	}
	return paramParserFunc()
}

func ReadNameParam() Params {
	args := os.Args
	if len(args) < 2 {
		return nil
	}
	paramWords := args[2:]
	param := make(map[string]string)
	param["name"] = strings.Join(paramWords, " ")
	return param
}

func ReadIdAndQtyParam() Params {
	args := os.Args
	if len(args) < 4 {
		return nil
	}
	param := make(map[string]string)
	param["bookId"] = args[2]
	param["bookQty"] = args[3]
	return param
}

func ReadBookIdParam() Params {
	args := os.Args
	if len(args) < 3 {
		return nil
	}
	param := make(map[string]string)
	param["bookId"] = args[2]
	return param
}

func HandleReadArgs() cmdReq {
	cmd, err := getCmdAndArgs()
	if err != nil {
		fmt.Println("You need to enter at least a cmdReq.")
	}
	return cmd
}

func getCmdAndArgs() (cmdReq, error) {
	args := os.Args
	var cmd cmdReq
	if len(args) < 2 {
		return cmd, errors.New("not enough args")
	}
	cmd.CmdKey = args[1]

	return cmd, nil
}
