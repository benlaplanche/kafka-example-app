package main

import (
	"encoding/json"
	"fmt"
)

type VCAPServices struct {
	Service []Details `json:"p-kafka"`
}

type Details struct {
	Label       string   `json:"label"`
	Name        string   `json:"name"`
	Plan        string   `json:"plan"`
	Tags        []string `json:"tags"`
	Credentials Process  `json:"credentials"`
}

type ConnectionDetails struct {
	NodeIps []string `json:"node_ips"`
	Port    int      `json:"port"`
}

type Process struct {
	Kafka     ConnectionDetails `json:"kafka"`
	Zookeeper ConnectionDetails `json:"zookeeper"`
}

func main() {
	env := `{"p-kafka":[{"credentials":{"kafka":{"node_ips":["10.244.9.2","10.244.9.6","10.244.9.10","10.244.9.14","10.244.9.18"],"port":9092},"zookeeper":{"node_ips":["10.244.9.22","10.244.9.26","10.244.9.30"],"port":2181}},"label":"p-kafka","name":"kafka","plan":"shared","tags":["pivotal","kafka"]}]}`

	var parent VCAPServices
	if err := json.Unmarshal([]byte(env), &parent); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parent.Service[0].Credentials.Kafka.NodeIps[0])

}
