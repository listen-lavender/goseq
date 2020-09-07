package dao

import (
	"context"

	"github.com/listen-lavender/goseq/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type hardSegmentHandler struct {
	client *mongo.Database
}

func NewHardSegmentHandler(client *mongo.Database) *hardSegmentHandler {
	return &hardSegmentHandler{
		client: client,
	}
}

func (hsh *hardSegmentHandler) StoreID() string {
	var o *model.HardSegment
	return o.TableName()
}

func (hsh *hardSegmentHandler) AtomicAdd(ctx context.Context, o *model.HardSegment) (*model.HardSegment, error) {
	_, err := hsh.client.Collection(hsh.StoreID()).InsertOne(ctx, o)
	if err != nil {
		println("===== add segment", err.Error())
	}
	return o, err
}

func (hsh *hardSegmentHandler) AtomicUpdate(ctx context.Context, uType string, id uint16, o *model.HardSegment) (*model.HardSegment, error) {

	cond := bson.M{"_id": o.HardSegmentID, "maxSeq": bson.M{"$lt": o.MaxSeq}}
	doc := bson.M{
		"maxSeq": o.MaxSeq,
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	err := hsh.client.Collection(hsh.StoreID()).FindOneAndUpdate(ctx, cond, bson.M{"$set": doc}, &opt).Decode(o)

	if err != nil {
		println("===== update segment", err.Error())
	}

	return o, err
}

func (hsh *hardSegmentHandler) Find(ctx context.Context, offset uint64, limit int, ftype string, ol []*model.HardSegment, filter func(*model.HardSegment) bool) ([]*model.HardSegment, error) {
	var hsList []*model.HardSegment
	cond := bson.M{}
	cursor, err := hsh.client.Collection(hsh.StoreID()).Find(ctx, cond)
	if err != nil {
		return hsList, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(ctx) {
		hs := model.HardSegment{}
		err := cursor.Decode(&hs)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				println("Query seq segment wrong", "error", err.Error())
			}
			continue
		}

		hsList = append(hsList, &hs)
	}
	return hsList, nil
}
