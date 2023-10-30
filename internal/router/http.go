package router

import (
	"github.com/gin-gonic/gin"
	"todo_list/internal/controllers"
	"todo_list/internal/middleware"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.Use(middleware.Cors())
	app.Use(gin.Recovery())

	test := app.Group("/test")
	test.Use(middleware.TestMiddleware())
	test.GET("/1", func(context *gin.Context) {
		context.JSON(200, "test 1 success")
	})
	user := app.Group("fUser")
	{
		u := controllers.UserCtl{}
		user.POST("/oneDaySuggestion", u.OneDaySuggestion)
		user.GET("/getSuggestion", u.GetSuggestion)
		user.POST("/login", u.Login)
		user.POST("/updateUserName", u.UpdateUserName)
		user.POST("/uploadAvatar", u.UploadAvatar)
		user.POST("/identifyCodeSend", u.IdentifyCodeSend)
		user.POST("/register", u.Register)
		user.POST("/resetPassword", u.ResetPassword)
		user.POST("/forgetPassword", u.ForgetPassword)
	}

	task := app.Group("/oneDayTask")
	{
		t := controllers.TaskCtl{}
		task.POST("/createTask", t.Create)
		task.POST("/updateTask", t.Update)
		task.POST("/getTasks", t.GetAllByUserId)
		task.POST("/deleteTask", t.Delete)
	}

	newUser := app.Group("user")
	{
		u := controllers.New()
		newUser.GET("/query", u.Query)
		newUser.POST("/login", u.NewLogin)
		newUser.POST("/register", u.Register)
		newUser.POST("/resetPassword", u.ResetPassword)
		newUser.POST("/forgetPassword", u.ForgetPassword)

	}

	newTask := app.Group("task")
	{
		t := controllers.TaskCtl{}
		newTask.POST("/add", t.AddNewTask)
		newTask.POST("/update", t.UpdateNew)
		newTask.GET("/list", t.QueryTasksNew)
		newTask.GET("/count", t.QueryTasksCountNew)

	}

	ng := app.Group("taskgroup")
	{
		t := controllers.NewTaskGroupController()
		ng.POST("/add", t.Add)
		ng.POST("/update", t.Update)
		ng.GET("/query", t.QueryByUserId)
	}

	return app
}
