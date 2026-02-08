package auth

import (
	"Lectures/ProdvinutiyGolang/Archetecture/config"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*config.Config
}

type AuthHandler struct {
	*config.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Register())
	router.HandleFunc("POST /auth/regisnter", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Auth")
	}
}
