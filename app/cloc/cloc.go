package cloc

import (
	"fmt"
	"log"

	"aira/app/command"
	"aira/config"

	"github.com/valyala/fastjson"
)

type Data struct {
	Files   int `json:"files"`
	Code    int `json:"code"`
	Blank   int `json:"blank"`
	Comment int `json:"comment"`
}

func Analyze(conf config.Config) {
	stdout := command.Exec("cloc", "./", "--json")

	result, err := parse(conf.Codes, stdout)
	if err != nil {
		log.Printf("parse json failed: %+v\n", err)
		return
	}
	fmt.Println(result)
}

func parse(codes []string, data string) (map[string]Data, error) {
	v, err := fastjson.Parse(data)
	if err != nil {
		return nil, err
	}

	result := make(map[string]Data, len(codes))
	for _, code := range codes {
		info := v.Get(code)

		result[code] = Data{
			Files:   info.GetInt("nFiles"),
			Code:    info.GetInt("code"),
			Blank:   info.GetInt("blank"),
			Comment: info.GetInt("comment"),
		}
	}
	return result, nil
}
