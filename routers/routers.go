package routers

import (
	"gin-practice/controller"
	"gin-practice/pkg/controllerHelper"
	"github.com/gin-gonic/gin"
	"reflect"
)

func RouterInit() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/auth/login", BindRouter(new(controller.AuthController), "Login"))
		v1.POST("/auth/register", BindRouter(new(controller.AuthController), "Register"))

		v1.POST("/publish/message", BindRouter(new(controller.RabbitMQController), "Producer"))
		v1.GET("/consume/message", BindRouter(new(controller.RabbitMQController), "Consumer"))
		v1.POST("/publish/work-queue", BindRouter(new(controller.RabbitMQController), "WorkerQueueProducer"))
		v1.GET("/consume/work-queue", BindRouter(new(controller.RabbitMQController), "WorkerQueueConsumer"))
	}

	return router
}

func BindRouter(ct controllerHelper.ControllerInterface, methodName string) gin.HandlerFunc {
	reflectValue := reflect.ValueOf(ct)
	execController, ok := reflectValue.Interface().(controllerHelper.ControllerInterface)
	if !ok {
		panic("controller is not impolent controller interface")
	}
	method := reflectValue.MethodByName(methodName)
	return func(ctx *gin.Context) {
		execController.Init(execController, ctx)
		execController.Prepare()
		method.Call(nil)
		execController.Finish()
	}
}
