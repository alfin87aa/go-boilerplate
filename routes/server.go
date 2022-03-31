package routes

import (
	"boilerplate/configs"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
		c.Next()
	}
}

func Init() *gin.Engine {
	r := gin.Default()
	if os.Getenv("APP_ENV") != "PRODUCTION" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	/**
	@description Setup Middleware
	*/
	r.Use(RequestIDMiddleware())
	r.Use(CORSMiddleware())
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.BestCompression))

	r.GET("/healtz", func(c *gin.Context) {
		db := configs.GetDB()
		dbSQL, _ := db.DB()
		redis := configs.GetRedis()

		if err := dbSQL.Ping(); err != nil {
			c.JSON(503, err)
			return
		}

		redisStatus, err := redis.Ping(redis.Context()).Result()
		if err != nil {
			c.JSON(503, err)
			return
		}
		c.JSON(200, gin.H{
			"DB SQL": dbSQL.Stats(),
			"Redis":  redisStatus,
		})
	})

	InitAuthRoutes(r)

	return r
}
