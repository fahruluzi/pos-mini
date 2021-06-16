package product

import (
	"github.com/gin-gonic/gin"
)

type CreateProductValidator struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Price       string `form:"price" json:"price" binding:"required"`
	ActualPrice string `form:"actual_price" json:"actual_price" binding:"required"`
	Stock       string `form:"stock" json:"stock" binding:"required"`
	Image       string `form:"image" json:"image" binding:"required"`
}

func NewCreateProductValidator() CreateProductValidator {
	return CreateProductValidator{}
}

func (createProductValidator *CreateProductValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindJSON(&createProductValidator)
	if err != nil {
		return err
	}

	return
}

type UpdateProductValidator struct {
	FullName    string `form:"full_name" json:"full_name" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required,email"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
}

func NewUpdateProductValidator() UpdateProductValidator {
	return UpdateProductValidator{}
}

func (createProductValidator *UpdateProductValidator) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindJSON(&createProductValidator)
	if err != nil {
		return err
	}

	return
}
