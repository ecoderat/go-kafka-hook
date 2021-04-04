package main

import (
	"context"
	"encoding/json"
	"fmt"

	models "github.com/ecoderat/go-kafka-hook/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
)

const (
	topic = "test"
	//brokerAddress = "172.17.0.1:9092"
	brokerAddress = "kafka"

	//host     = "172.17.0.1"
	host     = "kartaca-postgres"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "test"
)

func consume(ctx context.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&models.Logs{})

	/*___________________________________*/

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	/*
		BUNU SAKIN DEĞİŞTİRME!! 3 SAAT ARADIĞIN AUTO.OFFSET.RESET İN KARŞILIĞI
		LastOffset  int64 = -1
		FirstOffset int64 = -2
	*/
	r.SetOffset(-1)

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		var jsonBlob = msg.Value
		var logs models.Logs
		err = json.Unmarshal(jsonBlob, &logs)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("received:", logs)
		err = db.Create(&logs).Error
		if err != nil {
			panic(err)
		}
	}
}
