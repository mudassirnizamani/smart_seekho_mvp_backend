package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"smart_seekho_mvp/src/data"
	"smart_seekho_mvp/src/models"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOtpFromDb(phoneNumber string) (models.Otp, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var otp models.Otp

	err := data.OtpsCollection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&otp)

	if err != nil {
		return otp, err
	}

	return otp, nil
}

func SetOtpInDb(phoneNumber string) error {
	randomNumber := rand.Intn(9000) + 1000
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := data.OtpsCollection.CountDocuments(ctx, bson.M{"phone_number": phoneNumber})

	if err != nil {
		return errors.New("error occurred while counting otps")
	}

	if count != 0 {
		filter := bson.D{{"phone_number", phoneNumber}}
		update := bson.D{{"$set", bson.D{{"otp", randomNumber}}}}
		_, err := data.OtpsCollection.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return errors.New("error occurred while updating otp")
		}

		return nil
	} else {
		var otp models.Otp = models.Otp{
			ID:          primitive.NewObjectID(),
			PhoneNumber: phoneNumber,
			Otp:         strconv.Itoa(randomNumber),
		}

		_, err := data.OtpsCollection.InsertOne(ctx, otp)

		if err != nil {
			return errors.New("error occurred while updating otp")
		}

		return nil
	}
}

func CheckOtp(phoneNumber, otp string) error {
	var otpModel models.Otp

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data.OtpsCollection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&otpModel)

	if otpModel.Otp == otp {
		return nil
	}

	return errors.New("invalid otp")
}

func CheckPhoneNumberInDB(phonNumber string) (bool, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := data.UsersCollection.CountDocuments(ctx, bson.M{"phone_number": phonNumber})

	if err != nil {
		fmt.Println(err.Error())
		return false, errors.New("error occurred while fetching phone numbers")
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
