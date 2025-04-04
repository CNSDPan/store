package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"store/app/rpc/im/client/socket"
	"store/app/rpc/store/internal/config"
	"store/db/dao/query"
	"store/pkg/cache"
	"store/pkg/inital"
	"strconv"
	"time"
)

type ServiceContext struct {
	Config       config.Config
	Node         *snowflake.Node
	CacheConn    *redis.Client
	CacheConnApi *cache.CacheItem
	BizConn      *redis.Client
	Mysql        *gorm.DB
	MysqlQuery   *query.Query
	// 以下RPC服务
	ImRpcCl ImRpc
}

type ImRpc struct {
	Socket socket.Socket
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

	mysqlDb := inital.NewSqlDB(c.Sql, "model")

	sCtx := &ServiceContext{
		Config:       c,
		Node:         node,
		CacheConn:    cacheConn,
		CacheConnApi: cache.NewCache(cache.NewCacheUser(context.Background(), cacheConn), cache.NewCacheStore(context.Background(), cacheConn)),
		BizConn:      bizConn,
		Mysql:        mysqlDb,
		MysqlQuery:   query.Use(mysqlDb),
	}

	go func() {
		sCtx.ImRpcCl = ImRpc{
			Socket: socket.NewSocket(zrpc.MustNewClient(c.ImRpc, zrpc.WithTimeout(time.Second*10))),
		}
	}()
	return sCtx

}
