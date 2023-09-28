package controllers

import (
	"context"
	"log"
	"net/http"
	"smart_seekho_mvp/src/data"
	"smart_seekho_mvp/src/dtos"
	"smart_seekho_mvp/src/models"
	"smart_seekho_mvp/src/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// save the latest in otp with phone to make sure the otp is matches with database
// send and then verify the otp
// get the phone number
// check if a user with phone number already exists or not?
// if the user doesn't exist create an account
// if the user already exist sign in the user and send tokens
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto dtos.AuthDto

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "error ocurred while binding json"})
			return
		}

		err := services.CheckOtp(*dto.PhoneNumber, *dto.Otp)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid otp"})
			return
		}

		numberStatus, err := services.CheckPhoneNumberInDB(*dto.PhoneNumber)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if numberStatus == true {
			var user models.User

			data.UsersCollection.FindOne(ctx, bson.M{"phone_number": dto.PhoneNumber}).Decode(&user)

			c.JSON(http.StatusOK, gin.H{"user": user})
			return
		} else {
			var user models.User = models.User{
				ID:          primitive.NewObjectID(),
				PhoneNumber: dto.PhoneNumber,
				Name:        "",
				ProfilePic:  "",
			}
			user.UserId = user.ID.Hex()

			_, err := data.UsersCollection.InsertOne(ctx, user)

			if err != nil {
				log.Panic(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
			return
		}
	}
}

func GenerateOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto dtos.AuthDto

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "error ocurred while binding json"})
			return
		}

		err := services.SetOtpInDb(*dto.PhoneNumber)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "opt has been created"})
		return
	}
}
