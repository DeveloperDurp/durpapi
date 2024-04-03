package main

import (
	"fmt"
	"net/http"

	"github.com/swaggo/http-swagger"

	"gitlab.com/DeveloperDurp/DurpAPI/controller"
	"gitlab.com/DeveloperDurp/DurpAPI/docs"
	"gitlab.com/DeveloperDurp/DurpAPI/middleware"
)

//	@title			DurpAPI
//	@description	API for Durp's needs
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://durp.info
//	@contact.email	developerdurp@durp.info

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/api
//	@securityDefinitions.apikey	Authorization
//	@in							header
//	@name						Authorization

func main() {
	c := controller.NewController()

	docs.SwaggerInfo.Host = c.Cfg.Host
	docs.SwaggerInfo.Version = c.Cfg.Version

	router := http.NewServeMux()
	router.HandleFunc("/swagger/*", httpSwagger.Handler())
	router.HandleFunc("GET /api/health/gethealth", c.GetHealth)
	router.HandleFunc("GET /api/jokes/dadjoke", c.GetDadJoke)
	router.HandleFunc("POST /api/jokes/dadjoke", c.PostDadJoke)
	router.HandleFunc("DELETE /api/jokes/dadjoke", c.DeleteDadJoke)
	router.HandleFunc("GET /api/openai/general", c.GeneralOpenAI)
	router.HandleFunc("GET /api/openai/travelagent", c.TravelAgentOpenAI)
	// adminRouter := http.NewServeMux()

	// router.Handle("/", middleware.EnsureAdmin(adminRouter))

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
