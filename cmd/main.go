package main

import (
	"encoding/json"
	"net/http"

	"github.com/frasnym/go-boilerplate/app/controller"
	gorm "github.com/frasnym/go-boilerplate/app/infrastructure/gorm/database"
	"github.com/frasnym/go-boilerplate/app/infrastructure/gorm/repository"
	router "github.com/frasnym/go-boilerplate/app/infrastructure/http"
	"github.com/frasnym/go-boilerplate/app/infrastructure/logging"
	"github.com/frasnym/go-boilerplate/app/usecase/fanuser"
)

func main() {
	const port string = ":8000"

	// Initialize database
	db, err := gorm.NewConnectionDB("sqlite", "boilerplatedb", "localhost", "user", "password", 5432)
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
		log        logging.Logging = logging.NewZapLogging() // Use zap for logging
		httpRouter router.Router   = router.NewMuxRouter(log.LoggingMiddleware)
	)

	// Assign route
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("FrasNym Go Boilerplate version 1.0.0")
	})
	httpRouter.POST("/fanuser", fanuserController.SignUpFanuser)

	// Start service
	httpRouter.SERVE(port)
}
