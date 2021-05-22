package internal

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

func TestRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	res := rdb.Ping()
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func TestRabbitMQ() error {
	connUrl := fmt.Sprintf("amqp://%s:%s@%s/", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASSWORD"), os.Getenv("RABBIT_HOST"))
	conn, err := amqp.Dial(connUrl)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func TestPostgre() error {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("POSTGRES_HOST"),
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	defer db.Close()
	_, err := db.Exec("SELECT 1=1")
	return err
}
