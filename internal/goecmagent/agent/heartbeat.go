package agent

import (
	"bytes"
	"encoding/json"
	"go-ecm/internal/goecmagent/config"
	"go-ecm/pkg/log"
	"net/http"
	"time"
)

func Heartbeat(c *config.Config) {
	var heartbeatSpec = &heartbeatRequest{
		Token:             c.Agent.Token,
		AgentPort:         c.InsecureServing.BindPort,
		AgentAddr:         c.InsecureServing.BindAddress,
		RemoteServerAddr:  c.Agent.ServerAddr,
		HeartbeatInterval: c.Agent.HeartBeatInterval,
	}

	var buf = &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(heartbeatSpec)

	for {
		tp := &http.Transport{
			DisableKeepAlives: true,
		}
		cli := &http.Client{
			Transport: tp,
		}

		body, _ := json.Marshal(heartbeatSpec)
		req, _ := http.NewRequest(http.MethodPost, "http://"+heartbeatSpec.RemoteServerAddr+"/v1/agent/heartbeat", bytes.NewReader(body))
		req.Close = true
		req.Header.Set("Content-Type", "application/json")
		resp, err := cli.Do(req)
		if err != nil {
			log.Errorf("heart beat to goecm server failed, error: %s", err.Error())
			time.Sleep(heartbeatSpec.HeartbeatInterval * time.Second)
			continue
		} else {
			resp.Body.Close()
			log.Debugf("heart beat to goecm server %s success", heartbeatSpec.RemoteServerAddr)
			time.Sleep(heartbeatSpec.HeartbeatInterval * time.Second)
		}
	}
}
