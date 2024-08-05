package main

import (
	"context"
	"fmt"
	"gin_fw/src/database"
	"gin_fw/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	godotenv.Load()
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	r := gin.Default()

	routes := route.NewRouteFacade()
	routes.SetupRoutes(r)

	database.Connect()

	// Case single topic
	/*consumer := CreateKafkaConsumer()
	defer consumer.Close()
	topic := "gin-service"
	partition := int32(0)
	if err := StartConsuming(consumer, topic, partition); err != nil {
		log.Fatalf("Error while consuming: %v", err)
	}*/

	// Case consumer group
	groupID := "gin-service"
	topics := []string{"hello"}

	consumerGroup := CreateKafkaConsumerGroup(groupID)
	defer consumerGroup.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	go func() {
		for range sigchan {
			cancel()
		}
	}()

	if err := StartGroupConsuming(consumerGroup, topics, ctx); err != nil {
		log.Fatalf("Error while consuming: %v", err)
	}

	appPort := os.Getenv("APP_PORT")
	address := fmt.Sprintf(":%s", appPort)
	r.Run(address)
}
