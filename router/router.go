package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"go_cnode/controllers/reply"
	"go_cnode/controllers/sign"
	"go_cnode/controllers/site"
	"go_cnode/controllers/topic"
	"html/template"
	"net/http"
	// _ "net/http/pprof"
	// "log"
)

func add(left int, right int) int {
	return left + right
}
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	store := sessions.NewCookieStore([]byte("secret123"))
	router.Use(sessions.Middleware("my_session", store))
	//router.Delims("([{", "}])")//模板函数隔离符
	router.SetFuncMap(template.FuncMap{"add": add})
	router.LoadHTMLGlob("views/**/*")
	router.StaticFS("/public", http.Dir("./public"))
	router.StaticFile("/favicon.ico", "./public/images/cnode_icon_32.png")
	router.GET("/", site.Index)
	router.GET("/about", site.About)
	router.GET("/signup", sign.ShowSignup)
	router.POST("/signup", sign.Signup)
	router.POST("/signout", sign.Signout)
	router.GET("/signin", sign.Signin)
	router.GET("/passport/github", sign.GithubSignup)
	router.GET("/github/callback", sign.GithubCallBack)
	router.GET("/setting", sign.Setting)
	router.GET("/my/messages", sign.Message)

	router.POST("/passport/local", sign.Login)
	router.GET("/search_pass", sign.SearchPass)
	router.GET("/api", site.Api)
	router.GET("/getstart", site.Getstart)
	router.GET("/topic/:id", topic.Index)
	router.GET("/topic/:id/top", topic.Top)
	router.GET("/topics/create", topic.ShowCreate)
	router.POST("/topic/create", topic.Create)
	router.GET("/active_account", sign.ActiveAccount) // 帐号激活
	router.POST("/reply/:topic_id", reply.Add)
	router.POST("/edit/reply/:reply_id", reply.Edit)
	router.GET("/edit/reply/:reply_id", reply.ShowEdit)
	router.POST("/upload", topic.Upload) // 上传图片
	// go func() {
	//     log.Println(http.ListenAndServe("localhost:10000", nil))
	// }()

	return router
}
