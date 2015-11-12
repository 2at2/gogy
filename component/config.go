package component

import (
	"github.com/gotterdemarung/go-configfile"
	"gopkg.in/yaml.v2"
)

const ConfigFolder = "config"
const DefaultConfig = "params.yml"

func LoadConfig(configFile string) (Reader, error) {
	configReader := configfile.ConfigReader{
		Subfolder: ConfigFolder,
	}

	data, err := configReader.ReadBytes(configFile)

	if err != nil {
		return nil, err
	}

	var config Reader
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}
