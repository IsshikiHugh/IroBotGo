package model

type Configuration struct {
	Basic struct {
		QidStr string `yaml:"AccountId"`
		Qid    int64
		Url    string `yaml:"ServerUrl"`
		Retry  int    `yaml:"MaxRetry"`
		Key    string `yaml:"CommandKey"`
	} `yaml:"Basic"`
}
