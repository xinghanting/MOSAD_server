package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route 根据url进行路由的数据结构体
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes 一个路由链接的序列
type Routes []Route

// Service 方便和数据库端
var Service *NewService

// NewRouter 创建新的Router
func NewRouter(s *NewService) *mux.Router {
	Service = s
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index 首页请求，用于测试
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	// Route{
	// 	"AddOnetime",
	// 	strings.ToUpper("Post"),
	// 	"/onetime",
	// 	AddOnetime,
	// },

	// Route{
	// 	"DeleteOnetime",
	// 	strings.ToUpper("Delete"),
	// 	"/onetime/{onetimeID}",
	// 	DeleteOnetime,
	// },

	// Route{
	// 	"GetBlogByBlogID",
	// 	strings.ToUpper("Get"),
	// 	"/blog/{BlogID}",
	// 	GetBlogByBlogID,
	// },

	// Route{
	// 	"GetBlogs",
	// 	strings.ToUpper("Get"),
	// 	"/blogs",
	// 	GetBlogs,
	// },

	// Route{
	// 	"GetBlogsByID",
	// 	strings.ToUpper("Get"),
	// 	"/blogs/{ID}",
	// 	GetBlogsByID,
	// },

	// Route{
	// 	"GetCommentByBlogID",
	// 	strings.ToUpper("Get"),
	// 	"/comment/{BlogID}",
	// 	GetCommentByBlogId,
	// },

	// Route{
	// 	"PostComment",
	// 	strings.ToUpper("Post"),
	// 	"/comment/{BlogID}",
	// 	PostComment,
	// },

	// Route{
	// 	"GetTagByBlogID",
	// 	strings.ToUpper("Get"),
	// 	"/tag/{BlogID}",
	// 	GetTagByBlogId,
	// },

	// Route{
	// 	"PostTag",
	// 	strings.ToUpper("Post"),
	// 	"/tag/{BlogID}",
	// 	PostTag,
	// },

	// Route{
	// 	"CreateUser",
	// 	strings.ToUpper("Post"),
	// 	"/user/signup",
	// 	CreateUser,
	// },

	// Route{
	// 	"GetUserByID",
	// 	strings.ToUpper("Get"),
	// 	"/user/{ID}/info",
	// 	GetUserByID,
	// },

	// Route{
	// 	"LoginUser",
	// 	strings.ToUpper("Post"),
	// 	"/user/login",
	// 	LoginUser,
	// },

	// Route{
	// 	"LogoutUser",
	// 	strings.ToUpper("Get"),
	// 	"/user/logout",
	// 	LogoutUser,
	// },
}
