package j2yLib

import (
	"io"
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func TranslateStream(in io.Reader, out io.Writer) error {
	decoder := json.NewDecoder(in)
	encoder:=yaml.NewEncoder(out)
	for {
		var data interface{}
		err := decoder.Decode(&data)
		if err != nil {
			if err==io.EOF{
				return nil
			}
			return err
		}
		err = encoder.Encode(data)
		//output, err := yaml.Marshal(data)
		//if err != nil {
		//	return err
		//}
		//data = nil
		//_, err = out.Write(output)
		if err != nil {
			return err
		}
		//_, err = io.WriteString(out, "\n")
		//if err != nil {
		//	return err
		//}
	}
}