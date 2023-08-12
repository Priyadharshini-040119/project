package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/request", handleRequest)
}

func handleRequest(c *gin.Context) {

	var requestData RequestData

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output := request(RequestData)
	c.JSON(http.StatusOK, output.Data)
}

func request(requestData resources.RequestData) resources.ResponseData{
	ch := make(chan RequestData)

	ch <- requestData

	go worker(ch)

	ResponseData := <-ch

	return ResponseData
}

func worker(data <-chan RequestData) {
	for requestData := range data {
		ResponseData := ResponseData{
			Event:           requestData.Ev,
			EventType:       requestData.Et,
			AppID:           requestData.ID,
			UserID:          requestData.UID,
			MessageID:       requestData.MID,
			PageTitle:       requestData.T,
			PageURL:         requestData.P,
			BrowserLanguage: requestData.L,
			ScreenSize:      requestData.SC,
			Attributes: Attribute{
				FormVarient: Common{
					Value: requestData.Atrv1,
					Type:  requestData.requestD,
				},
				Ref: Common{
					Value: requestData.Atrv2,
					Type:  requestData.Atrt2,
				},
			},
			Traits: Trait{
				Name: Common{
					Value: requestData.UAtrv1,
					Type:  requestData.UAtrt1,
				},
			},
			Email: Common{
				Value: requestData.UAtrv1,
				Type:  requestData.UAtrt1,
			},
			Age: Common{
				Value: requestData.UAtrv1,
				Type:  requestData.UAtrt1,
			},
		}
	}
	data <- ResponseData
}
