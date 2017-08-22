package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Register",
		"GET",
		"/register",
		Register,
	},
	Route{
		"GoogleLogin",
		"GET",
		"/googlelogin",
		GoogleLogin,
	},
	Route{
		"GoogleCallback",
		"GET",
		"/googlecallback",
		GoogleCallback,
	},
	Route{
		"GameIndex",
		"GET",
		"/games",
		GameIndex,
	},
	Route{
		"GameShow",
		"GET",
		"/games/{gameId}",
		GameShow,
	},
}
