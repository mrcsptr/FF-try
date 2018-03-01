package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SomeNumber int
	SomeString string
}

func NewConfig(filepath string) (Config, error){
  raw, err := ioutil.ReadFile(filepath)
  if err != nil {
    return Config{}, err
  }

  var c Config
  if err := json.Unmarshal(raw, &c); err != nil {
    return Config{}, err
  }
  return c, nil
}


func (c Config) String() string {
	raw, _ := json.MarshalIndent(c, "", "    ")
	return string(raw)
}
