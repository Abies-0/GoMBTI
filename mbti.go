package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Group struct {
	Q string `yaml:"Q"`
	A string `yaml:"A"`
	B string `yaml:"B"`
}

func fetch_data(i int) map[string]interface{} {

	databytes, _ := ioutil.ReadFile("qa.yaml")

	var group []Group
	_ = yaml.Unmarshal(databytes, &group)

	res := make(map[string]interface{})
	res["item"] = i
	res["Q"] = group[i].Q
	res["A"] = group[i].A
	res["B"] = group[i].B

	return res
	
}
