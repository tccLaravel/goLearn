package main

import (
	"encoding/json"
	"bytes"
	"fmt"
)


type peerInfo struct {
	HTTPPort string `json:"http_port"`
	TCPPort  int `json:"tcp_port,omitempty"`
	Name string `json:"name,omitempty"`
	versiong string
}

func main() {
	b := []byte(`{"http_port":"polaris","tcp_port":30,"versiong":"0.10","name":""}`)

	//var person = make(map[string]interface{})
	person := new(peerInfo)
	decoder := json.NewDecoder(bytes.NewReader(b))

	decoder.UseNumber()
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Printf("json unmarshal error: %+v \n", err)
	}
	fmt.Printf(" %+v \n",person)
}
