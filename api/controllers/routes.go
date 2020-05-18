package controllers

import "github.com/goapi/api/middleware"

func (server *Server) initializeRoutes(){
	//HomeRoute
	server.Router.HandleFunc("/", middleware.SetMiddleWareJSON(server.Home)).Methods("GET")
	//Login route
	server.Router.HandleFunc("/", middleware.SetMiddleWareJSON(server.Login)).Methods("POST")
	//Users routes
	server.Router.HandleFunc("/users", middleware.SetMiddleWareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middleware.SetMiddleWareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middleware.SetMiddleWareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middleware.SetMiddleWareJSON(middleware.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareAuthentication(server.DeleteUser)).Methods("DELETE")

	//Posts routes
	server.Router.HandleFunc("/posts", middleware.SetMiddleWareJSON(server.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/posts", middleware.SetMiddleWareJSON(server.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middleware.SetMiddleWareJSON(server.GetPost)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middleware.SetMiddleWareJSON(middleware.SetMiddlewareAuthentication(server.UpdatePost))).Methods("PUT")
	server.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareAuthentication(server.DeletePost)).Methods("DELETE")
}
