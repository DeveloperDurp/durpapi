package controller

import (
	"fmt"
	"gitlab.com/DeveloperDurp/DurpAPI/docs"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/dadjoke"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/openai"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.com/DeveloperDurp/DurpAPI/middleware"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/health"
	"gitlab.com/developerdurp/durpify/handlers"
)

type Controller struct {
	Cfg   Config
	Dbcfg DBConfig
	Db    *gorm.DB
}

type Config struct {
	Host      string `env:"host"`
	Version   string `env:"version"`
	Groupsenv string `env:"groupsenv"`
	JwksURL   string `env:"jwksurl"`
	LlamaURL  string `env:"llamaurl"`
	LlamaPass string `env:"llamapass"`
}

type DBConfig struct {
	Host     string `env:"db_host"`
	Port     string `env:"db_port"`
	Password string `env:"db_pass"`
	User     string `env:"db_user"`
	DBName   string `env:"db_name"`
	SSLMode  string `env:"db_sslmode"`
}

func NewController() *Controller {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("no env file found")
	}

	controller := &Controller{
		Cfg:   Config{},
		Dbcfg: DBConfig{},
	}

	err = env.Parse(&controller.Cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}
	err = env.Parse(&controller.Dbcfg)
	if err != nil {
		log.Fatalf("unable to parse database variables: %e", err)
	}

	config := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		controller.Dbcfg.Host,
		controller.Dbcfg.Port,
		controller.Dbcfg.User,
		controller.Dbcfg.Password,
		controller.Dbcfg.DBName,
		controller.Dbcfg.SSLMode,
	)
	Db, err := connectDB(config)
	if err != nil {
		panic("Failed to connect to database")
	}
	controller.Db = Db

	return controller
}

func (c *Controller) Run() error {

	docs.SwaggerInfo.Host = c.Cfg.Host
	docs.SwaggerInfo.Version = c.Cfg.Version
	docs.SwaggerInfo.BasePath = "/"

	router := http.NewServeMux()

	err := c.loadAll(router)

	if err != nil {
		return err
	}
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Headers,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening on port :8080")
	return server.ListenAndServe()
}

func (c *Controller) loadAll(router *http.ServeMux) error {

	// adminRouter := http.NewServeMux()

	// router.Handle("/", middleware.EnsureAdmin(adminRouter))

	router.HandleFunc("/", handlers.Make(defaultHandler))
	router.HandleFunc("/swagger/", httpSwagger.Handler())

	healthHandler := health.NewHandler()
	err := health.RegisterOpenAIHandler(router, healthHandler)
	if err != nil {
		return err
	}

	dadjokeHandler := dadjoke.NewHandler(
		c.Db,
	)
	err = dadjoke.RegisterDadJokeHandler(router, dadjokeHandler)
	if err != nil {
		return err
	}

	openAIHandler := openai.NewHandler(
		c.Cfg.LlamaURL,
		c.Cfg.LlamaPass,
	)
	err = openai.RegisterOpenAIHandler(router, openAIHandler)
	if err != nil {
		return err
	}

	return nil
}

func defaultHandler(
	w http.ResponseWriter,
	r *http.Request,
) (
	*handlers.StandardMessage,
	error,
) {
	resp := handlers.NewFailureResponse(
		"Page does not exist",
		http.StatusNotFound,
		[]string{"Page not found"},
	)
	return nil, resp
}

func connectDB(config string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
