package controller

import "context"

func (i CTRLImplement) checkEmail(ctx context.Context, userEmail string) error {
	checkEmailExist, err := i.repo.CheckUserByEmail(ctx, userEmail)
	if err != nil {
		return ServerError
	}
	if !checkEmailExist {
		return UserNotFound
	}
	return nil
}
