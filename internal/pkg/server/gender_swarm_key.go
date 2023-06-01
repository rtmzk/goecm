package server

import (
	"context"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"gorm.io/gorm"
)

func GetTokens() error {
	var ctx = context.Background()
	_, err := store.Client().Swarm().GetToken(ctx, "manager", metav1.GetOptions{})
	if err != nil && err == gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
