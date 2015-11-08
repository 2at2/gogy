package component
import (
    "path/filepath"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type Config struct {
    init bool
    Logstash struct {
        Host     string `yaml:"host"`
        Login    string `yaml:"login"`
        Password string `yaml:"password"`
    } `yaml:"logstash"`
}

func (o *Config) Init() {
    if o.init == false {
        filename, _ := filepath.Abs("./config/params.yml")
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
