package main

import "net/http"

// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	Host        string
// 	HandlerFunc http.HandlerFunc
// }
//
// type Routes []Route
//
// var routes = Routes{
// 	Route{
// 		"IndieExpo",
// 		"POST",
// 		"/send",
// 		"www.indieexpo.co",
// 		SendEmail,
// 	},
// }
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"IndieExpo",
		"POST",
		"/send",
		SendEmail,
	},
	Route{
		"HireMe",
		"POST",
		"/demo",
		SmartBiz,
	},
}
