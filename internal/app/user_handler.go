package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"template/internal/model"
	"template/internal/utils"
	"time"
)

func (u handler) RegisterUser(c *gin.Context) {
	customer := new(model.MemberParam)
	if err := c.Bind(customer); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
		return
	}

	validator := validator.New()

	// Validasi struktur data customer
	if err := validator.Struct(customer); err != nil {
		c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
		return
	}

	err := u.User.RegisterCustomer(c, *customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success register user",
	})
	return
}

func (u handler) LoginUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	userInfo, err := u.User.GetUserInfoByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to login",
			Error:    err.Error(),
		})
		return
	}

	if !utils.VerifyPassword(password, userInfo.Password) {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "invalid username/password",
			Error:    "username or password is mismatch",
		})
		return
	}
	claims := &jwtCustomClaims{
		userInfo.Email,
		int64(userInfo.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when generate token",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   accessToken,
	})
	return
}

func (u handler) CreateGathering(c *gin.Context) {
	gathering := new(model.GatheringParam)
	if err := c.Bind(gathering); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
		return
	}

	validator := validator.New()

	// Validasi struktur data customer
	if err := validator.Struct(gathering); err != nil {
		c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or missing claims",
		})
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to claims map",
		})
		return
	}

	userId, ok := claimsMap["id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user id",
		})
		return
	}

	gathering.Creator = int64(userId)

	err := u.Gathering.CreateNewGathering(c, *gathering)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to create gathering",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success create gathering",
	})
	return
}

func (u handler) SendInvitation(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or missing claims",
		})
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to claims map",
		})
		return
	}

	userId, ok := claimsMap["id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user id",
		})
		return
	}

	idStr := c.Param("id")
	gatheringID, _ := strconv.Atoi(idStr)

	err := u.Gathering.SendInvitation(c, int(userId), gatheringID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to send invitation",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success send invitation",
	})
	return
}

func (u handler) ApproveInvitation(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or missing claims",
		})
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to claims map",
		})
		return
	}

	userId, ok := claimsMap["id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user id",
		})
		return
	}

	idStr := c.Param("id")
	gatheringID, _ := strconv.Atoi(idStr)

	err := u.Gathering.ApproveInvitation(c, model.Invitation{
		GatheringID: gatheringID,
		MemberID:    int(userId),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to approve invitation",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success approve invitation",
	})
	return
}

func (u handler) RejectInvitation(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or missing claims",
		})
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to claims map",
		})
		return
	}

	userId, ok := claimsMap["id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user id",
		})
		return
	}

	idStr := c.Param("id")
	gatheringID, _ := strconv.Atoi(idStr)

	err := u.Gathering.RejectInvitation(c, model.Invitation{
		GatheringID: gatheringID,
		MemberID:    int(userId),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to reject invitation",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success reject invitation",
	})
	return
}

func (u handler) UserFriend(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or missing claims",
		})
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to claims map",
		})
		return
	}

	userId, ok := claimsMap["id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user id",
		})
		return
	}

	data, err := u.User.GetUserFriends(c, int(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to reject invitation",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch user friends",
		Data:     data,
	})
	return
}

func (u handler) GetAttendeesInfo(c *gin.Context) {
	idStr := c.Param("id")
	gatheringID, _ := strconv.Atoi(idStr)

	data, err := u.Gathering.GetGatheringInfo(c, gatheringID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get gathering info",
			Error:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success get gathering info",
		Data:     data,
	})
	return
}
