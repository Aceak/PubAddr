package cli

import (
	"PubAddr/internal/config"
	"PubAddr/internal/logger"
	"PubAddr/internal/server"
	"PubAddr/internal/version"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	cfgPath string
)

var rootCmd = &cobra.Command{
	Use:   "pubaddr",
	Short: "PubAddr - A lightweight public IP query service",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	// 配置文件路径
	rootCmd.Flags().StringVarP(&cfgPath, "config", "c", "./config.yaml", "Config file path")

	// 版本号
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	// 执行前处理
	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Println("PubAddr", version.Version)
			os.Exit(0)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() {
	// 默认 info，之后由配置覆盖
	logger.InitLogger("info")

	cfg, err := config.Load(cfgPath)
	if err != nil {
		logger.Fatal("Failed to load config: %v", err)
	}

	logger.SetLevel(cfg.Server.LogLevel)

	srv, err := server.NewHTTPServer(cfg)
	if err != nil {
		logger.Fatal("Failed to create HTTP server: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()
	logger.Debug("Received shutdown signal...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("Shutdown unexpectly: %v", err)
	}
}
