package routes

import (
	"github.com/barbaraLuersen/GoGin/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos", controller.CriaAluno)
	r.GET("/alunos", controller.ListarTodos)
	r.GET("/alunos/:id", controller.ListarPorId)
	r.DELETE("/alunos/:id", controller.DeletarAluno)
	r.PATCH("/alunos/:id", controller.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controller.ListarPorCPF)

	r.Run() //listening and servering HTTP on :8080.
	//É possível mudar isso setando outra porta entre os parenteses

}
