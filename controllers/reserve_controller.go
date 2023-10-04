package controllers

import (
	"HPE-golang-test/services/helper"
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

	message := models.UserMessage{
		UserID:     ctx.PostForm("user_id"),
		ReplyToken: ctx.PostForm("reply_token"),
		Content:    ctx.PostForm("reserve_content"),
		ExpireTime: helper.ParseHtmlDateTime(ctx.PostForm("reserve_datetime")),
		CreateTime: time.Now(),
	}

	invalidMsg := validate.Run(message, "struct")
	if len(invalidMsg) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": invalidMsg,
		})
	}

	log.Println(message)
	models.UserModel.Save(message)
}
