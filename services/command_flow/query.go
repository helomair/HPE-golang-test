package commandflow

import (
	"HPE-golang-test/services/component"
	"HPE-golang-test/services/models"
	"time"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func queryCommand(userId string) lineSDK.SendingMessage {
	var ret lineSDK.SendingMessage
	// Get Datas
	filter := makeFilter(userId)
	reservations, _ := models.ReservationModel.Query(filter)

	// Make Response
	retMsg := ""
	for _, v := range reservations {
		retMsg += v.ExpireTime.Format(time.RFC1123) + "\n"
		retMsg += v.Content + "\n"
		retMsg += "\n"
	}

	ret = component.NormalMessage(retMsg)

	return ret
}

func makeFilter(userId string) primitive.M {
	filter := bson.M{
		"user_id":     userId,
		"expire_time": bson.M{"$gte": time.Now()},
	}

	return filter
}
