package config

import (
    "path/filepath"
    "os"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/yaml.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/caarlos0/env"
)

// New fills the provided obj with values from given file(s) (add env.env for parsing the environment)
func New(obj interface{}, files ...string) {

    if len(files) > 0 {

        for file := range files {
            parseByFile(file, obj)
        }
    }
}

func parseByFile(file string, c interface{}) {

    ext := filepath.Ext(file)

    switch ext {

    case "json":
        parseByJSON(file, c)
        break
    case "yml":
    case "yaml":
        parseByYaml(file, c)
        break
    case "bson":
        parseByBson(file, c)
        break
    case "env":
        parseByEnv(c)
    }
}

func parseByEnv(c interface{}) {

    if err := env.Parse(c); err != nil {
        panic(err)
    }
}

func parseByJSON(file string, c interface{}) {

    if _, err := os.Stat(file); err != nil {
        panic(err)
    }

    bytes, err := ioutil.ReadFile(file)
    if err != nil { panic(err) }

    if err := json.Unmarshal(bytes, c); err != nil {
        panic(err)
    }
}

func parseByYaml(file string, c interface{}) {

    if _, err := os.Stat(file); err != nil {
        panic(err)
    }

    bytes, err := ioutil.ReadFile(file)
    if err != nil { panic(err) }

    if err := yaml.Unmarshal(bytes, c); err != nil {
        panic(err)
    }
}

func parseByBson(file string, c interface{}) {

    if _, err := os.Stat(file); err != nil {
        panic(err)
    }

    bytes, err := ioutil.ReadFile(file)
    if err != nil { panic(err) }

    if err := bson.Unmarshal(bytes, c); err != nil {
        panic(err)
    }
}