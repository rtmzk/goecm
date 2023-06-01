package constand

import v1 "go-ecm/internal/goecmserver/model/v1"

type processStatus struct {
	Percent string `json:"percent"`
	Message string `json:"message"`
}
type importStatus struct {
	Percent float64 `json:"percent"`
	Message string  `json:"message"`
	Status  string  `json:"status,omitempty"`
}

var Token = ""
var ImagePackageName = ""
var DeployPercent = ""

var ReportItems = []*v1.ReportItems{}

var HeartbeatItem = []string{}

var MainPackageVersion string = "v5.18.0.0"

var DeployProcess = new(processStatus)
var ImportProgress = new(importStatus)
