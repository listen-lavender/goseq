package dao

import (
	"context"

	"github.com/listen-lavender/goseq/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type hardRegionHandler struct {
	client *mongo.Database
}

func NewHardRegionHandler(client *mongo.Database) *hardRegionHandler {
	return &hardRegionHandler{
		client: client,
	}
}

func (hrh *hardRegionHandler) StoreID() string {
	var o *model.HardRegion
	return o.TableName()
}

func (hrh *hardRegionHandler) AtomicAdd(ctx context.Context, o *model.HardRegion) (*model.HardRegion, error) {
	_, err := hrh.client.Collection(hrh.StoreID()).InsertOne(ctx, o)
	return o, err
}

func (hrh *hardRegionHandler) Find(ctx context.Context, offset uint64, limit int, ftype string, ol []*model.HardRegion, filter func(*model.HardRegion) bool) ([]*model.HardRegion, error) {
	var hrList []*model.HardRegion
	cond := bson.M{}
	cursor, err := hrh.client.Collection(hrh.StoreID()).Find(ctx, cond)
	if err != nil {
		return hrList, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(ctx) {
		hr := model.HardRegion{}
		err := cursor.Decode(&hr)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				println("Query namespace region wrong", "error", err.Error())
			}
			continue
		}

		hrList = append(hrList, &hr)
	}
	return hrList, nil
}

func (hrh *hardRegionHandler) FindByID(ctx context.Context, namespace string) (*model.HardRegion, error) {
	hr := &model.HardRegion{}
	err := hrh.client.Collection(hrh.StoreID()).FindOne(ctx, bson.M{"_id": namespace}).Decode(hr)
	return hr, err
}
