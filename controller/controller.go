package controller

import (
	"net/http"

	"github.com/barbaraLuersen/GoGin/database"
	"github.com/barbaraLuersen/GoGin/models"
	"github.com/gin-gonic/gin"
)

// c *gin.Context = contexto do Gin que faz com que seja possível retornar o JSON
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz: ": "E ai " + nome + " tudo beleza?",
	})
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ListarTodos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func ListarPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			//Mensagem de erro
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func ListarPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	//Buca no banco de dados um aluno quando o CPF do aluno
	//for igual ao cpf que busquei e armazena em uma variável
	//aluno
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			//Mensagem de erro
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
