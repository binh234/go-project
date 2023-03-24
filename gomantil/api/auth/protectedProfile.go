package auth

import (
	"context"
)

type ProfileRequest struct {
	Token string
}

func (t *Token) ProtectedProfile(ctx context.Context, req *ProfileRequest) string {
	value := Validate(req.Token)

	if value == "valid" {
		return "let's now proceed"
	}
	return value
}
