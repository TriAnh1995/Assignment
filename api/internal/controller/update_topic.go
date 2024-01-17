package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) UpdateTopic(ctx context.Context, updateInfo model.UpdateInfo) ([]string, error) {

	if err := i.checkEmail(ctx, updateInfo.Sender); err != nil {
		return []string{}, err
	}

	if err := i.repo.UpdateTopic(ctx, updateInfo); err != nil {
		return []string{}, middleware.ServerError
	}

	finalList, err := i.repo.GetReceivedUpdateList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, middleware.ServerError
	}
	return finalList, nil
}
