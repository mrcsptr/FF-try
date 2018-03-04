package main

import (
	"encoding/json"
    "errors"
	"io/ioutil"
)

type Config struct {
	TeamsLocation string
	DudesLocation string
	ATeam string
	BTeam string
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
  
  if err := c.Check(); err!= nil {
      return Config{}, err
  }
  
  return c, nil
}

func (c Config) Check () error{
    if c.TeamsLocation==""{
        return errors.New("undefined TeamsLocation")
    }
    if c.DudesLocation==""{
        return errors.New("undefined DudesLocation")
    }
    if c.ATeam==""{
        return errors.New("undefined ATeam")
    }
    if c.BTeam==""{
        return errors.New("undefined BTeam")
    }
    return nil
}

func (c Config) String() string {
	raw, _ := json.MarshalIndent(c, "", "    ")
	return string(raw)
}
