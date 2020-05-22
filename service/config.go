package service

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/arthurkushman/pgo"
	"github.com/spf13/viper"

	"github.com/mops1k/go-mvc/cli"
)

type config struct {
	reader *viper.Viper
}

var Config *config

func init() {
	Config = &config{reader: viper.New()}
	Config.reader.AutomaticEnv()
	Config.reader.SetConfigType("yaml")

	dirName := "./config/"
	_ = os.Mkdir(dirName, 0666)

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !pgo.InArray(filepath.Ext(file.Name()), []string{".yaml", ".yml"}) {
			continue
		}

		Config.AddFile(dirName + file.Name())
	}
}

// Get config parameter value by key
func (c *config) Get(key string) interface{} {
	return c.reader.Get(key)
}

// Get config string value by key
func (c *config) GetString(key string) string {
	return c.reader.GetString(key)
}

// Get config integer value by key
func (c *config) GetInt(key string) int {
	return c.reader.GetInt(key)
}

// Get config boolean value by key
func (c *config) GetBool(key string) bool {
	return c.reader.GetBool(key)
}

// Get config string slice by key
func (c *config) GetStringSlice(key string) []string {
	return c.reader.GetStringSlice(key)
}

// Get config string map by key
func (c *config) GetStringMap(key string) map[string]interface{} {
	return c.reader.GetStringMap(key)
}

// Get config duration value by key
func (c *config) GetDuration(key string) time.Duration {
	return c.reader.GetDuration(key)
}

// Add custom file to config
func (c *config) AddFile(path string) {
	c.reader.SetConfigFile(path)
	err := c.reader.MergeInConfig()
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Fatal(err)
	}
}
