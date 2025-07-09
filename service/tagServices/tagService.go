package tagServices

import (
	"Yattask/dto/tagDTO"
	"context"
)

type TagService interface {
	FindOrCreate(ctx context.Context, task tagDTO.TagServiceRequest) (tagDTO.TagServiceResponse, error)
}
