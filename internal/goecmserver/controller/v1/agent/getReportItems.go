package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/constand"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"strings"
)

func (a *AgentController) GetReportItems(c *gin.Context) {
	node := c.Query("n")
	queryType := c.Query("t")

	if node == "" {
		node = "all"
	}
	if queryType == "" {
		queryType = "all"
	}

	switch queryType {
	case "image":
		var res []v1.ImageReportItem
		var temp v1.ImageReportItem
		for _, v := range constand.ReportItems {
			if strings.Contains(v.Host, node) || node == "all" {
				for _, item := range v.Items.ImageItem {
					temp.ImageSummary = item
					temp.Host = v.Host
					res = append(res, temp)
				}
			}
		}
		core.WriteResponse(c, nil, res)
		return
	case "container":
		var res []v1.ContainerReportItem
		var temp v1.ContainerReportItem
		for _, v := range constand.ReportItems {
			if strings.Contains(v.Host, node) || node == "all" {
				for _, item := range v.Items.ContainerItem {
					temp.Container = item
					temp.Host = v.Host
					res = append(res, temp)
				}
			}
		}
		core.WriteResponse(c, nil, res)
		return
		//	case "service":
		//		var res []v1.ServiceReportItem
		//		var temp v1.ServiceReportItem
		//		for _, v := range constand.ReportItems {
		//			if strings.Contains(v.Host, node) || node == "all" {
		//				for _, item := range v.Items.ServiceItem {
		//					temp.Host = v.Host
		//					temp.Service = item
		//					res = append(res, temp)
		//				}
		//			}
		//		}
		//		core.WriteResponse(c, nil, res)
		//		return
	case "all":
		//		var res []v1.AllReportItems
		//		var temp v1.AllReportItems
		//		for _, v := range constand.ReportItems {
		//			if strings.Contains(v.Host, node) || node == "all" {
		//				for _, item := range v.Items.ImageItem {
		//					temp.ImageReportItem.Host = v.Host
		//					temp.ImageReportItem.ImageSummary = item
		//				}
		//			}
		//		}
		//		core.WriteResponse(c, nil, res)
		//		return
	case "*":
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "读取传参失败"), nil)
	}
}
