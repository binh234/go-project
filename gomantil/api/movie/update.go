package movie

import (
	"context"

	auth "github.com/mantil-io/go-mantil-template/api/auth"
)

type SaveRequest struct {
	Movie MovieItem
	Token string
}

type SaveResponse struct {
	Message string
}

func (t *Movie) Update(ctx context.Context, req *SaveRequest) *SaveResponse {
	value := auth.Validate(req.Token)
	if value == "valid" {
		t.kv.Put(req.Movie.ID, req.Movie)
		return &SaveResponse{
			Message: "data saved",
		}
	} else {
		return &SaveResponse{
			Message: "token invalid",
		}
	}
}
