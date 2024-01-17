package controller

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressController struct {
	AddressUsecase domain.AddressUsecase
}

// エラー発生箇所のファイル名と行番号を返す関数
func where() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?:0"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (tc *AddressController) Create(c *gin.Context) {
	var address domain.Address

	//addressにmappingする
	err := c.ShouldBind(&address)
	if err != nil {
		where()
		c.JSON(400, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userID := c.GetString("x-user-id")
	address.ID = primitive.NewObjectID()

	address.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.AddressUsecase.Create(c, &address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Address created successfully",
	})

}



func (u *AddressController) Fetch(c *gin.Context) {
	userId := c.GetString("x-user-id")
	address, err := u.AddressUsecase.FetchAddressByUserID(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, address)
}
