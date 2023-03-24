package movie

import (
	"context"

	auth "github.com/mantil-io/go-mantil-template/api/auth"
)

type DestroyRequest struct {
	ID    string
	Token string
}

type DestroyResponse struct {
	Message string
}

func (t *Movie) Delete(ctx context.Context, req *DestroyRequest) *DestroyResponse {
	value := auth.Validate(req.Token)
	if value == "valid" {
		t.kv.Delete(req.ID)
		return &DestroyResponse{
			Message: "data deleted",
		}
	} else {
		return &DestroyResponse{
			Message: "token invalid",
		}
	}

}
