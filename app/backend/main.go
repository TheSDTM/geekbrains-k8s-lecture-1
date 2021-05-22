package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/TheSDTM/geekbrains-k8s-lecture-1/app/backend/internal"
)

var pvDetected = false
var lock = sync.Mutex{}
var res = &Result{}

type Result struct {
	Hostname   string `json:"hostname"`
	AppVersion string `json:"appVersion"`
	EnvName    string `json:"envName"`
	Secret     string `json:"secret"`
	DataDir    string `json:"dataDir"`
	Config     string `json:"config"`

	RedisRes   string `json:"redisRes"`
	RabbitRes  string `json:"rabbitRes"`
	PostgreRes string `json:"postgreRes"`
}

func (r *Result) Update() {
	r.Hostname = internal.GetHostname()
	r.AppVersion = internal.GetAppVersion()
	r.EnvName = internal.GetEnvName()
	r.Secret = internal.GetSecret()
	r.DataDir = internal.ReadDataDir()
	r.Config = internal.GetConfig()

	r.RedisRes = "connected!"
	redisErr := internal.TestRedis()
	if redisErr != nil {
		r.RedisRes = redisErr.Error()
	}

	r.RabbitRes = "connected!"
	rabbitErr := internal.TestRabbitMQ()
	if rabbitErr != nil {
		r.RabbitRes = rabbitErr.Error()
	}

	r.PostgreRes = "connected!"
	postgreErr := internal.TestPostgre()
	if postgreErr != nil {
		r.PostgreRes = postgreErr.Error()
	}
}

func (r *Result) String() string {
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
	-------------------------
	Redis: %s
	-------------------------
	RabbitMQ: %s
	-------------------------
	PostgreSQL: %s
		`
	output = fmt.Sprintf(
		output,
		r.Hostname,
		r.AppVersion,
		r.EnvName,
		r.Secret,
		pvDetected,
		r.DataDir,
		r.Config,
		r.RedisRes,
		r.RabbitRes,
		r.PostgreRes,
	)
	return output
}

func handleRawPage(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	res.Update()
	fmt.Fprint(w, res.String())
}

func handleApiPage(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	res.Update()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	pvDetected = internal.CheckDirExistence()
	internal.WriteToDataDir()
	http.HandleFunc("/", handleRawPage)
	http.HandleFunc("/api", handleApiPage)
	http.ListenAndServe(":80", nil)
}
