package adapter

import (
	"encoding/json"
	"errors"
	"log"
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

	r.Use(CORS())

	authorized := r.Group("/api")
	authorized.Use(JWTAuth())

	authorized.GET("", ctrl.index)                         // fetch all or by user id
	authorized.GET("/projects", ctrl.projects)             // fetch all or by user id
	authorized.GET("/projects/:id", ctrl.getProjectByID)   // fetch project
	authorized.POST("/projects", ctrl.createProject)       // create project
	authorized.PUT("/projects/:id", ctrl.updateProject)    // update project
	authorized.DELETE("/projects/:id", ctrl.deleteProject) // delete project

	authorized.GET("/tasks", ctrl.tasks)                    // fetch all or by project id
	authorized.GET("/tasks/:id", ctrl.getTaskByID)          // fetch task
	authorized.POST("/tasks/:id/actions", ctrl.taskActions) // Execute action to task (e.g finish)
	authorized.POST("/tasks", ctrl.createTask)              // create task
	authorized.PUT("/tasks/:id", ctrl.updateTask)           // update task
	authorized.DELETE("/tasks/:id", ctrl.deleteTask)        // delete task

	r.POST("/login", ctrl.login)
	r.POST("/signup", ctrl.signup)

	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// LoginInput Binding from JSON TODO: valide the length of the password and maybe other things?
type LoginInput struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func (ctrl Controller) login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	args := usecase.LoginArgs{
		Email:          input.Email,
		Password:       input.Password,
		UserRepository: userRepository,
	}

	cookie, err := usecase.Login(args)
	if err != nil {
		if errors.Is(err, usecase.UserErrCouldNotCreateJWT) {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Something happened",
				"status":  "error",
			})
			return
		} else if errors.Is(err, usecase.UserErrWrongPassword) {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "The password or the email does not match any account.", // Same error for obscurity purposes
				"status":  "error",
			})
			return
		} else {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"message": "The password or the email does not match any account.", // Same error for obscurity purposes
					"status":  "error",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something happened.",
				"status":  "error",
			})
			log.Fatal(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   cookie.Value,
		"status":  "ok",
	})
}

// RegisterInput from JSON TODO: valide the length of the password and maybe other things?
type RegisterInput struct {
	Email           string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password        string `form:"password" json:"password" xml:"password"  binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" xml:"confirm_password"  binding:"required"`
	Username        string `form:"username" json:"username" xml:"username"  binding:"required"`
}

func (ctrl Controller) signup(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.AbortWithStatusJSON(200, gin.H{
			"message": "Passwords provided do not match",
			"status":  "error",
		})
		return
	}

	args := usecase.RegisterArgs{
		Email:          input.Email,
		Password:       input.Password,
		Username:       input.Username,
		UserRepository: userRepository,
	}

	user, err := usecase.Register(args)
	if err != nil {
		if errors.Is(err, usecase.UserErrAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  "error",
			})
			return
		}
		panic(err)
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "authenticated",
		"status":  "ok",
		"user":    string(jsonUser),
	})
}

func (ctrl Controller) projects(c *gin.Context) {

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	projects := usecase.Projects(projectRepository, user.ID) // Dependency Injection
	c.JSON(200, gin.H{
		"data":   projects,
		"status": "ok",
	})
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

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	proj, err := usecase.GetProjectByID(projectRepository, uint(val))

	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this project.",
		})
		return
	}

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

// ProjectInput from JSON
type ProjectInput struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}

