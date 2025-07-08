package db

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/looker/v1"
)

func InitLookerService(credentialsFile string) (*looker.Service, error) {
	ctx := context.Background()
	lookerService, err := looker.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}
	return lookerService, nil
}