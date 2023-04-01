package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     					primitive.ObjectID 		`bson:"_id"`
	FirstName  				*string 				`json:"first_name"`
	LastName 				*string 				`json:"last_name"`
	Password 				*string					`json:"password"`
	Email 					*string					`json:"email"`
	Phone 					*string					`json:"phone"`
	Token 					*string					`json:"token"`
	Refresh_Token 			*string					`json:"refresh_token"`
	Created_at				time.Time				`json:"created_at"`
	Updated_at				time.Time				`json:"updated_at"`
	User_id					string					`json:"user_id"`
  }