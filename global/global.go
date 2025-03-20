package global

import (
	"neocex/v2/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONF  config.Server
	GVA_DB    *gorm.DB
	GVA_REDIS redis.UniversalClient
	GVA_LOG   *zap.Logger
)
