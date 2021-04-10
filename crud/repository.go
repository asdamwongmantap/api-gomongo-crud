package crud

import (
	"context"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/model"
)

type CrudRepositoryI interface {
	GetAllData(ctx context.Context) (crudResp model.GetDataResponse, err error)
}
