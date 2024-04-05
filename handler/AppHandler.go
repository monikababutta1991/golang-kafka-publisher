package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"publisher/service"

	"github.com/gin-gonic/gin"
)

type appHandler struct{}

type IAppHandler interface {
	ProduceMessage() gin.HandlerFunc
}

type InputData struct {
	MAC          string `json:"mac"`
	ProjectID    int    `json:"project_id"`
	SessionID    int    `json:"session_id"`
	FieldAnother string `json:"other_fields"`
}

func NewAppHandler() IAppHandler {
	return &appHandler{}
}

func (a *appHandler) ProduceMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		kafkaHost := os.Getenv("KAFKA_HOST")
		kafkaTopic := os.Getenv("KAFKA_TOPIC")

		var postData InputData

		if err := ctx.ShouldBindJSON(&postData); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		jsonresp, _ := json.Marshal(postData)

		kafkaSvc := service.NewKafkaService(kafkaHost, kafkaTopic)

		go kafkaSvc.ProduceMessage(string(jsonresp))

		ctx.JSON(http.StatusOK, gin.H{"msg": "message published successfully"})
	}
}
