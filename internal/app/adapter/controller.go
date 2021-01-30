package adapter

import (
	"net/http"

	"github.com/Dalot/goddd/internal/app/adapter/repository"
	"github.com/Dalot/goddd/internal/app/application/usecase"
	"github.com/Dalot/goddd/internal/app/domain/factory"
	"github.com/gin-gonic/gin"
)

var (
	projectRepository = repository.Project{}
	userRepository    = repository.User{}
)

// Controller is a controller
type Controller struct{}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	// NOTICE: following path is from CURRENT directory, so please run Gin from root directory
	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/", ctrl.index)
	r.GET("/projects", ctrl.projects)
	r.POST("/projects", ctrl.createProject)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Goilerplate",
	})
}

func (ctrl Controller) projects(c *gin.Context) {
	project := factory.Project{}
	proj := project.Generate()
	projectRepository.Save(proj)

	projects := usecase.Projects(projectRepository) // Dependency Injection
	c.JSON(200, projects)
}

func (ctrl Controller) createProject(c *gin.Context) {
	Name := c.Query("Name")
	args := usecase.CreateProjectArgs{
		Name:              Name,
		ProjectRepository: projectRepository,
		UserRepository:    userRepository,
	}
	project := usecase.CreateProject(args) // Dependency Injection
	c.JSON(200, project)
}
