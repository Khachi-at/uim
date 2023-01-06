package service

import (
	"context"
	"fmt"
	"hash/crc32"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
	"github.com/uim/logger"
	"github.com/uim/naming"
	"github.com/uim/naming/consul"
	"github.com/uim/services/service/conf"
	"github.com/uim/services/service/database"
	"github.com/uim/services/service/handler"
	"github.com/uim/wire"
)

type ServerStartOptions struct {
	Config string
}

// NewServerStartCmd creates a new http server command.
func NewServerStartCmd(ctx context.Context, version string) *cobra.Command {
	opts := &ServerStartOptions{}

	cmd := &cobra.Command{
		Use:   "royal",
		Short: "Start a rpc service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx, opts, version)
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.Config, "config", "c", "conf.yaml", "Config file")
	return cmd
}

// RunServerStart run http server.
func RunServerStart(ctx context.Context, opts *ServerStartOptions, version string) error {
	config, err := conf.Init(opts.Config)
	if err != nil {
		return err
	}
	_ = logger.Init(logger.Setting{
		Level:    config.LogLevel,
		FileName: "./data/royal.log",
	})

	// Init database.
	db, err := database.InitMysqlDb(config.BaseDb)
	if err != nil {
		return err
	}
	_ = db.AutoMigrate(&database.Group{}, &database.GroupMember{})

	messageDb, err := database.InitMysqlDb(config.MessageDb)
	if err != nil {
		return err
	}
	_ = messageDb.AutoMigrate(&database.MessageIndex{}, &database.MessageContent{})

	if config.NodeID == 0 {
		config.NodeID = int64(HashCode(config.ServiceID))
	}
	idgen, err := database.NewIDGenerator(config.NodeID)
	if err != nil {
		return err
	}

	rdb, err := conf.InitRedis(config.RedisAddrs, "")
	if err != nil {
		return err
	}
	ns, err := consul.NewNaming(config.ConsulURL)
	if err != nil {
		return err
	}
	_ = ns.Register(&naming.DefaultService{
		Id:       config.ServiceID,
		Name:     wire.SNService,
		Address:  config.PublicAddress,
		Port:     config.PublicPort,
		Protocol: "http",
		Tags:     config.Tags,
		Meta: map[string]string{
			consul.KeyHealthURL: fmt.Sprintf("http://%s:%d/health", config.PublicAddress, config.PublicPort),
		},
	})
	defer func() {
		_ = ns.Deregister(config.ServiceID)
	}()
	serviceHandler := handler.ServiceHandler{
		BaseDb:    db,
		MessageDb: messageDb,
		Idgen:     idgen,
		Cache:     rdb,
	}

	ac := conf.MakeAccessLog()
	defer ac.Close()

	app := newApp(&serviceHandler)
	app.UseRouter(ac.Handler)
	app.UseRouter(setAllowedResponses)

	// start server.
	return app.Listen(config.Listen, iris.WithOptimizations)
}

func newApp(serviceHandler *handler.ServiceHandler) *iris.Application {
	app := iris.Default()

	app.Get("/health", func(ctx iris.Context) {
		_, _ = ctx.WriteString("ok")
	})
	messageAPI := app.Party("/api/:app/message")
	{
		messageAPI.Post("/user", serviceHandler.InsertUserMessage)
		messageAPI.Post("/group", serviceHandler.InsertGroupMessage)
		messageAPI.Post("/ack", serviceHandler.MessageAck)
	}

	groupAPI := app.Party("/api/:app/group")
	{
		groupAPI.Get("/:id", serviceHandler.GroupGet)
		groupAPI.Post("", serviceHandler.GroupCreate)
		groupAPI.Post("/member", serviceHandler.GroupJoin)
		groupAPI.Delete("/member", serviceHandler.GroupQuit)
		groupAPI.Get("/members/:id", serviceHandler.GroupMembers)
	}

	offlineAPI := app.Party("/api/:app/offline")
	{
		offlineAPI.Use(iris.Compression)
		offlineAPI.Post("/index", serviceHandler.GetOfflineMessageIndex)
		offlineAPI.Post("/content", serviceHandler.GetOfflineMessageContent)
	}
	return app
}

func setAllowedResponses(ctx iris.Context) {
	ctx.Negotiation().JSON().Protobuf().MsgPack()
	ctx.Negotiation().Accept.JSON()
	ctx.Next()
}

func HashCode(key string) uint32 {
	hash32 := crc32.NewIEEE()
	hash32.Write([]byte(key))
	return hash32.Sum32() % 1000
}
