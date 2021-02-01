<template>
  <div>
    <div class="ui center aligned basic segment">
      <div class="ui left icon action input">
        <i class="plus icon"></i>
        <input
          type="text"
          name="new_project_name"
          v-model="new_project_name"
          placeholder="Project"
          @keydown.enter="createNewProject"
        />
        <div
          class="ui positive submit button"
          v-on:click="createNewProject"
        >
          Create New Project
        </div>
      </div>
    </div>
    <div
      class="ui three column doubling stackable grid container"
      v-for="(projectsGroup, groupIndex) in projects"
      v-bind:key="groupIndex"
    >
      <div
        class="column"
        v-for="(project, projectIndex) in projectsGroup"
        v-bind:key="project.ID"
      >
        <div class="ui centered card fluid">
          <div class="content">
            <div class="header">
              <span
                v-show="!project.edit"
                :id="'project_name_span_' + project.id"
                v-on:click.prevent="
                  toggleProjectEdit(
                    $event,
                    project.id,
                    projectIndex,
                    groupIndex
                  )
                "
                >{{ project.name }}</span
              >
              <input
                type="text"
                v-model="project.name"
                v-show="project.edit"
                :id="'project_name_' + project.id"
                @keydown.enter="
                  updateProject($event, project.id, projectIndex, groupIndex)
                "
              />
              <a
                class="trash icon right floated"
                v-on:click.prevent="
                  deleteProject(project.id, projectIndex, groupIndex)
                "
              >
                <i class="trash red icon"></i>
              </a>
            </div>
            <div class="description">
              <div class="ui centered card">
                <div
                  class="content"
                  v-for="(task, taskIndex) in project.tasks"
                  v-bind:key="task.id"
                >
                  <div
                    v-show="task.finishedAt.length"
                    class="ui horizontal label small"
                  >
                    {{ task.finishedAt }}
                  </div>
                  <div
                    class="ui horizontal label small"
                    v-bind:class="{
                      blue: isNew(taskIndex, projectIndex, groupIndex),
                      green: isFinished(taskIndex, projectIndex, groupIndex),
                    }"
                  >
                    {{ task.status }}
                  </div>
                  <div class="description">
                    <span
                      style="font-size: 1em"
                      v-show="!task.edit"
                      :id="'task_name_span_' + task.id"
                      v-on:click.prevent="
                        toggleTaskEdit(
                          task.id,
                          taskIndex,
                          projectIndex,
                          groupIndex
                        )
                      "
                      >{{ task.name }}</span
                    >
                    <input
                      type="text"
                      v-model="task.name"
                      v-show="task.edit"
                      :id="'task_name_' + task.id"
                      @keydown.enter="
                        updateTask(task.id, taskIndex, projectIndex, groupIndex)
                      "
                    />
                    <a
                      class="trash icon right floated"
                      v-on:click.prevent="
                        deleteTask(task.id, taskIndex, projectIndex, groupIndex)
                      "
                      ><i class="trash red icon"></i
                    ></a>

                    <a
                      class="check circle right floated"
                      v-show="
                        projects[groupIndex][projectIndex].tasks[taskIndex]
                          .status === 'new'
                      "
                      v-on:click.prevent="
                        finishTask(task.id, taskIndex, projectIndex, groupIndex)
                      "
                    >
                      <i class="check circle green icon"></i>
                    </a>
                  </div>
                  <div class="meta">
                    {{ task.description }}
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="ui action input">
            <input
              type="text"
              name="new_task_name"
              v-model="form.task_name[project.id]"
              placeholder="My new task"
              value=""
            />

            <div
              class="ui positive right icon button"
              v-on:click="createNewTask(project.id, projectIndex, groupIndex)"
            >
              <i class="plus icon"></i>
              Add Task
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="ui basic segment">
      <div class="ui submit button" v-on:click="logout">Logout</div>
    </div>
  </div>
</template>

<script type = "text/javascript" >
import axios from 'axios';

