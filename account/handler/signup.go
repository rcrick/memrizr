package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rcrick/memrizr/account/model"
	"github.com/rcrick/memrizr/account/model/apperrors"
)

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	PassWord string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var req signupReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.PassWord,
	}
	err := h.UserService.SignUp(c, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
}
