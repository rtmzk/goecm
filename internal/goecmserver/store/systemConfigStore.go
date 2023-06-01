package store

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type SystemConfigStore interface {
	GetSystemConfig(c context.Context, options metav1.GetOptions) (*v1.SystemConfig, error)
	GetAgentToken(c context.Context, options metav1.GetOptions) (string, error)
}
