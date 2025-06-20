package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var clienteRedis *redis.Client

func init() {
	clienteRedis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	_, err := clienteRedis.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error al conectar con Redis: %v", err)
	}
}

func GuardarEnCache(nombre string) {
	clienteRedis.Set(ctx, nombre, "1", time.Minute*10)
}

func ObtenerDeCache(nombre string) string {
	val, err := clienteRedis.Get(ctx, nombre).Result()
	if err != nil {
		return ""
	}
	return val
}
