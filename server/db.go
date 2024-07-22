package server

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Pool *redis.Pool
