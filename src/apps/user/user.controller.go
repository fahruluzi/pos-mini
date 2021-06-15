package user

import (
	"net/http"

	"github.com/fahruluzi/pos-mini/src/apps/merchant"
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-gonic/gin"
)

// * Register godoc
// @Summary Register User
// @Description register app
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body RegisterValidator true "Register User"
// @Success 200 {object} utils.ResponseModel
// @Router /user/register [post]
func UserCreate(c *gin.Context) {
	response := utils.Response{C: c}
	registerValidator := NewRegisterValidator()

	if err := registerValidator.Bind(c); err != nil {
		response.ResponseFormatter(http.StatusNotAcceptable, "Invalid Form", err, gin.H{"err_message": err.Error()})
		return
	}

	newMerchant := merchant.Merchants{Name: registerValidator.MerchantName, Phone: registerValidator.MerchantPhone, Address: registerValidator.MerchantAddress}
	merchantUUID, resultSaveMerchant := merchant.Save(newMerchant)
	if resultSaveMerchant.Error != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save data", resultSaveMerchant.Error, nil)
		return
	}

	if registerValidator.Password != registerValidator.CorfirmationPassword {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", nil, gin.H{"error": map[string]string{
			"confirm_password": "not match",
		}})
		return
	}

	hashPassword, _ := utils.HashPassword(registerValidator.Password)

	newUser := Users{
		Name:         registerValidator.FullName,
		Email:        registerValidator.Email,
		Password:     hashPassword,
		MerchantUuid: merchantUUID,
	}
	userUUID, resultSaveUser := Save(newUser)
	if resultSaveUser.Error != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save data", resultSaveUser.Error, nil)
		return
	}

	token, err := utils.GenerateJWT(newUser.Name, newUser.Email)
	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Create Token JWT", err, gin.H{
			"message": err.Error(),
		})
		return
	}

	response.ResponseFormatter(http.StatusOK, "Register Successfully", nil, gin.H{
		"userInsertedUuid":     userUUID,
		"merchantInsertedUuid": merchantUUID,
		"access_token":         token,
		"users": map[string]interface{}{
			"email": newUser.Email,
			"name":  newUser.Name,
		},
	})
}

// * Login godoc
// @Summary Login User
// @Description login app
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body LoginValidator true "Login User"
// @Success 200 {object} utils.ResponseModel
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	response := utils.Response{C: c}
	loginValidator := NewLoginValidator()

	if err := loginValidator.Bind(c); err != nil {
		response.ResponseFormatter(http.StatusNotAcceptable, "Invalid Form", err, gin.H{"err_message": err.Error()})
		return
	}

	usersLogin, err := GetUserByEmail(loginValidator.Email)
	if err != nil {
		response.ResponseFormatter(http.StatusNotFound, "Email Unregistered", err, gin.H{"err_message": err.Error()})
		return
	}

	if !utils.CheckPasswordHash(loginValidator.Password, usersLogin.Password) {
		response.ResponseFormatter(http.StatusBadRequest, "Password Wrong", nil, gin.H{"err_message": "Password Wrong"})
		return
	}

	token, err := utils.GenerateJWT(usersLogin.Name, usersLogin.Email)
	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Create Token JWT", err, gin.H{
			"message": err.Error(),
		})
		return
	}

	response.ResponseFormatter(http.StatusOK, "Register Successfully", nil, gin.H{
		"access_token": token,
		"users": map[string]interface{}{
			"email": usersLogin.Email,
			"name":  usersLogin.Name,
		},
	})
}
