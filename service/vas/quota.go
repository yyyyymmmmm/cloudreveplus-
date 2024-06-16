package vas

import (
	"github.com/gin-gonic/gin"
	model "github.com/yyyyymmmmm/Test/models"
	"github.com/yyyyymmmmm/Test/pkg/serializer"
)

// GeneralVASService 通用增值服务
type GeneralVASService struct {
}

// Quota 获取容量配额信息
func (service *GeneralVASService) Quota(c *gin.Context, user *model.User) serializer.Response {
	packs := user.GetAvailableStoragePacks()
	return serializer.BuildUserQuotaResponse(user, packs)
}
