package component

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"os"
	"fmt"
	"log"
)

type Config struct {
	init     bool
	Env string
	Filename string
	Source map[string]interface{}
}

func (o *Config) Init() {
	if o.init == false {
		if len(o.Filename) == 0 {
			panic(errors.New("Config file must be declared"))
		}

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)

		var filename string
		if _, err := os.Stat("./config/"); os.IsExist(err) {
			filename = "./config/" + o.Filename + ".yml"
		} else if _, err := os.Stat("./../config/"); os.IsExist(err) {
			filename = "./../config/" + o.Filename + ".yml"
		} else {
			panic(errors.New("Cannot find config directory"))
		}

		path, _ := filepath.Abs(filename)

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
