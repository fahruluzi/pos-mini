package product

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-gonic/gin"
)

// * GetList godoc
// @Summary Get List Product
// @Description List product
// @Tags Products
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /product [get]
func GetListProduct(c *gin.Context) {
	response := utils.Response{C: c}
	pg, err := utils.GeneratePaginationFromRequest(c)
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Type Queries Pagination", err, gin.H{"err_message": err.Error()})
		return
	}

	products, err := GetProducts(&pg)
	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Fetch Data", err, gin.H{"err_message": err.Error()})

		return
	}

	countProducts, err := CountProducts()
	if err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Count Data", err, gin.H{"err_message": err.Error()})
		return
	}

	countTotalPage := utils.CountTotalPage(countProducts, &pg)

	response.ResponseFormatter(http.StatusOK, "List Products", nil, gin.H{
		"data":       products,
		"total_data": countProducts,
		"total_page": countTotalPage,
		"page":       pg.Page,
	})
}

// * Get Detail godoc
// @Summary Get Detail Product
// @Description Detail product
// @Tags Products
// @Param uuid path string true "Product UUID"
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /product/{uuid} [get]
func GetProductDetail(c *gin.Context) {
	response := utils.Response{C: c}
	getUuid := c.Param("uuid")

	ok := utils.ValidateUuid(getUuid)
	if !ok {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid UUID", nil, nil)
		return
	}

	productDetail, err := GetProduct(getUuid)
	if err != nil {
		response.ResponseFormatter(http.StatusNotFound, "Product Not Found", err, gin.H{"err_message": err.Error()})
		return
	}

	response.ResponseFormatter(http.StatusOK, "Detail Products", nil, gin.H{"data": productDetail})
}

// * Create Product godoc
// @Summary Create Product
// @Description Create product
// @Tags Products
// @Param product_name formData string true "Product Name"
// @Param product_image formData file true "Product Image"
// @Param product_price formData integer true "Product Price"
// @Param product_actual_price formData integer true "Product Actual Price"
// @Param product_stock formData integer true "Product Stock"
// @Accept  application/form-data
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	response := utils.Response{C: c}
	reqData, _ := c.MustGet("my_user_model").(jwt.MapClaims)

	name := c.PostForm("product_name")
	if len(name) <= 0 {
		response.ResponseFormatter(http.StatusBadRequest, "Form Name Required", nil, gin.H{"product_name": "required"})
		return
	}

	price := c.PostForm("product_price")
	intPrice, err := strconv.Atoi(price)
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Form Price must be Integer", err, gin.H{"err_message": err.Error()})
		return
	}

	actualPrice := c.PostForm("product_actual_price")
	intActualPrice, err := strconv.Atoi(actualPrice)
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Form Actual Price must be Integer", err, gin.H{"err_message": err.Error()})
		return
	}

	stock := c.PostForm("product_stock")
	intStock, err := strconv.Atoi(stock)
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Form Stock must be Integer", err, gin.H{"err_message": err.Error()})
		return
	}

	file, err := c.FormFile("product_image")
	if err != nil {
		response.ResponseFormatter(http.StatusBadRequest, "Form Image Required", err, gin.H{"err_message": err.Error()})
		return
	}

	mimeType := strings.Join(file.Header["Content-Type"], "")
	if !strings.Contains(mimeType, "image") {
		response.ResponseFormatter(http.StatusBadRequest, "Form Product Image must be Image file", nil, gin.H{"product_image": "Form Product Image must be Image file"})
		return
	}

	generatedSKU := utils.GenerateSKU(name)
	splittedFileName := strings.Split(file.Filename, ".")
	fileName := generatedSKU + "." + splittedFileName[len(splittedFileName)-1]
	path := "./public/uploads/" + fileName
	if err := c.SaveUploadedFile(file, path); err != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Upload Image", err, gin.H{"err_message": err.Error()})
		return
	}

	newProduct := Products{
		Name:         name,
		SKU:          generatedSKU,
		Price:        int32(intPrice),
		ActualPrice:  int32(intActualPrice),
		Stock:        int32(intStock),
		Image:        "/uploads/" + fileName,
		MerchantUuid: reqData["merchant"].(string),
	}

	insertedUUID, resultSave := Save(newProduct)
	if resultSave.Error != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Create Product", err, gin.H{"err_message": err.Error()})
		return
	}

	response.ResponseFormatter(http.StatusOK, "Successfully Create Product", nil, gin.H{"insertedUUID": insertedUUID})
}

