package component

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	init     bool
	Env      string
	Filename string
	File     *os.File
	Source   map[string]interface{}
}

func (o *Config) InitConfigFile(filename string) {
	if o.init == false {
		o.Filename = filename + "." + o.Env + ".yml"
		if len(o.Filename) == 0 {
			panic(errors.New("Config file must be declared"))
		}

		var path string

		path, _ = filepath.Abs(filename)

		yamlFile, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(yamlFile, &o.Source)
		if err != nil {
			panic(err)
		}

		o.init = true
	}
}

func (o *Config) existsFile(folder, filename string) bool {
	pathSeparator := string(os.PathSeparator)

	if len(folder) > 0 && folder[len(folder)-1:] != pathSeparator {
		folder = folder + pathSeparator
	}

	name := folder + filename

	if _, err := os.Stat(name); err == nil {
		return true
	} else {
		return false
	}
}
