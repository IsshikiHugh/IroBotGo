package cTypeExplainer

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var explainerUrl string = "https://xwd733f66f.execute-api.us-west-1.amazonaws.com/prod/cdecl_backend?q="

func Explain(sentence string) (string, error) {
	sentence = strings.Replace(sentence, " ", "%20", 1)
	logrus.Info(sentence)
	resp, err := http.Get(explainerUrl + sentence)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ret := string(body)
	if ret == "\"syntax error\"" {
		return "", errors.New("invalid syntax")
	} else if strings.HasPrefix(ret, "<html>") {
		// 502 Bad Gateway
		return "", errors.New("502 bad gateway")
	} else if strings.HasPrefix(ret, "\"declare") {
		ret = strings.TrimPrefix(ret, "\"")
		ret = strings.TrimSuffix(ret, "\"")
		return ret, nil
	}
	return "", errors.New("receive unaccepted result")
}
