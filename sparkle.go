package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apsdehal/go-logger"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
)

var VERSION = "0.2"

var log, _ = logger.New("sparkle", 1, os.Stdout)

const DefaultTarget = "https://sparklebase.com/api/update-host"

var (
	target  = kingpin.Flag("target", "Server URL").Default(DefaultTarget).String()
	verbose = kingpin.Flag("verbose", "Sets the log level to DEBUG").Short('v').Bool()
	token   = kingpin.Arg("token", "Host specific update token").Required().String()
)

func sendTelemetry(metrics map[string]interface{}) (success bool, err error) {
	jsonValue, err := json.Marshal(metrics)

	if err != nil {
		panic("Failed marshaling system metrics")
	}

	baseURL := *target + "/" + *token

	client := &http.Client{}

	req, _ := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonValue))

	req.Header.Set("User-Agent", "Sparkle/"+VERSION)
	req.Header.Set("Content-Type", "application/json")

	log.Debugf("Trying to send data to: %s", *target)

	resp, err := client.Do(req)

	if err != nil {
		log.Errorf("Could not send telemetry to %s", *target)
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		defer resp.Body.Close()

		var data map[string]interface{}

		var errorMessage = "Server did not accept our metrics. "

		if json.NewDecoder(resp.Body).Decode(&data) != nil {
			errorMessage += fmt.Sprintf("but we could not decode the error message. HTTP-Code %d", resp.StatusCode)
		} else {
			errorMessage += fmt.Sprintf("HTTP-Code %d: %v", resp.StatusCode, data["error"])
		}

		log.Error(errorMessage)
	}

	return true, nil
}

func main() {
	kingpin.Version(VERSION)
	kingpin.Parse()

	if *verbose {
		log.SetLogLevel(logger.DebugLevel)
	}

	info, err := host.Info()

	if err != nil {
		log.Error("Failed collecting Host info")
	}

	virtualMem, err := mem.VirtualMemory()

	if err != nil {
		log.Warning("Failed collecting Memory info")
	}

	values := map[string]interface{}{
		"host":     info,
		"vmem":     virtualMem,
		"_version": "1",
	}

	sendTelemetry(values)
}
