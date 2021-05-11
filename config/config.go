package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	configType      = "json"
	configDirectory = "./config/json"
)

/*
example
	obj := config.Country{}
	config.ReadAndMarshal("./config/json/countries.json","json", &obj)
	fmt.Println(obj)
*/
func ReadAndMarshal(path string, ext string, cfg interface{}) {
	//Read non env config from a different directory
	viper.AddConfigPath(".")
	viper.SetConfigName(path) //Path should n
	viper.SetConfigType(ext)
	err := viper.MergeInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.AllKeys())
}

func ReadConfigFromDir(dir string) {
	viper.AddConfigPath(".")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		if !f.IsDir() {
			fileName := f.Name()
			viper.SetConfigName(dir + "/" + strings.TrimSuffix(fileName, filepath.Ext(fileName)))
			err = viper.MergeInConfig()
			if err != nil {
				panic(err)
			}
		}
	}
}
func ReadConfigsRecursive(searchDir string) {
	e := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && isConfigFile(info.Name()) {
			viper.SetConfigName(path)
			viper.SetConfigType(configType)
			err := viper.MergeInConfig()
			if err != nil {
				panic(err)
			}
		}
		return err
	})
	if e != nil {
		panic(e)
	}
}

func isConfigFile(filename string) bool {
	return strings.HasSuffix(filename, "."+configType)

}
