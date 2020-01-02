package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var VERSION = "0.1"

var (
	target = kingpin.Flag("target", "Server URL").Default("https://sparklebase.com/api/update-host").String()
	token  = kingpin.Arg("token", "Host specific update token").Required().String()
)

func sendTelemetry(metrics map[string]interface{}) (success bool, err error) {
	jsonValue, _ := json.Marshal(metrics)

	baseURL := *target + "/" + *token
	_, err = http.Post(baseURL, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Errorf("An error occured while sending telemetry: ", err)
		return false, err
	}

	return true, nil
}

func main() {
	kingpin.Version(VERSION)
	kingpin.Parse()

	info, _ := host.Info()
	virtualMem, _ := mem.VirtualMemory()

	values := map[string]interface{}{
		"host": info,
		"vmem": virtualMem,
	}

	sendTelemetry(values)
}
