package cmd

import (
	"context"
	"ethereum-explorer/config"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/logger"
	"ethereum-explorer/routes"
	"ethereum-explorer/subscriber"
	"fmt"
	"log"
	"os"
	"time"

	server "ethereum-explorer/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbosity string
)

var rootCmd = &cobra.Command{
  Use:   "ethereum-explorer",
  RunE: func(command *cobra.Command, args []string) error {
    err := logger.NewLogger(verbosity)

    cfg := config.NewConfig()

    db, err := db.NewDB(context.Background(), cfg.MongoUri, "explorer", []string{"blocks", "transactions"})
    if err != nil {}

    gin := gin.Default()

    timeout := time.Duration(1) * time.Second

    ethClient := ethClient.NewEthClient(cfg)

	  sub := subscriber.NewSubscriber(ethClient)

    go sub.ProcessSubscribe(ethClient)

    sv := server.NewServer(db, cfg, gin, ethClient, sub, timeout)
    
    routes.Setup(&sv)

    sv.Start()

    return nil
  },
}

func init() {
  rootCmd.Flags().String("host", "0.0.0.0", "server ip address")
  rootCmd.Flags().String("port", "5000", "server port")
  rootCmd.Flags().String("chainUrl", "http://localhost:8545", "Chain Url")
  rootCmd.Flags().String("mongoUri", "mongodb://localhost:27017", "Mongo DB URI")
  rootCmd.Flags().StringVar(&verbosity, "verbosity", "info", "Verbosity Level [debug, info, warn, error]")

  if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}