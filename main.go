package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/ksensehq/eventnative/appconfig"
	"github.com/ksensehq/eventnative/appstatus"
	"github.com/ksensehq/eventnative/events"
	"github.com/ksensehq/eventnative/handlers"
	"github.com/ksensehq/eventnative/middleware"
	"github.com/ksensehq/eventnative/storages"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

//some inner parameters
const (
	staticFileDir      = "./web"
	uploaderFileMask   = "*-event-*-20*.log"
	uploaderBatchSize  = 20
	uploaderLoadEveryS = 60
)

var (
	configFilePath = flag.String("cfg", "./eventnative.yaml", "config file path")
)

func readInViperConfig() error {
	flag.Parse()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//custom config
	viper.SetConfigFile(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Custom eventnative.yaml wasn't provided", err)
	}
	return nil
}

//go:generate easyjson -all handlers/static.go useragent/resolver.go
func main() {
	// Setup seed for globalRand
	rand.Seed(time.Now().Unix())

	//Setup default timezone for time.Now() calls
	time.Local = time.UTC

	if err := readInViperConfig(); err != nil {
		log.Fatal("Error while reading viper config: ", err)
	}

	if err := appconfig.Init(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP)
	go func() {
		<-c
		appstatus.Instance.Idle = true
		cancel()
		appconfig.Instance.Close()
		os.Exit(0)
	}()

	destinationsViper := viper.Sub("destinations")

	//override with config from os env
	jsonConfig := viper.GetString("destinations_json")
	if jsonConfig != "" && jsonConfig != "{}" {
		envJsonViper := viper.New()
		envJsonViper.SetConfigType("json")
		if err := envJsonViper.ReadConfig(bytes.NewBufferString(jsonConfig)); err != nil {
			log.Println("Error reading/parsing json config from DESTINATIONS_JSON", err)
		} else {
			destinationsViper = envJsonViper.Sub("destinations")
		}
	}

	tokenizedEventStorages := storages.CreateStorages(ctx, destinationsViper)
	for _, eStorages := range tokenizedEventStorages {
		for _, es := range eStorages {
			appconfig.Instance.ScheduleClosing(es)
		}
	}

	logEventPath := viper.GetString("log.path")
	if !strings.HasSuffix(logEventPath, "/") {
		logEventPath += "/"
	}
	uploader := events.NewUploader(logEventPath+uploaderFileMask, uploaderBatchSize, uploaderLoadEveryS, tokenizedEventStorages)
	uploader.Start()

	router := SetupRouter()

	log.Println("Started listen and server: " + appconfig.Instance.Authority)
	server := &http.Server{
		Addr:              appconfig.Instance.Authority,
		Handler:           router,
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 60,
		IdleTimeout:       time.Second * 65,
	}
	log.Fatal(server.ListenAndServe())
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New() //gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	publicUrl := viper.GetString("server.public_url")

	htmlHandler := handlers.NewPageHandler(staticFileDir, publicUrl, viper.GetBool("server.disable_welcome_page"))
	router.GET("/p/:filename", htmlHandler.Handler)

	staticHandler := handlers.NewStaticHandler(staticFileDir, publicUrl)
	router.GET("/s/:filename", staticHandler.Handler)
	router.GET("/t/:filename", staticHandler.Handler)

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/event", middleware.TokenAuth(handlers.NewEventHandler().Handler))
	}

	router.Use(middleware.Cors)
	return router
}
