package service

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/arthurkushman/pgo"
	"github.com/spf13/viper"
	"golang.org/x/text/language"

	"github.com/mops1k/go-mvc/cli"
)

type translation struct {
	defaultLocale  language.Tag
	fallbackLocale language.Tag
	dictionary     map[language.Tag]map[string]string
}

var Translation *translation

func init() {
	Translation = &translation{
		defaultLocale:  language.Make(Config.GetString("translation.default")),
		fallbackLocale: language.Make(Config.GetString("translation.fallback")),
		dictionary:     make(map[language.Tag]map[string]string),
	}
	dirname := "./translations/"
	_ = os.Mkdir(dirname, 0666)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	for _, file := range files {
		if file.IsDir() || !pgo.InArray(filepath.Ext(file.Name()), []string{".yaml", ".yml"}) {
			continue
		}

		v := viper.GetViper()
		v.SetConfigType("yaml")
		v.SetConfigFile(dirname + file.Name())
		err = v.ReadInConfig()
		if err != nil {
			cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
		}

		pattern, err := regexp.Compile(`[\w\d\S]*\.(\w{2,5})\.(yaml|yml)$`)
		if err != nil {
			cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
		}

		matches := pattern.FindAllStringSubmatch(file.Name(), -1)
		locale := language.Make(matches[0][1])

		for _, key := range v.AllKeys() {
			if Translation.dictionary[locale] == nil {
				Translation.dictionary[locale] = make(map[string]string)
			}

			Translation.dictionary[locale][key] = v.GetString(key)
		}
	}
}

func (t *translation) TransFor(key string, arguments map[string]interface{}, locale language.Tag) string {
	text, exists := t.dictionary[locale][key]
	if !exists {
		text, exists = t.dictionary[t.fallbackLocale][key]
		if !exists {
			return key
		}
	}

	return Template.RenderString(text, arguments)
}

func (t *translation) Trans(key string, arguments map[string]interface{}) string {
	locale := t.defaultLocale
	text, exists := t.dictionary[locale][key]
	if !exists {
		text, exists = t.dictionary[t.fallbackLocale][key]
		if !exists {
			return key
		}
	}

	return Template.RenderString(text, arguments)
}
