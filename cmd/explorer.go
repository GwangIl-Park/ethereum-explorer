package cmd

import (
	"ethereum-explorer/config"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
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

var rootCmd = &cobra.Command{
  Use:   "ethereum-explorer",
  RunE: func(command *cobra.Command, args []string) error {
    cfg := config.NewConfig()
    db, err := db.NewDB(cfg)
    if err != nil {}
    defer db.Close()

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

var (
	verbosity string
)

func init() {
  rootCmd.Flags().String("host", "0.0.0.0", "server ip address")
  rootCmd.Flags().String("port", "5000", "server port")
  rootCmd.Flags().String("dbhost", "127.0.0.1", "db ip address")
  rootCmd.Flags().String("dbport", "3306", "db port")
  rootCmd.Flags().String("dbuser", "test_user", "server ip address")
  rootCmd.Flags().String("dbpassword", "1234", "server ip address")
  rootCmd.Flags().String("dbname", "testdb", "server ip address")
  rootCmd.Flags().StringVar(&verbosity, "verbosity", "info", "Verbosity Level [debug, info, warn, error]")
  rootCmd.Flags().String("chainUrl", "http://localhost:8545", "Chain Url")

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