export default {
  data() {
    return {
      projects: [],
      form: {
        task_name: [],
      },
      new_project_name: '',
      new_task_name: '',
      isNew: (taskIndex, projectIndex, groupIndex) =>
        this.projects[groupIndex][projectIndex].tasks[taskIndex].status ===
        'new',
      isFinished: (taskIndex, projectIndex, groupIndex) =>
        this.projects[groupIndex][projectIndex].tasks[taskIndex].status ===
        'finished',
    };
  },
  methods: {
    logout() {
      localStorage.removeItem('token');
      this.$router.push('login');
    },
    finishTask(taskId, taskIndex, projectIndex, groupIndex) {
      axios
        .post(`http://localhost:8080/api/tasks/${taskId}/actions`, {
          type: 'finish',
        })
        .then((response) => {
          const data = response.data;
          if (data.status === 'ok') {
            this.projects[groupIndex][projectIndex].tasks[taskIndex].status =
              data.data.Status;
          }
        });
    },
    deleteTask(taskId, taskIndex, projectIndex, groupIndex) {
      axios
        .delete(`http://localhost:8080/api/tasks/${taskId}`)
        .then((response) => {
          // TODO: Something happened flash message?
          if (response.status === 204) {
            this.projects[groupIndex][projectIndex].tasks.splice(taskIndex, 1);
          } else if (response.status === 200) {
            // TODO: remove the alert
            // eslint-disable-next-line no-alert
            alert(response.data.message);
          }
        });
    },
    deleteProject(projectId, projectIndex, groupIndex) {
      axios
        .delete(`http://localhost:8080/api/projects/${projectId}`)
        .then((response) => {
          // TODO: Something happened flash message?
          if (response.status === 204) {
            this.projects[groupIndex].splice(projectIndex, 1);
          }
        });
    },
    toggleTaskEdit(taskId, taskIndex, projectIndex, groupIndex) {
      const input = document.getElementById(`task_name_${taskId}`);
      const span = document.getElementById(`task_name_span_${taskId}`);
      input.style.display = input.style.display === '' ? 'none' : '';
      span.style.display = span.style.display === 'none' ? '' : 'none';
      // Focus input field
      if (this.projects[groupIndex][projectIndex].tasks[taskIndex].edit) {
        this.$nextTick(() => {
          input.focus();
        });
      }
    },
    updateTask(taskId, taskIndex, projectIndex, groupIndex) {
      // save your changes
      const taskUpdated = new Promise((resolve, reject) => {
        const task = this.projects[groupIndex][projectIndex].tasks[taskIndex];
        axios
          .put(`http://localhost:8080/api/tasks/${taskId}`, {
            name: task.name,
          })
          .then((response) => {
            const data = response.data;
            if (data.status === 'error') {
              // TODO: Something happened flash message?
              // TODO: Put the old value
              // TODO: remove the alert
              // eslint-disable-next-line no-alert
              alert(data.message);
              resolve();
            } else if (data.status === 'ok') {
              resolve();
            } else {
              reject();
            }
          });
      });
      taskUpdated.then(() => {
        this.toggleTaskEdit(taskId, taskIndex, projectIndex, groupIndex);
      });
    },
    toggleProjectEdit(ev, projectId, projectIndex, groupIndex) {
      const input = document.getElementById(`project_name_${projectId}`);
      const span = document.getElementById(`project_name_span_${projectId}`);
      input.style.display = input.style.display === '' ? 'none' : '';
      span.style.display = span.style.display === 'none' ? '' : 'none';
      // Focus input field
      if (this.projects[groupIndex][projectIndex].edit) {
        this.$nextTick(() => {
          input.focus();
        });
      }
    },
    updateProject(ev, projectId, projectIndex, groupIndex) {
      // save your changes
      const projectUpdated = new Promise((resolve, reject) => {
        const project = this.projects[groupIndex][projectIndex];
        axios
          .put(`http://localhost:8080/api/projects/${projectId}`, {
            name: project.name,
          })
          .then((response) => {
            const data = response.data;
            if (data.status === 'error') {
              // TODO: Something happened flash message?
              // TODO: Put the old value
              reject();
            } else if (data.status === 'ok') {
              resolve();
            }
          });
      });
      projectUpdated.then(() => {
        this.toggleProjectEdit(ev, projectId, projectIndex, groupIndex);
      });
    },
    createNewProject() {
      const projectCreated = new Promise((resolve, reject) => {
        axios
          .post('http://localhost:8080/api/projects', {
            name: this.new_project_name,
          })
          .then((response) => {
            const data = response.data;
            if (data.status === 'error') {
              // TODO: Something happened flash message?
              reject();
            } else if (data.status === 'ok') {
              const project = data.data;
              const newProject = {};
              const firstGroupOfProjects = this.projects[0];
              newProject.name = project.Name;
              newProject.id = project.ID;
              newProject.tasks = [];
              if (firstGroupOfProjects) {
                // eslint-disable-next-line no-console
                console.log(firstGroupOfProjects.length);
                if (firstGroupOfProjects.length === 3) {
                  this.projects.unshift([newProject]);
                } else {
                  firstGroupOfProjects.unshift(newProject);
                }
              } else {
                this.projects.unshift([newProject]);
              }
            }
            this.new_project_name = '';
            resolve();
          });
      });

      projectCreated.then(() => {
        // TODO: add flash message
      });
    },
    createNewTask(projectId, projectIndex, groupIndex) {
      const taskCreated = new Promise((resolve, reject) => {
        axios
          .post('http://localhost:8080/api/tasks', {
            name: this.form.task_name[projectId],
            project_id: projectId,
          })
          .then((response) => {
            const data = response.data;
            if (data.status === 'error') {
              // TODO: Something happened flash message?
              reject();
            } else if (data.status === 'ok') {
              const task = data.data;

              const newTask = {
                name: task.Name,
                id: task.ID,
                edit: false,
                description: '',
                status: 'new',
                finishedAt: this.createDate(),
              };
              const selectedProject = this.projects[groupIndex][projectIndex];
              selectedProject.tasks.unshift(newTask);
            }
            this.form.task_name[projectId] = '';
            resolve();
          });
      });

      taskCreated.then(() => {
        // TODO: add flash message
      });
    },
    createDate() {
      const d = new Date(2010, 7, 5);
      const ye = new Intl.DateTimeFormat('en', { year: 'numeric' }).format(d);
      const mo = new Intl.DateTimeFormat('en', { month: 'short' }).format(d);
      const da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(d);

      return `${da}-${mo}-${ye}`;
    },
  },
  created() {
    let projects;
    let tasks;
    const projectsDone = new Promise((resolve, reject) => {
      axios.get('http://localhost:8080/api/projects').then((response) => {
        const data = response.data;

        if (data.status === 'error') {
          // TODO: Something happened flash message?
          reject();
        } else if (data.status === 'ok') {
          projects = data.data;
        }

        resolve();
      });
    });
    const tasksDone = new Promise((resolve, reject) => {
      axios.get('http://localhost:8080/api/tasks').then((response) => {
        const data = response.data;

        if (data.status === 'error') {
          // TODO: Something happened flash message?
          reject();
        } else if (data.status === 'ok') {
          tasks = data.data;
        }
        resolve();
      });
    });

    Promise.all([projectsDone, tasksDone]).then(() => {
      const aggregateRoot = [];
      const step = 3;
      let index = 0;
      for (let i = 0; i < Math.ceil(projects.length / 3); i += 1) {
        const projectsArray = [];
        for (let j = 0; j < 3; j += 1) {
          // eslint-disable-next-line no-mixed-operators
          const currentProject = projects[step * index + j];

          if (!currentProject) break;
          const projectWithTasks = {};
          projectWithTasks.name = currentProject.Name;
          projectWithTasks.id = currentProject.ID;
          projectWithTasks.edit = false;
          projectWithTasks.tasks = [];
          for (let k = 0; k < tasks.length; k += 1) {
            if (tasks[k].ProjectID === currentProject.ID) {
              const task = {
                id: tasks[k].ID,
                name: tasks[k].Name,
                description: tasks[k].Description,
                edit: false,
                status: tasks[k].Status,
                finishedAt: tasks[k].finished_at,
              };
              projectWithTasks.tasks.push(task);
            }
          }
          projectsArray.push(projectWithTasks);
        }
        index += 1;
        aggregateRoot.push(projectsArray);
      }
      this.projects = aggregateRoot;
    });
  },
};
</script>
<style>
</style>