// * Update Product godoc
// @Summary Update Product
// @Description Update product
// @Tags Products
// @Param uuid path string true "Product UUID"
// @Param product_name formData string false "Product Name"
// @Param product_image formData file false "Product Image"
// @Param product_price formData integer false "Product Price"
// @Param product_actual_price formData integer false "Product Actual Price"
// @Param product_stock formData integer false "Product Stock"
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /product/{uuid} [patch]
func UpdateProduct(c *gin.Context) {
	response := utils.Response{C: c}
	productUpdate := make(map[string]interface{})
	getUuid := c.Param("uuid")

	ok := utils.ValidateUuid(getUuid)
	if !ok {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid Path UUID", nil, nil)
		return
	}

	productDetail, err := GetProduct(getUuid)
	if err != nil {
		response.ResponseFormatter(http.StatusNotFound, "Product Not Found", err, gin.H{"err_message": err.Error()})
		return
	}

	name := c.PostForm("product_name")
	if name != "" {
		productUpdate["name"] = name
	}

	price := c.PostForm("product_price")
	if price != "" {
		intPrice, err := strconv.Atoi(price)
		if err != nil {
			response.ResponseFormatter(http.StatusBadRequest, "Form Price must be Integer", err, gin.H{"err_message": err.Error()})
			return
		}
		productUpdate["price"] = intPrice
	}

	actualPrice := c.PostForm("product_actual_price")
	if actualPrice != "" {
		intActualPrice, err := strconv.Atoi(actualPrice)
		if err != nil {
			response.ResponseFormatter(http.StatusBadRequest, "Form Price must be Integer", err, gin.H{"err_message": err.Error()})
			return
		}
		productUpdate["actual_price"] = intActualPrice
	}

	stock := c.PostForm("product_stock")
	if stock != "" {
		intStock, err := strconv.Atoi(stock)
		if err != nil {
			response.ResponseFormatter(http.StatusBadRequest, "Form Price must be Integer", err, gin.H{"err_message": err.Error()})
			return
		}
		productUpdate["stock"] = intStock
	}

	file, err := c.FormFile("product_image")
	if err == nil {
		mimeType := strings.Join(file.Header["Content-Type"], "")
		if !strings.Contains(mimeType, "image") {
			response.ResponseFormatter(http.StatusBadRequest, "Form Product Image must be Image file", nil, gin.H{"product_image": "Form Product Image must be Image file"})
			return
		}

		updateName := name
		if updateName == "" {
			updateName = productDetail.Name
		}
		generatedSKU := utils.GenerateSKU(updateName)
		splittedFileName := strings.Split(file.Filename, ".")
		fileName := generatedSKU + "." + splittedFileName[len(splittedFileName)-1]
		path := "./public/uploads/" + fileName
		if err := c.SaveUploadedFile(file, path); err != nil {
			response.ResponseFormatter(http.StatusInternalServerError, "Failed Upload Image", err, gin.H{"err_message": err.Error()})
			return
		}

		productUpdate["image"] = "/uploads/" + fileName
	}

	productUUID, resultUpdateProduct := Update(productUpdate, getUuid)
	if resultUpdateProduct.Error != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Failed Update data", resultUpdateProduct.Error, nil)
		return
	}

	e := os.Remove("./public" + productDetail.Image)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("%v\n", productUpdate)
	response.ResponseFormatter(http.StatusOK, "Successfully Update Product", nil, gin.H{"updatedUUID": productUUID})
}

// * Get Delete godoc
// @Summary Get Delete Product
// @Description Delete product
// @Tags Products
// @Param uuid path string true "Product UUID"
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Success 200 {object} utils.ResponseModel
// @Router /product/{uuid} [delete]
func DeleteProduct(c *gin.Context) {
	response := utils.Response{C: c}
	getUuid := c.Param("uuid")

	ok := utils.ValidateUuid(getUuid)
	if !ok {
		response.ResponseFormatter(http.StatusBadRequest, "Invalid UUID", nil, nil)
		return
	}

	productDetail, err := GetProduct(getUuid)
	if err != nil {
		response.ResponseFormatter(http.StatusNotFound, "Product Not Found", err, gin.H{"err_message": err.Error()})
		return
	}

	productUUID, result := Delete(getUuid)
	if result.Error != nil {
		response.ResponseFormatter(http.StatusInternalServerError, "Product Not Found", result.Error, gin.H{"err_message": result.Error.Error()})
		return
	}

	e := os.Remove("./public" + productDetail.Image)
	if e != nil {
		log.Fatal(e)
	}

	response.ResponseFormatter(http.StatusOK, "Successfully Delete Product", nil, gin.H{"deletedUUID": productUUID})
}
