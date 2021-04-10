package crud

import (
	"context"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/model"
)

type CrudUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataResponse, err error)
}
