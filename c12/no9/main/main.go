package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("shibei", err)
		return false
	}
	filepath := "F:/111111.txt"
	ioutil.WriteFile(filepath, data, 0222)
	if err != nil {
		fmt.Println("err", err)
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	filepath := "F:/111111.txt"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("err", err)
		return false
	}
	err = json.Unmarshal(data, &this)
	if err != nil {
		fmt.Println("err", err)
	}
	return true
}
