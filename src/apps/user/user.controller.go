package user

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

	token, err := utils.GenerateJWT(newUser.Name, newUser.Email, newUser.MerchantUuid)
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

	token, err := utils.GenerateJWT(usersLogin.Name, usersLogin.Email, usersLogin.MerchantUuid)
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

// * GetList godoc
// @Summary Get List User
// @Description List user
// @Tags Users
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /user [get]
func GetListUser(c *gin.Context) {
	response := utils.Response{C: c}
	pg, err := utils.GeneratePaginationFromRequest(c)
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Type Queries Pagination", err, gin.H{"err_message": err.Error()})
		return
	}

	users, err := GetUsers(&pg)
	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Fetch Data", err, gin.H{"err_message": err.Error()})

		return
	}

	countUsers, err := CountUsers()
	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Count Data", err, gin.H{"err_message": err.Error()})
		return
	}

	countTotalPage := utils.CountTotalPage(countUsers, &pg)

	response.ResponseFormatter(http.StatusOK, "List Majors", nil, gin.H{
		"data":       users,
		"total_data": countUsers,
		"total_page": countTotalPage,
		"page":       pg.Page,
	})
}

// * Get Detail godoc
// @Summary Get Detail User
// @Description Detail user
// @Tags Users
// @Param uuid path string true "User UUID"
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /user/{uuid} [get]
func GetUserDetail(c *gin.Context) {
	response := utils.Response{C: c}
	getUuid := c.Param("uuid")

	ok := utils.ValidateUuid(getUuid)
	if !ok {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid UUID", nil, nil)
		return
	}

	userDetail, err := GetUser(getUuid)
	if err != nil {
		response.ResponseFormatter(http.StatusNotFound, "Major Not Found", err, gin.H{"err_message": err.Error()})
		return
	}

	response.ResponseFormatter(http.StatusOK, "List Majors", nil, gin.H{"data": userDetail})
}

// * Create User godoc
// @Summary Create User
// @Description Create user
// @Tags Users
// @Param user body CreateUserValidator true "Create User"
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /user [post]
func CreateUser(c *gin.Context) {
	response := utils.Response{C: c}
	createUserValidator := NewCreateUserValidator()
	reqData, _ := c.MustGet("my_user_model").(jwt.MapClaims)

	if err := createUserValidator.Bind(c); err != nil {
		response.ResponseFormatter(http.StatusNotAcceptable, "Invalid Form", err, gin.H{"err_message": err.Error()})
		return
	}

	if createUserValidator.Password != createUserValidator.CorfirmationPassword {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Form", nil, gin.H{"error": map[string]string{
			"confirm_password": "not match",
		}})
		return
	}

	hashPassword, _ := utils.HashPassword(createUserValidator.Password)

	newUser := Users{
		Name:         createUserValidator.FullName,
		Email:        createUserValidator.Email,
		Password:     hashPassword,
		MerchantUuid: reqData["merchant"].(string),
	}

	ok := utils.ValidateUuid(newUser.MerchantUuid)
	if !ok {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Merchant UUID", nil, nil)
		return
	}

	_, err := merchant.GetMerchant(newUser.MerchantUuid)
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Merchant UUID", nil, nil)
		return
	}

	userUUID, resultSaveUser := Save(newUser)
	if resultSaveUser.Error != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Save data", resultSaveUser.Error, nil)
		return
	}

	response.ResponseFormatter(http.StatusOK, "Successfully Create User", nil, gin.H{"insertedUUID": userUUID})
}
