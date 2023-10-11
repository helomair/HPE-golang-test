package controllers

import (
	"HPE-golang-test/services/component"
	"HPE-golang-test/services/helper"
	"HPE-golang-test/services/line"
	"HPE-golang-test/services/models"
	"HPE-golang-test/services/validate"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MakeReserveForm(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"reserve_form.tmpl",
		gin.H{
			"title":       "Reserve Form",
			"form_action": "/reserve",
			"user_id":     ctx.Param("user_id"),
			"reply_token": ctx.Param("reply_token"),
		},
	)
}

func ReserveNew(ctx *gin.Context) {

	reserve := models.Reserve{
		UserID:     ctx.PostForm("user_id"),
		ReplyToken: ctx.PostForm("reply_token"),
		Content:    ctx.PostForm("reserve_content"),
		ExpireTime: helper.ParseHtmlDateTime(ctx.PostForm("reserve_datetime")),
		CreateTime: time.Now(),
	}

	invalidMsg := validate.Run(reserve, "struct")
	if len(invalidMsg) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": invalidMsg,
		})
		return
	}

	log.Println(reserve)
	id, _ := models.ReservationModel.Save(reserve)

	lineService := line.GetLineBotServiceInstance()
	lineService.Push(line.LineEventMessageHandler{
		UserId:     reserve.UserID,
		ReplyToken: reserve.ReplyToken,
		Message:    component.NormalMessage("New Reservation!\nExpire: " + reserve.ExpireTime.Format(time.RFC1123) + "\nContent: " + reserve.Content),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "New reserve done",
		"reserve_id": id,
	})
}

func ReserveQuery(ctx *gin.Context) {
	var reservations []models.Reserve
	var err error

	userId := ctx.Param("user_id")

	if len(userId) == 0 { // query all
		reservations, err = models.ReservationModel.QueryAll()
	} else {
		reservations, err = models.ReservationModel.QueryByUser(userId)
	}
	if err != nil {
		log.Println("MessageQuery failed : " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Message query failed, error msg : " + err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		reservations,
	)
}
