package factory

import "gorm.io/gorm"

func Seed() {
	factoryProject := Project{}
	for i := 0; i < 20; i++ {
		proj := factoryProject.Generate()
		projectRepository.Save(proj)
	}

	projects := projectRepository.Index()
	for _, project := range projects {
		factoryTask := Task{}
		for i := 0; i < 20; i++ {
			task := factoryTask.Generate(project.ID)
			taskRepository.Save(task)
		}
	}
}

func Wipe(conn *gorm.DB) {
	conn.Exec("DROP TABLE IF EXISTS projects").Exec("DROP TABLE IF EXISTS users").Exec("DROP TABLE IF EXISTS tasks")

}
