package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func getAppVersion() string {
	return "1.0"
}

// getHostname возвращает хостнейм машины
func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}

// getEnvName возвращает перменную окружения SOME_VARIABLE
func getEnvName() string {
	envVar := os.Getenv("SOME_VARIABLE")
	if len(envVar) == 0 {
		return "unknown"
	}
	return envVar
}

func getSecret() string {
	data, err := ioutil.ReadFile("/etc/geekbrains/username")
	if err != nil {
		return "unknown"
	}
	return string(data)
}

func getConfig() string {
	data, err := ioutil.ReadFile("/config/config.yaml")
	if err != nil {
		return "unknown"
	}
	return string(data)
}

func checkDirExistence() bool {
	if _, err := os.Stat("/data"); !os.IsNotExist(err) {
		return true
	}
	return false
}

func writeToDataDir() {
	randomFloat := rand.Float32()
	data := fmt.Sprintf("%f\n\r", randomFloat)
	ioutil.WriteFile("/data/data", []byte(data), 0777)
}

func readDataDir() string {
	data, err := ioutil.ReadFile("/data/data")
	if err != nil {
		return "unknown"
	}
	return string(data)
}

func main() {
	pvDetected := checkDirExistence()
	writeToDataDir()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname := getHostname()
		appVersion := getAppVersion()
		envName := getEnvName()
		secret := getSecret()
		dataDir := readDataDir()
		config := getConfig()

		output := `Hostname: %s
-------------------------
App version: %s
-------------------------
Environment Variable: %s
-------------------------
Secret: %s
-------------------------
Persistent volume data (persistence: %t):
%s
-------------------------
Config content:
%s
		`
		output = fmt.Sprintf(output, hostname, appVersion, envName, secret, pvDetected, dataDir, config)
		fmt.Fprint(w, output)
	})
	http.ListenAndServe(":80", nil)
}
