package server

import (
	"net/http"

	"github.ducminhsw.prepare-project/internal/handlers"
	"github.ducminhsw.prepare-project/internal/middlewares"
)

func (server *Server) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// auth route (using email for logging)
	authHandler := handlers.NewAuthHandler(server.Conf.Database)
	mux.HandleFunc("POST /auth/register", middlewares.HeaderAddMiddleware(authHandler.HandlerRegister()))
	mux.HandleFunc("POST /auth/login", middlewares.HeaderAddMiddleware(authHandler.HandlerLogin()))

	// user route
	userHandler := handlers.NewUserHandler(server.Conf.Database)
	mux.HandleFunc("POST /user/pair/{target}/{heart-key}", middlewares.HeaderAddMiddleware(userHandler.HandlePairLover()))
	mux.HandleFunc("PUT /user/setting/setup", middlewares.HeaderAddMiddleware(userHandler.HandleSettingUser()))

	// conversation route
	mux.HandleFunc("POST /conversation/{page}/{numberOfMessages}", middlewares.HeaderAddMiddleware(userHandler.HandleGetConversation()))

	// memories route
	mux.HandleFunc("GET /memo/all", middlewares.HeaderAddMiddleware(userHandler.HandleGetAllMemo()))
	mux.HandleFunc("POST /memo/more", middlewares.HeaderAddMiddleware(userHandler.HandleGetPartOfMemo()))
	mux.HandleFunc("POST /memo/{id}/detail", middlewares.HeaderAddMiddleware(userHandler.HandleGetDetailMemo()))

	return mux
}
