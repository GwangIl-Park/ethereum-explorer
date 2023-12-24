package cmd

import (
	"context"
	"ethereum-explorer/config"
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/logger"
	"ethereum-explorer/repository"
	"ethereum-explorer/router"
	"ethereum-explorer/service"
	"ethereum-explorer/subscriber"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	server "ethereum-explorer/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbosity string
)

var rootCmd = &cobra.Command{
	Use: "ethereum-explorer",
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
		
		db, err := db.NewDB(context.Background(), *cfg, []string{"blocks", "transactions"})
		if err != nil {
			logger.Logger.WithError(err).Error("NewDB Error")
			return err
		}
		defer db.Close()

		timeout := time.Duration(1) * time.Second

		ethClient, err := ethClient.NewEthClient(cfg)
		if err != nil {
			logger.Logger.WithError(err).Error("NewEthClient Error")
			return err
		}
		defer ethClient.Http.Close()
		defer ethClient.Ws.Close()

		errorChan := make(chan error)
		initBlockNumberChan := make(chan *big.Int)

		sub, err := subscriber.NewSubscriber(ethClient, db, errorChan)
		if err != nil {
			logger.Logger.WithError(err).Error("NewSubscriber Error")
			return err
		}

		go sub.ProcessSubscribe(ethClient, initBlockNumberChan)

		initBlock := <-initBlockNumberChan

		go sub.ProcessPrevious(ethClient, db, initBlock)
		
		sv := server.NewServer(db, cfg, ethClient, sub, timeout)

		r := http.NewServeMux()

		accountRepository := repository.NewAccountRepository(db)
		blockRepository := repository.NewBlockRepository(db)
		mainRepository := repository.NewMainRepository(db)
		transactionRepository := repository.NewTransactionRepository(db)

		accountService := service.NewAccountService(accountRepository)
		blockService := service.NewBlockService(blockRepository)
		mainService := service.NewMainService(mainRepository)
		transactionService := service.NewTransactionService(transactionRepository)

		accountController := controller.NewAccountController(accountService)
		blockController := controller.NewBlockController(blockService)
		mainController := controller.NewMainController(mainService)
		transactionController := controller.NewTransactionController(transactionService)

		router.NewAccountRouter(sv.Timeout, accountController, r)
		router.NewBlockRouter(sv.Timeout, blockController, r)
		router.NewMainRouter(sv.Timeout, mainController, r)
		router.NewTransactionRouter(sv.Timeout, transactionController, r)

		go sv.Start(errorChan, r)

		err = <-errorChan
		if err != nil {
			logger.Logger.WithError(err).Error("Error")
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().String("url", "localhost:5000", "server url")
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
