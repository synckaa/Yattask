package tagServices

import (
	"Yattask/dto/tagDTO"
	"Yattask/dto/taskDTO"
	"context"
)

type TagService interface {
	FindOrCreate(ctx context.Context, userId uint, input taskDTO.TaskCreateUpdateRequest) (tagDTO.TagServiceResponse, error)
}
