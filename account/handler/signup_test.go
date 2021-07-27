package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rcrick/memrizr/account/model/mocks"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Email and password Required", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("SignUp", mock.Anything, mock.Anything).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()
		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(
			gin.H{
				"Email": "",
			},
		)
		assert.Nil(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.Nil(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})
}
