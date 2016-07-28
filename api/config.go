package api

import (
  "io/ioutil"
  "os"

  "github.com/naoina/toml"
)

type Config struct {
  Mode     string
  Port     string
  Host     string
  Database struct {
    Url string
  }
}

func ParseConfig(filename string) Config {
  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  buf, err := ioutil.ReadAll(f)
  if err != nil {
    panic(err)
  }
  var config Config
  if err := toml.Unmarshal(buf, &config); err != nil {
    panic(err)
  }
  return config
}
