package j2yLib

import (
	"testing"
	"bytes"
)

func TestTranslate(ot *testing.T){
	type tCas struct{
		in string
		out string
	}
	var casList = []tCas{
		{`{"property":"0testS_1value"}`+"\n",`property: 0testS_1value`+"\n"},
		{`{  "property"  :  "0testS_1value"  }`+"\n\n",`property: 0testS_1value`+"\n"},
		{`{"v":"hi"}{"v":"hi"}`,`v: hi
---
v: hi
`},
		{`{"v":4294967296}`,`v: 4.294967296e+09
`},
		{`{"a":1,"b":2}`,`a: 1
b: 2
`},
		{`{"56":"bar"}`,`"56": bar
`},
		{`{"m":false}`,`m: false
`},
		{`{"m":true}`,`m: true
`},
		{`{"null":null}`,`"null": null
`},
		{`[null,null]`,`- null
- null
`},
		{`{"5.6":5.6}`,`"5.6": 5.6
`},
}
	for _,cas:=range casList{
		out:=mustTranslateStreamString(cas.in)
		if out!=cas.out{
			panic("fail in:["+cas.in+"] thisOut:["+string(out)+"] expect:["+cas.out+"]")
		}
	}
}

func mustTranslateStreamString(inSlice string) (outSlice string){
	inBuf:=bytes.NewBufferString(inSlice)
	outBuf:=&bytes.Buffer{}
	err:= TranslateStream(inBuf,outBuf)
	if err!=nil{
		panic(err)
	}
	return outBuf.String()
}