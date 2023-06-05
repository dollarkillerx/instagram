package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dollarkillerx/graphql_template/internal/conf"
	"github.com/dollarkillerx/graphql_template/internal/generated"
	"github.com/dollarkillerx/graphql_template/internal/middlewares"
	"github.com/dollarkillerx/graphql_template/internal/resolvers"
	"github.com/dollarkillerx/graphql_template/internal/storage/simple"
	"github.com/dollarkillerx/graphql_template/internal/utils"
	"github.com/go-chi/chi/v5"
)

var configFilename string
var configDirs string

func init() {
	const (
		defaultConfigFilename = "config"
		configUsage           = "Name of the config file, without extension"
		defaultConfigDirs     = "./,./configs/"
		configDirUsage        = "Directories to search for config file, separated by ','"
	)
	flag.StringVar(&configFilename, "c", defaultConfigFilename, configUsage)
	flag.StringVar(&configFilename, "config", defaultConfigFilename, configUsage)
	flag.StringVar(&configDirs, "cPath", defaultConfigDirs, configDirUsage)
}

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	utils.InitJWT()
	flag.Parse()

	// Setting up configurations
	err := conf.InitConfiguration(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		panic(fmt.Errorf("Error parsing config, %s", err))
	}

	utils.InitLogger(conf.CONFIG.LoggerConfig)

	if conf.CONFIG.LoggerConfig.Level.IsDebugMode() {
		utils.Logger.Warningln("running on debug mode")
	}

	router := chi.NewRouter()
	router.Use(middlewares.Cors())
	router.Use(middlewares.Context())

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ack"))
	})

	if conf.CONFIG.EnablePlayground {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

		utils.Logger.Warningln("playground is enabled!")
	}

	//将所有的rpcClient放入 gqlgen的Resolver
	newSimple, err := simple.NewSimple(conf.CONFIG.PostgresConfig)
	if err != nil {
		log.Fatalln(err)
	}
	cf := generated.Config{
		Resolvers: resolvers.NewResolver(newSimple),
	}

	cf.Directives.HasLogined = middlewares.HasLoginFunc

	graphQLServer := handler.NewDefaultServer(generated.NewExecutableSchema(cf))

	graphQLServer.SetRecoverFunc(middlewares.RecoverFunc)
	graphQLServer.SetErrorPresenter(middlewares.MiddleError)

	router.Handle("/graphql", graphQLServer)

	utils.Logger.Infof("connect to http://%s/playground for GraphQL playground", conf.CONFIG.ListenAddr)

	if err := http.ListenAndServe(conf.CONFIG.ListenAddr, router); err != nil {
		log.Fatalln(err)
	}
}
