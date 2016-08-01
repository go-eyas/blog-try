package config

import (
  "os"
  "io/ioutil"
  "github.com/naoina/toml"
)

type Conf struct {
  Mode       string
  ClientMode string
  Port       string
  Host       string
  Database   struct {
    Url string
  }
}

func parseConfig(filename string) Conf {
  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  buf, err := ioutil.ReadAll(f)
  if err != nil {
    panic(err)
  }
  var config Conf
  if err := toml.Unmarshal(buf, &config); err != nil {
    panic(err)
  }
  return config
}

func Get() Conf {
  return parseConfig("config.ini")
}