package agent

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"reflect"
	"runtime"

	"github.com/lena-zima/golang-metrics-project/global"
)

func UpdateMetrics() {
	runtime.ReadMemStats(&global.RuntimeMetrics)
	global.PollCount++
	global.RandomValue = rand.Float64()
}

func SendMetrics() {
	for k, v := range global.RtMetrics {
		val := reflect.ValueOf(global.RuntimeMetrics).FieldByName(k)
		sendMetric(v, k, val)
	}

	PollCountvalue := reflect.ValueOf(global.PollCount)
	sendMetric("counter", "PollCount", PollCountvalue)

	RandomValuevalue := reflect.ValueOf(global.RandomValue)
	sendMetric("gauge", "RandomValue", RandomValuevalue)
}

func sendMetric(mtype string, mname string, mvalue reflect.Value) {

	client := &http.Client{}

	url := fmt.Sprint(global.ServerAddr, "/update/", mtype, "/", mname, "/", mvalue)

	//fmt.Println(url)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
}