func (ctrl Controller) createProject(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := GetAuthUser(c)

	args := usecase.CreateProjectArgs{
		Name:              input.Name,
		UserID:            user.ID,
		ProjectRepository: projectRepository,
		UserRepository:    userRepository,
	}
	project, err := usecase.CreateProject(args) // Dependency Injection
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something happened",
			"status": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": project,
		"status": "ok",
	})
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

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	proj, err := usecase.GetProjectByID(projectRepository, uint(val))

	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this project.",
		})
		return
	}

	args := usecase.UpdateProjectArgs{
		ProjectID:         uint(val),
		Name:              input.Name,
		ProjectRepository: projectRepository,
	}
	project, err := usecase.UpdateProject(args) // Dependency Injection
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Project not found",
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": project,
		"status": "ok",
	})
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

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	proj, err := usecase.GetProjectByID(projectRepository, uint(val))

	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this project.",
		})
		return
	}

	err = usecase.DeleteProject(projectRepository, uint(val))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Project not found",
			})
			return
		}

		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return

	}
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

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	if projectID != 0 {
		proj, _ := usecase.GetProjectByID(projectRepository, projectID)

		if user.ID != proj.UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You have no access to this project.",
			})
			return
		}
		tasks := usecase.Tasks(taskRepository, projectID) // Dependency Injection
		c.JSON(200, tasks)
		return
	}
	
	projects := usecase.Projects(projectRepository, user.ID)
	projectIDs := []uint{}
	for _, proj := range projects {
		projectIDs = append(projectIDs, proj.ID)
	}
	tasks := usecase.TasksByProjectIDs(taskRepository, projectIDs)
	c.JSON(200, gin.H{
		"data":   tasks,
		"status": "ok",
	})
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

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	proj, err := usecase.GetProjectByID(projectRepository, task.ProjectID)
	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this task.",
		})
		return
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
	task, err := usecase.GetTaskByID(taskRepository, uint(val))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Task not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	proj, err := usecase.GetProjectByID(projectRepository, task.ProjectID)
	user, err := GetAuthUser(c)
	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this task.",
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

		c.JSON(200, gin.H{
			"data": task,
			"status": "ok",
		})
	} else {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "There is no such action",
		})
	}

}

// CreateTaskInput from JSON
type CreateTaskInput struct {
	Name        string `form:"name" json:"name" xml:"name"  binding:"required"`
	ProjectID   uint   `form:"project_id" json:"project_id" xml:"project_id"  binding:"required"`
	Description string `form:"description" json:"description" xml:"description" `
}

func (ctrl Controller) createTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	proj, err := usecase.GetProjectByID(projectRepository, input.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Project not found",
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := GetAuthUser(c)
	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this project.",
		})
		return
	}

	args := usecase.CreateTaskArgs{
		Name:              input.Name,
		ProjectID:         input.ProjectID,
		Description:       input.Description,
		ProjectRepository: projectRepository,
		TaskRepository:    taskRepository,
	}
	task, err := usecase.CreateTask(args) // Dependency Injection

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Project not found",
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data":   task,
		"status": "ok",
	})
}

// UpdateTaskInput Binding from JSON
type UpdateTaskInput struct {
	Name        string `form:"name" json:"name" xml:"name"  binding:"required"`
	Description string `form:"description" json:"description" xml:"description"`
}

func (ctrl Controller) updateTask(c *gin.Context) {
	var input UpdateTaskInput
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

	task, err := usecase.GetTaskByID(taskRepository, uint(val))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Task not found",
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	proj, err := usecase.GetProjectByID(projectRepository, task.ProjectID)
	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this task.",
		})
		return
	}

	args := usecase.UpdateTaskArgs{
		TaskID:         task.ID,
		Name:           input.Name,
		Description:    input.Description,
		TaskRepository: taskRepository,
	}
	task, err = usecase.UpdateTask(args) // Dependency Injection
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Task not found",
			})
			return
		}
		if errors.Is(err, usecase.TaskErrCannotBeUpdated) {
			c.AbortWithStatusJSON(200, gin.H{
				"message": err.Error(),
				"status": "error",
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": task,
		"status": "ok",
	})
}

func (ctrl Controller) deleteTask(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	user, err := GetAuthUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "error",
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
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	proj, err := usecase.GetProjectByID(projectRepository, task.ProjectID)
	if user.ID != proj.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You have no access to this task.",
		})
		return
	}

	err = usecase.DeleteTask(taskRepository, task.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "Task not found",
			})
			return
		}
		if errors.Is(err, usecase.TaskErrCannotBeDeleted) {
			c.AbortWithStatusJSON(200, gin.H{
				"message": err.Error(),
				"status": "error",
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(204, gin.H{})
}
