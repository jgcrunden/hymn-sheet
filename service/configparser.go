package service

import (
	"encoding/json"

	"github.com/jgcrunden/hymn-sheet/model"
)

type ConfigParser struct {
	Config model.Config
	reader FileReader
}

func NewConfigParser(reader FileReader) ConfigParser {
	return ConfigParser{reader: reader}
}

func (c *ConfigParser) ReadConfigFile(filename string) error {
	content, err := c.reader.ReadFile(filename)
	if err != nil {
		return err
	}

	var conf model.Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return err
	}
	c.Config = conf
	return nil
}

func (c *ConfigParser) ValidateConfig() error {
	// TODO: To implement
	return nil
}
