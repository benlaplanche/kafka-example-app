package services_json

import (
	"encoding/json"
	"errors"
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

type Result struct {
	Nodes []string
	Port  int
}

func Parse(vcap_services string) (kafka Result, zookeeper Result, err error) {
	var parent VCAPServices

	if err := json.Unmarshal([]byte(vcap_services), &parent); err != nil {
		// fmt.Println(err)
		return Result{}, Result{}, err
	}
	// fmt.Printf("%+v\n", parent.Service)
	if parent.Service == nil {
		return Result{}, Result{}, errors.New("Error parsing VCAP_SERVICES")
	}

	kafka = Result{
		Nodes: parent.Service[0].Credentials.Kafka.NodeIps,
		Port:  parent.Service[0].Credentials.Kafka.Port,
	}

	zookeeper = Result{
		Nodes: parent.Service[0].Credentials.Zookeeper.NodeIps,
		Port:  parent.Service[0].Credentials.Zookeeper.Port,
	}

	return kafka, zookeeper, err
}
