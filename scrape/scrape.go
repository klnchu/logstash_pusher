package scrape

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/common/log"
)

const (
	jobName = "logstash"
)

// IntervalScrape interval scrape from http and push
// endpoint as push gateway endpoint, like http://pushgateway.simple.com
// interval as interval scrape logstash metric
func IntervalScrape(endpoint string, intervel int, metricPort string) {

	duration := time.Duration(int64(intervel))
	ticker := time.NewTicker(duration * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			scrapeAndPush(endpoint, metricPort)
		}
	}
}

func scrapeAndPush(endpoint string, metricPort string) {

	body, err := scrape(metricPort)
	if err != nil {
		log.Errorf("Cannot get metric from local server: %v", err)
		return
	}
	err = push(endpoint, body)
	if err != nil {
		log.Errorf("Cannot push metric data to pushgateway: %v", err)
	}
}

func scrape(metricPort string) ([]byte, error) {

	url := fmt.Sprintf("http://localhost%s/metrics", metricPort)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getIntranetIP() (string, error) {

	addrLst, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrLst {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("not found ip address")
}

func push(endpoint string, body []byte) error {

	instance, err := getIntranetIP()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/metrics/job/%s/instance/%s", endpoint, jobName, instance)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
