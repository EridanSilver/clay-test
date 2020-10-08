package main

import (
	"github.com/EridanSilver/clay-test/internal/user_service/service"
	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/utrack/clay/v2/log"
	"github.com/utrack/clay/v2/transport/middlewares/mwgrpc"
	"github.com/utrack/clay/v2/transport/server"

	_ "github.com/utrack/clay/doc/example/static/statik"
)

func main() {
	staticFS, err := fs.New()
	if err != nil {
		logrus.Fatal(err)
	}
	hmux := chi.NewRouter()
	hmux.Mount("/", http.FileServer(staticFS))

	impl := service.NewService()
	srv := server.NewServer(
		8080,
		// Pass our mux with Swagger UI
		server.WithHTTPMux(hmux),
		// Recover from both HTTP and gRPC panics and use our own middleware
		server.WithGRPCUnaryMiddlewares(mwgrpc.UnaryPanicHandler(log.Default)),
	)
	err = srv.Run(impl)
	if err != nil {
		logrus.Fatal(err)
	}
}
