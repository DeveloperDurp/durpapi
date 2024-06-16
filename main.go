package main

import (
	"log"

	"gitlab.com/DeveloperDurp/DurpAPI/controller"
	"gitlab.com/DeveloperDurp/DurpAPI/docs"
)

//	@title			DurpAPI
//	@description	API for Durp's needs
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://durp.info
//	@contact.email	developerdurp@durp.info

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/
//	@securityDefinitions.apikey	Authorization
//	@in							header
//	@name						Authorization

func main() {
	c := controller.NewController()

	docs.SwaggerInfo.Host = c.Cfg.Host
	docs.SwaggerInfo.Version = c.Cfg.Version

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
	// adminRouter := http.NewServeMux()

	// router.Handle("/", middleware.EnsureAdmin(adminRouter))

}
