package model

type Configuration struct {
	Basic struct {
		Qid   int64  `yaml:"AccountId"`
		Url   string `yaml:"ServerUrl"`
		Retry int    `yaml:"MaxRetry"`
		Key   string `yaml:"CommandKey"`
	} `yaml:"Basic"`
}
