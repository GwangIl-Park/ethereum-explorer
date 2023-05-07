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
	"math/big"
	"os"
	"time"

	server "ethereum-explorer/server"

	"github.com/gin-contrib/cors"
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
    if err != nil {
      logger.Logger.WithError(err).Error("NewLogger Error")
      return err
    }

    cfg, err := config.NewConfig()
    if err != nil {
      logger.Logger.WithError(err).Error("NewConfig Error")
      return err
    }

    db, err := db.NewDB(context.Background(), cfg.MongoUri, "explorer", []string{"blocks", "transactions"})
    if err != nil {
      logger.Logger.WithError(err).Error("NewDB Error")
      return err
    }

    gin := gin.Default()
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    gin.Use(cors.New(config))

    timeout := time.Duration(1) * time.Second

    ethClient, err := ethClient.NewEthClient(cfg)
    if err != nil {
      logger.Logger.WithError(err).Error("NewEthClient Error")
      return err
    }

	  sub, initBlock, err := subscriber.NewSubscriber(ethClient, db)
    if err != nil {
      logger.Logger.WithError(err).Error("NewSubscriber Error")
      return err
    }

    errorChan := make(chan error)
    go sub.ProcessSubscribe(ethClient, db, errorChan)

    go sub.ProcessPrevious(ethClient, db, big.NewInt(cfg.StartBlock), initBlock, errorChan)

    sv := server.NewServer(db, cfg, gin, ethClient, sub, timeout)
    
    routes.Setup(&sv)

    go sv.Start(errorChan)

    err = <-errorChan
    if err != nil {
      logger.Logger.WithError(err).Error("Error")
      return err
    }

    return nil
  },
}

func init() {
  rootCmd.Flags().String("host", "0.0.0.0", "server ip address")
  rootCmd.Flags().String("port", "5000", "server port")
  rootCmd.Flags().String("chainHttp", "http://localhost:8545", "Chain Http Url")
  rootCmd.Flags().String("chainWs", "ws://localhost:8546", "Chain Websocket Url")
  rootCmd.Flags().String("mongoUri", "mongodb://localhost:27017", "Mongo DB URI")
  rootCmd.Flags().Int64("startBlock", 0, "explorer start block")
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