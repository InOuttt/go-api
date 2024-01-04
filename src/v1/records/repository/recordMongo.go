package repository

import (
	"context"
	"log"
	"time"

	"github.com/inouttt/test-go-mezink/pkg/db"
	"github.com/inouttt/test-go-mezink/src/v1/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const RecordCollection = "records"

type recordMongo struct {
	DB *db.MongoDB
}

func NewRecordMongo(db *db.MongoDB) domain.RecordRepository {
	return &recordMongo{db}
}

func (rm *recordMongo) GetAll(ctx context.Context, req domain.FetchRecordRequest) (res []domain.Record, err error) {
	ctxTo, cls := context.WithTimeout(ctx, db.MongoMaxTimeExec*time.Second)
	defer cls()

	// pipeline for $match
	qfilter := bson.M{}
	if req.StartDate != "" {
		sd, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, err
		}
		qfilter["createdAt"] = bson.M{"$gte": primitive.NewDateTimeFromTime(sd)}
	}
	if req.EndDate != "" {
		ed, err := time.Parse("2006-01-02T15:04:05", req.EndDate+"T23:59:59")
		if err != nil {
			return nil, err
		}
		qfilter["createdAt"] = bson.M{"$lte": primitive.NewDateTimeFromTime(ed)}
	}

	if req.MinCount != 0 {
		qfilter["totalMarks"] = bson.M{"$gte": req.MinCount}
	}
	if req.MaxCount != 0 {
		qfilter["totalMarks"] = bson.M{"$lte": req.MaxCount}
	}
	pFilter := bson.M{
		"$match": qfilter,
	}

	// pipeline $project
	pProject := bson.M{
		"$project": bson.M{
			"id":   1,
			"name": 1,
			"totalMarks": bson.M{
				"$sum": "$marks",
			},
			"createdAt": 1,
		},
	}

	data, err := rm.DB.Client.Database(rm.DB.DBName).Collection(RecordCollection).Aggregate(ctxTo, bson.A{pProject, pFilter})
	if err != nil {
		log.Println("error on fetch data, ", err)
		return
	}

	if err = data.All(ctx, &res); err != nil {
		log.Println("error on fetch data, ", err)
		return
	}

	return
}
