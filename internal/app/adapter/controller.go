package adapter

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dalot/goddd/internal/app/adapter/repository"
	"github.com/Dalot/goddd/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	projectRepository = repository.Project{}
	userRepository    = repository.User{}
	taskRepository    = repository.Task{}
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
	r.GET("/projects/:id", ctrl.getProjectByID)
	r.POST("/projects", ctrl.createProject)
	r.PUT("/projects/:id", ctrl.updateProject)
	r.DELETE("/projects/:id", ctrl.deleteProject)

	r.GET("/tasks", ctrl.tasks)
	r.GET("/tasks/:id", ctrl.getTaskByID)
	r.POST("/tasks/:id/actions", ctrl.taskActions)
	r.POST("/tasks", ctrl.createTask)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Code Challenge",
	})
}

func (ctrl Controller) projects(c *gin.Context) {
	//TODO: Needs authentication
	userIDInput := c.Query("user_id")
	var userID uint
	if len(userIDInput) > 0 {
		val, err := strconv.ParseUint(userIDInput, 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Could not parse user_id",
			})
			return
		}

		userID = uint(val)
	} else {
		userID = 0
	}

	projects := usecase.Projects(projectRepository, userID) // Dependency Injection
	c.JSON(200, projects)
}

func (ctrl Controller) getProjectByID(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	proj, err := usecase.GetProjectByID(projectRepository, uint(val))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Project not found",
			})
			return
		}
	}

	c.JSON(200, proj)
}
// Binding from JSON
type ProjectInput struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}
func (ctrl Controller) createProject(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	args := usecase.CreateProjectArgs{
		Name:              input.Name,
		UserID: 			   1,// TODO: insert here authenticated user 
		ProjectRepository: projectRepository,
		UserRepository:    userRepository,
	}
	project, err := usecase.CreateProject(args) // Dependency Injection
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Authenticated user not found", // TODO: When I do have authetnication this will not make sense anymore
			})
			return
		} else {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Something happened",
			})
			return
		}
	} 
	
	c.JSON(200, project)
}

func (ctrl Controller) updateProject(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	args := usecase.UpdateProjectArgs{
		ProjectID: uint(val),
		Name: input.Name,
		ProjectRepository: projectRepository,
	}
	project, err := usecase.UpdateProject(args) // Dependency Injection
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Project not found",
			})
			return
		} else {
			c.AbortWithStatusJSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(200, project)
}

func (ctrl Controller) deleteProject(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	usecase.DeleteProject(projectRepository, uint(val)) // Dependency Injection
	c.JSON(204, gin.H{})
}

func (ctrl Controller) tasks(c *gin.Context) {
	//TODO: Needs authentication
	projectIDInput := c.Query("project_id")
	var projectID uint
	if len(projectIDInput) > 0 {
		val, err := strconv.ParseUint(projectIDInput, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Could not parse project_id",
			})
		}

		projectID = uint(val)
	} else {
		projectID = 0
	}

	tasks := usecase.Tasks(taskRepository, projectID) // Dependency Injection
	c.JSON(200, tasks)
}

func (ctrl Controller) getTaskByID(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	task, err := usecase.GetTaskByID(taskRepository, uint(val))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Task not found",
			})
			return
		}
	}

	c.JSON(200, task)
}

// Binding from JSON
type taskAction struct {
	Type string `form:"type" json:"type" xml:"type"  binding:"required"`
}

func (ctrl Controller) taskActions(c *gin.Context) {
	var action taskAction
	if err := c.ShouldBindJSON(&action); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	if action.Type == "finish" {
		task, err := usecase.FinishTask(taskRepository, uint(val))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(404, gin.H{
					"message": "Task not found",
				})
				return
			} else if errors.Is(err, usecase.TaskErrAlreadyFinished) {
				c.JSON(200, gin.H{
					"message": err.Error(),
				})
				return
			} else {
				c.AbortWithStatusJSON(500, gin.H{
					"message": "Something happened",
					"err":     err,
				})
				return
			}
		}

		c.JSON(200, task)
	} else {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "There is no such action",
		})
	}

}

func (ctrl Controller) createTask(c *gin.Context) {
	Name := c.Query("Name")
	args := usecase.CreateProjectArgs{
		Name:              Name,
		ProjectRepository: projectRepository,
		UserRepository:    userRepository,
	}
	project, err := usecase.CreateProject(args) // Dependency Injection

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "Project not found",
		})
		return
	}

	c.JSON(200, project)
}
