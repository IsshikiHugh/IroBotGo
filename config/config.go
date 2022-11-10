package config

import (
	"IroBot/model"
	"io/ioutil"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var localConfig model.Configuration

func Init(filename string) {
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
		logrus.Fatal("Config initialization failed while read yaml file.")
	} else if err = yaml.Unmarshal(yamlFile, &localConfig); err != nil {
		logrus.Error(err)
		logrus.Fatal("Config initialization failed while unmarshal yaml file.")
	}
	localConfig.Basic.Qid, _ = strconv.ParseInt(localConfig.Basic.QidStr, 10, 64)
}

func Config() *model.Configuration {
	return &localConfig
}
