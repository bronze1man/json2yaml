package main

import (
	"os"
	"github.com/bronze1man/json2yaml/j2yLib"
)

const version = "1.0"

func main() {
	firstArg:=""
	if len(os.Args)>=2{
		firstArg = os.Args[1]
	}
	if firstArg=="--help" || firstArg == "-h"{
		os.Stdout.WriteString(`Transform json string to yaml string without the type infomation.
Usage:
echo '{"a": 1}' | json2yaml
json2yaml < 1.yml > 2.json
`)
		os.Exit(1)
		return
	}
	if firstArg=="--version" || firstArg == "-v"{
		os.Stdout.WriteString(version)
		os.Stdout.WriteString("\n")
		os.Exit(0)
		return
	}
	err := j2yLib.TranslateStream(os.Stdin, os.Stdout)
	if err == nil{
		os.Exit(0)
	}
	os.Stderr.WriteString(err.Error())
	os.Stderr.WriteString("\n")
	os.Exit(2)
}

