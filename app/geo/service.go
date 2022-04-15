package geo

import (
	"context"
	"github.com/weplanx/openapi/common"
	"github.com/weplanx/openapi/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	*common.Inject
}

func (x *Service) FindCountries(
	ctx context.Context,
	fields []string,
) (data []model.Country, err error) {
	projection := bson.M{
		"name":   1,
		"iso2":   1,
		"native": 1,
	}
	for _, v := range fields {
		projection[v] = 1
	}
	var cursor *mongo.Cursor
	if cursor, err = x.Db.Collection("countries").
		Find(ctx,
			bson.M{},
			options.Find().SetProjection(projection),
		); err != nil {
		return
	}
	data = make([]model.Country, 0)
	if err = cursor.All(ctx, &data); err != nil {
		return
	}
	return
}

func (x *Service) FindStates(
	ctx context.Context,
	country string,
	fields []string,
) (data []model.State, err error) {
	projection := bson.M{
		"name":       1,
		"state_code": 1,
	}
	for _, v := range fields {
		projection[v] = 1
	}
	var cursor *mongo.Cursor
	if cursor, err = x.Db.Collection("states").
		Find(ctx,
			bson.M{"country_code": country},
			options.Find().SetProjection(projection),
		); err != nil {
		return
	}
	data = make([]model.State, 0)
	if err = cursor.All(ctx, &data); err != nil {
		return
	}
	return
}

func (x *Service) FindCities(
	ctx context.Context,
	country string,
	state string,
	fields []string,
) (data []model.City, err error) {
	projection := bson.M{"name": 1}
	for _, v := range fields {
		projection[v] = 1
	}
	var cursor *mongo.Cursor
	if cursor, err = x.Db.Collection("cities").
		Find(ctx,
			bson.M{"country_code": country, "state_code": state},
			options.Find().SetProjection(projection),
		); err != nil {
		return
	}
	data = make([]model.City, 0)
	if err = cursor.All(ctx, &data); err != nil {
		return
	}
	return
}
