package route

import (
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/controller"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/repository"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

func NewAddressRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewAddressRepository(db, domain.CollectionAddress)
	tc := &controller.AddressController{
		AddressUsecase: usecase.NewAddressUsecase(tr, timeout),
	}
	group.GET("/address", tc.Fetch)
	group.POST("/address", tc.Create)
}
