package repository

import (
	"context"
	"github.com/asdamwongmantap/api-gomongo-crud/crud"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/model"
	mongoDB "github.com/asdamwongmantap/api-gomongo-crud/lib/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type CrudRepository struct {
	mongoDB mongoDB.Mongo
}

func NewCrudRepository(mongo mongoDB.Mongo) crud.CrudRepositoryI {
	return &CrudRepository{
		mongoDB: mongo,
	}
}

func (cr CrudRepository) GetAllData(ctx context.Context) (crudResp model.GetDataResponse, err error) {

	errPing := cr.mongoDB.Ping()
	if errPing != nil {
		log.Println("errorping", errPing)
	}

	query, err := cr.mongoDB.DB().Collection("product").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.GetDataResponse{}, err
	}
	defer query.Close(ctx)

	listDataProduct := make([]model.DataProduct, 0)
	for query.Next(ctx) {
		var row model.DataProduct
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataProduct = append(listDataProduct, row)
	}

	crudResp = model.GetDataResponse{Data: listDataProduct}

	return crudResp, err
}
