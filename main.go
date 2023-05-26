package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"

	"github.com/trendy0413/sns/src/apis"
)

const (
	accountSidEnv = "A_SID"
	authTokenEnv  = "AUTH_TOKEN"
)

var (
	accountSid string
	authToken  string
)

// @title My API
// @version 1.0
// @description This is Lottery SMS Notification Service API
// @host localhost:3000
// @BasePath /api/v1
// @schemes http
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	accountSid = os.Getenv(accountSidEnv)
	authToken = os.Getenv(authTokenEnv)

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	server := apis.Server{
		Logger: logger,
		Client: twilioClient,
	}

	go func() {
		if err := apis.NewServer(server); err != nil {
			panic(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-interrupt
	logger.Info("Closing the Server")
}
