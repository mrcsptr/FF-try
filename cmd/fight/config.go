package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Config is a type hosting the data coming from config.json
type Config struct {
	TeamsLocation string
	DudesLocation string
	ATeam         string
	BTeam         string
}

// NewConfig gets the information from config.json to get the data to fight
func NewConfig(filepath string) (Config, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var c Config
	//reading the content of config.json
	if err := json.Unmarshal(raw, &c); err != nil {
		return Config{}, err
	}
	//checking conformity of informations
	if err := c.Check(); err != nil {
		return Config{}, err
	}

	return c, nil
}

//Check verifies the content of Config to be in accordance with what's expected
func (c Config) Check() error {
	if c.TeamsLocation == "" {
		return errors.New("undefined TeamsLocation")
	}
	if c.DudesLocation == "" {
		return errors.New("undefined DudesLocation")
	}
	if c.ATeam == "" {
		return errors.New("undefined ATeam")
	}
	if c.BTeam == "" {
		return errors.New("undefined BTeam")
	}
	return nil
}

// String implements the fmt.Stringer interface.
// It returns the json representation of the config.

func (c Config) String() string {
	raw, _ := json.MarshalIndent(c, "", "    ")
	return string(raw)
}
