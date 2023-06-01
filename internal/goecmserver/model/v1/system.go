package v1

import metav1 "go-ecm/internal/pkg/meta/v1"

type SystemConfig struct {
	metav1.ObjectMeta `json:"meta"`
	Token             string `json:"agentToken"`
}
