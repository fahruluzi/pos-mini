package user

import (
	"github.com/gin-gonic/gin"
)

type RegisterValidator struct {
	Email                string `form:"email" json:"email" binding:"required,email"`
	Password             string `form:"password" json:"password" binding:"required"`
	CorfirmationPassword string `form:"confirmation_password" json:"confirmation_password" binding:"required"`
	FullName             string `form:"full_name" json:"full_name" binding:"required"`
	MerchantName         string `form:"merchant_name" json:"merchant_name" binding:"required"`
	MerchantPhone        string `form:"merchant_phone" json:"merchant_phone" binding:"required"`
	MerchantAddress      string `form:"merchant_address" json:"merchant_address" binding:"required"`
}

func NewRegisterValidator() RegisterValidator {
	return RegisterValidator{}
}

func (registerValidator *RegisterValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindJSON(&registerValidator)
	if err != nil {
		return err
	}

	return
}

type LoginValidator struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func NewLoginValidator() LoginValidator {
	return LoginValidator{}
}

func (loginValidator *LoginValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindJSON(&loginValidator)
	if err != nil {
		return err
	}

	return
}

type CreateUserValidator struct {
	FullName             string `form:"full_name" json:"full_name" binding:"required"`
	Email                string `form:"email" json:"email" binding:"required,email"`
	Password             string `form:"password" json:"password" binding:"required"`
	CorfirmationPassword string `form:"confirmation_password" json:"confirmation_password" binding:"required"`
}

func NewCreateUserValidator() CreateUserValidator {
	return CreateUserValidator{}
}

func (createUserValidator *CreateUserValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindJSON(&createUserValidator)
	if err != nil {
		return err
	}

	return
}
