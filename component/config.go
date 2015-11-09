package component

import (
    "path/filepath"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "errors"
)

type Config struct {
    init bool
    Filename string
    Logstash struct {
        Host     string `yaml:"host"`
        Login    string `yaml:"login"`
        Password string `yaml:"password"`
    } `yaml:"logstash"`
}

func (o *Config) Init() {
    if o.init == false {
        if len(o.Filename) == 0 {
            panic(errors.New("Config file must be declared"))
        }

        filename, _ := filepath.Abs(o.Filename)

        yamlFile, err := ioutil.ReadFile(filename)

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
