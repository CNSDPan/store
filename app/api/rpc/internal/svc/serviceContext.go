package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"store/app/api/rpc/internal/config"
	"store/app/user/model/sqls"
	"store/pkg/cache"
	"store/pkg/inital"
	"strconv"
)

type ServiceContext struct {
	Config       config.Config
	Node         *snowflake.Node
	CacheConn    *redis.Client
	CacheConnApi *cache.CacheItem
	BizConn      *redis.Client
	UserModel    *sqls.UsersMgr
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 服务uuid节点池
	nodeId, err := strconv.ParseInt(c.ServiceId, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("%s nodeId节点初始化失败 fail:%s", c.ServiceName, err.Error()))
	}
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic(fmt.Sprintf("%s node节点池初始化失败 fail:%s", c.ServiceName, err.Error()))
	}

	cacheConn := inital.NewCacheRedisConn(c.CacheRedis, c.Name)
	bizConn := inital.NewBizRedisConn(c.BizRedis, c.Name)

	return &ServiceContext{
		Config:       c,
		Node:         node,
		CacheConn:    cacheConn,
		CacheConnApi: cache.NewCache(cache.NewCacheUser(context.Background(), cacheConn), cache.NewCacheStore(context.Background(), cacheConn)),
		BizConn:      bizConn,
		UserModel:    sqls.NewUserMgr(inital.NewSqlDB(c.Sql, "source.userModel")),
	}
}
