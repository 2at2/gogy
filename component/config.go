package component

import (
	"bytes"
	"github.com/gotterdemarung/go-configfile"
)

func LoadConfig(configFile string) ([]byte, error) {
	configReader := configfile.ConfigReader{Subfolder: "config"}

	if file, err := configReader.GetFile(configFile); err == nil {
		buf := new(bytes.Buffer)

		if _, err := buf.ReadFrom(file); err == nil {
			return buf.Bytes(), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
