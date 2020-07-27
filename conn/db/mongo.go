package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(dsn string, dataDB string, authDB string, username string, password string) (*mongo.Database, error) {
	ctx := context.Background()
	// writeTimeout := cfg.GetDuration(fmt.Sprintf("mongo.%s.writetimeout", mark))
	// readTimeout := cfg.GetDuration(fmt.Sprintf("mongo.%s.readtimeout", mark))
	auth := options.Credential{
		AuthSource: authDB,   // cfg.GetString(fmt.Sprintf("mongo.%s.authdb", mark)),
		Username:   username, // cfg.GetString(fmt.Sprintf("mongo.%s.authuser", mark)),
		Password:   password, // cfg.GetString(fmt.Sprintf("mongo.%s.authpass", mark)),
	}

	conn, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(dsn).SetAuth(auth),
		// options.Client().SetMaxPoolSize(100),
		// options.Client().SetMaxConnIdleTime(3600*time.Second),
	)
	if err != nil {
		panic(fmt.Sprintf("mongo connect failed error: %s", err.Error()))
		return nil, nil
	}
	db := conn.Database(dataDB)
	return db, err
}
