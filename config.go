package main

import "errors"

type Config struct {
	Filename string
}

func (c Config) Validate() error {
	if c.Filename == "" {
		return errors.New("filename must not be empty")
	}

	return nil
}
