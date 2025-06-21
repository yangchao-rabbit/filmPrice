package cmd

import (
	"context"
	"filmPrice/config"
	"filmPrice/internal/apps"
	"filmPrice/pkg/logger"
	"filmPrice/pkg/shutdown"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"

	_ "filmPrice/internal/apps/all"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动后端",
	Long:  `启动后端`,
	Run: func(cmd *cobra.Command, args []string) {
		// 加载配置文件
		config.LoadConfigYaml()

		sqlLog := logger.New(logger.WithDefault(), logger.WithPrefix(" [SQL] "), logger.WithFile(config.Get().Log.To))
		if err := config.Get().Mysql.Init(sqlLog); err != nil {
			panic(err)
		}

		ginLog := logger.New(logger.WithDefault(), logger.WithPrefix(" [GinServer] "), logger.WithFile(config.Get().Log.To))

		if err := apps.InitImplApps(ginLog); err != nil {
			panic(fmt.Sprintf("初始化ImplApps错误：%v", err))
		}
		// 启动Gin Server
		ginServer, err := apps.NewGinServer()
		if err != nil {
			panic(fmt.Sprintf("启动GinServer错误：%v", err))
		}

		srv := &http.Server{
			Addr:    config.Get().App.URL(),
			Handler: ginServer,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// 优雅关闭
		shutdown.NewHook().Close(
			// 关闭http服务
			func() {
				ginLog.Println("Gin Server 关闭中...")
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()

				if err := srv.Shutdown(ctx); err != nil {
					ginLog.Printf("Server shutdown err %v", err.Error())
				}
				ginLog.Println("Gin Server 关闭成功")
			},
			// 关闭cron
			func() {
				ginLog.Println("Cron关闭中...")
				entries := config.GetCron().Entries()
				for _, v := range entries {
					config.GetCron().Remove(v.ID)
				}

				config.GetCron().Stop()
				ginLog.Println("Cron关闭成功")
			},
			// 关闭数据库
			func() {
				sqlDB, _ := config.GetDB().DB()
				if err := sqlDB.Close(); err != nil {
					sqlLog.Printf("关闭SQL连接错误:%v\n", err)
				}
				sqlLog.Println("数据库连接关闭成功")
			})

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
