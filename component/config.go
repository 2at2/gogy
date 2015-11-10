package component

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"os"
	"github.com/mitchellh/go-homedir"
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

		// Read from current folder
		current, _ := os.Getwd()
		if o.existsFile(current, filename) {
			path = stat
		}

		home, err := homedir.Dir()
		if err != nil {
			if o.existsFile(home, filename) {
				path = stat
			}
		}

		if len(path) == 0 {
			panic(errors.New("Cannot find config file"))
		}

		path, _ = filepath.Abs(path)

		yamlFile, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(yamlFile, &o)
		if err != nil {
			panic(err)
		}

		o.init = true
	}
}

func(o *Config) existsFile(folder, filename string) (bool) {
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
