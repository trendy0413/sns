package apis

import (
	"encoding/json"
	"net/http"

	//"strings"

	"github.com/gin-gonic/gin"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type MessageForm struct {
	Msg string `json:"msg"`
}

func (s Server) SMS(to, message string) (string, error) {
	if err := validatePhoneNumber(to); err != nil {
		s.Logger.Error(err)
		return "", err
	}

	params := fillCreateMessageParams(to, message)
	resp, err := s.Client.Api.CreateMessage(params)
	if err != nil {
		s.Logger.Errorln("Error sending SMS message: " + err.Error())
		return "", err
	}

	response, _ := json.Marshal(*resp)
	s.Logger.Infoln("Response: " + string(response))
	return string(response), nil
}

func (s Server) SMSCheck(c *gin.Context) {
	to, err := returnPhoneFromCookie(c)
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var message MessageForm
	if err := c.ShouldBind(&message); err != nil {
		s.Logger.Error("couldn't bind JSON")
		c.JSON(http.StatusBadRequest, "bad Format")
		return
	}

	if err := validatePhoneNumber(to); err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	params := fillCreateMessageParams(to, message.Msg)
	resp, err := s.Client.Api.CreateMessage(params)
	if err != nil {
		s.Logger.Errorln("Error sending SMS message: " + err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, _ := json.Marshal(*resp)
	c.JSON(http.StatusOK, "message sent successfully")
	s.Logger.Infoln("Response: " + string(response))
}

func fillCreateMessageParams(to, message string) *twilioApi.CreateMessageParams {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(twilioNumber)
	params.SetBody(message)

	return params
}
