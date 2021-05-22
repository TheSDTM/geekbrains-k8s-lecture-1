package internal

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

func GetAppVersion() string {
	return "1.0"
}

// getHostname возвращает хостнейм машины
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}

// getEnvName возвращает перменную окружения SOME_VARIABLE
func GetEnvName() string {
	envVar := os.Getenv("SOME_VARIABLE")
	if len(envVar) == 0 {
		return "unknown"
	}
	return envVar
}

func GetSecret() string {
	data, err := ioutil.ReadFile("/etc/geekbrains/username")
	if err != nil {
		return "unknown"
	}
	return string(data)
}

func GetConfig() string {
	data, err := ioutil.ReadFile("/config/config.yaml")
	if err != nil {
		return "unknown"
	}
	return string(data)
}

func CheckDirExistence() bool {
	if _, err := os.Stat("/data"); !os.IsNotExist(err) {
		return true
	}
	return false
}

func WriteToDataDir() {
	randomFloat := rand.Float32()
	data := fmt.Sprintf("%f\n\r", randomFloat)
	ioutil.WriteFile("/data/data", []byte(data), 0777)
}

func ReadDataDir() string {
	data, err := ioutil.ReadFile("/data/data")
	if err != nil {
		return "unknown"
	}
	return string(data)
}
