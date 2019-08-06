package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"test/routers"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run: func(_ *cobra.Command, _ []string) {
			router := routers.InitRouter()

			srv := &http.Server{
				Addr:    viper.GetString("listen"),
				Handler: router,
			}

			logrus.Info("Server Starting...")

			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					logrus.Fatalf("Server start failed: %s", err)
				}
			}()

			logrus.Infof("Server started at http://%s", viper.GetString("listen"))

			// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
			quit := make(chan os.Signal, 1)

			// kill (no param) default send syscall.SIGTERM
			// kill -2 is syscall.SIGINT
			// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			logrus.Info("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				logrus.Fatalf("Server Shutdown: %v", err)
			}

			logrus.Info("Server exiting")

		},
	}

	return cmd

}
