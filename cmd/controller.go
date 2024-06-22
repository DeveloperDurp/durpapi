package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.com/DeveloperDurp/DurpAPI/middleware"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/dadjoke"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/health"
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/openai"
	"gitlab.com/developerdurp/stdmodels"
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

	Db, err := connectDB(controller.Dbcfg)
	if err != nil {
		panic("Failed to connect to database")
	}
	controller.Db = Db

	return controller
}

func (c *Controller) Run() error {

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

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/swagger/*", httpSwagger.Handler())

	health, err := health.NewHandler()
	router.HandleFunc("GET /health/gethealth", health.Get)

	dadjoke, err := dadjoke.NewHandler(c.Db)
	router.HandleFunc("GET /jokes/dadjoke", dadjoke.Get)
	router.HandleFunc("POST /jokes/dadjoke", dadjoke.Post)
	router.HandleFunc("DELETE /jokes/dadjoke", dadjoke.Delete)

	openai, err := openai.NewHandler(c.Cfg.LlamaURL)
	router.HandleFunc("GET /openai/general", openai.GeneralOpenAI)
	router.HandleFunc("GET /openai/travelagent", openai.TravelAgentOpenAI)

	if err != nil {
		return err
	}
	return nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	stdmodels.FailureReponse("Page Does not exist", w, http.StatusNotFound, []string{"Page is not found"})
}

func connectDB(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
