package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/frasnym/go-boilerplate/app/controller"
	gorm "github.com/frasnym/go-boilerplate/app/infrastructure/gorm/database"
	"github.com/frasnym/go-boilerplate/app/infrastructure/gorm/repository"
	router "github.com/frasnym/go-boilerplate/app/infrastructure/http"
	"github.com/frasnym/go-boilerplate/app/infrastructure/logging"
	"github.com/frasnym/go-boilerplate/app/usecase/fanuser"
	"github.com/spf13/viper"
)

func main() {
	// Initialize configuration
	var configProfile = "local"
	if os.Getenv("BP_ENV") != "" {
		configProfile = os.Getenv("BP_ENV")
	}

	var configFileName []string
	configFileName = append(configFileName, "config-")
	configFileName = append(configFileName, configProfile)

	viper.SetConfigType("yaml")
	viper.SetConfigName(strings.Join(configFileName, ""))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Initialize database
	db, err := gorm.NewConnectionDB(viper.GetString("database.driver"), viper.GetString("database.name"), viper.GetString("database.host"), viper.GetString("database.username"), viper.GetString("database.password"), viper.GetInt("database.port"))
	if err != nil {
		panic(err)
	}

	// Fanuser
	var (
		fanuserRepository fanuser.FanuserRepository    = repository.NewFanuserRepository(repository.NewBaseRepository(db))
		fanuserService    fanuser.FanuserUseCase       = fanuser.NewFanuserService(fanuserRepository)
		fanuserController controller.FanuserController = controller.NewFanuserController(fanuserService)
	)

	// Initialize route
	var (
		// log        logging.Logging = logging.NewDefaultLogging() // Use default logging
		log logging.Logging = logging.NewZapLogging() // Use zap for logging
		// httpRouter router.Router   = router.NewMuxRouter(log.LoggingMiddleware)
		httpRouter router.Router = router.NewGinRouter(log.LoggingMiddleware)
	)

	// Assign route
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("FrasNym Go Boilerplate version 1.0.0")
	})
	httpRouter.POST("/fanuser", fanuserController.SignUpFanuser)

	// Start service
	httpRouter.SERVE(viper.GetString("server.port"))
}
