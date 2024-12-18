package svc

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"store/app/chat/rpc/chat/socket"
	"store/app/store/model/sqls"
	"store/app/store/rpc/internal/config"
	"store/pkg/cache"
	"store/pkg/inital"
	"strconv"
)

type ServiceContext struct {
	Config           config.Config
	Node             *snowflake.Node
	CacheConn        *redis.Client
	CacheConnApi     *cache.CacheItem
	BizConn          *redis.Client
	StoreModel       *sqls.StoresMgr
	StoreUserModel   *sqls.StoreUsersMgr
	StoreMemberModel *sqls.StoresMemberMgr
	ChatLogModel     *sqls.ChatLogMgr
	// 以下RPC服务
	ChatRpcCl ChatRpc
}

type ChatRpc struct {
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

	// chaRPC节点
	chatRpcConf := zrpc.MustNewClient(c.ChatRpc)
	chatRpc := ChatRpc{
		Socket: socket.NewSocket(chatRpcConf),
	}

	return &ServiceContext{
		Config:           c,
		Node:             node,
		CacheConn:        cacheConn,
		CacheConnApi:     cache.NewCache(cache.NewCacheUser(context.Background(), cacheConn), cache.NewCacheStore(context.Background(), cacheConn)),
		BizConn:          bizConn,
		StoreModel:       sqls.NewStoresMgr(inital.NewSqlDB(c.Sql, "storeModel")),
		StoreUserModel:   sqls.NewStoreUsersMgr(inital.NewSqlDB(c.Sql, "storeUserModel")),
		StoreMemberModel: sqls.NewStoresMemberMgr(inital.NewSqlDB(c.Sql, "storeMemberModel")),
		ChatLogModel:     sqls.NewChatLogMgr(inital.NewSqlDB(c.Sql, "chatLogModel")),
		ChatRpcCl:        chatRpc,
	}
}
