package main 
import (
	// "fmt"
	"github.com/bisratlike/toGo/pkg/db" 
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/bisratlike/toGo/pkg/config"
	"github.com/bisratlike/toGo/internal/auth_service/router"
)

func main() {
     cfg := config.LoadConfig()
    db.Connect(cfg)
    db.RunMigrations()
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    router.AuthRoutes(r, db.DB)
    log.Printf("auth service started on port 8080")
    http.ListenAndServe(":8080", r)
}