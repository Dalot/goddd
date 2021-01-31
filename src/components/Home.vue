<template>
  <div>
    <div class="ui center aligned basic segment">
      <div class="ui left icon action input">
        <i class="search icon"></i>
        <input type="text" placeholder="Project">
        <div class="ui blue submit button">Search</div>
      </div>
      <div class="ui horizontal divider">
        Or
      </div>
      <div class="ui left icon action input">
        <i class="plus icon"></i>
        <input type="text" placeholder="Project">
        <div class="ui positive submit button">Create New Project</div>
      </div>
    </div>
    <div
      class="ui three column doubling stackable grid container"
      v-for="(projectsDuo, index) in projects"
      v-bind:key="index"
    >
      <div
        class="column"
        v-for="project in projectsDuo"
        v-bind:key="project.name"
      >
        <div class="ui centered card">
          <div class="content">
            <div class="header">
              {{ project.name }}
              <a class="edit icon right floated">
                <i class="trash red icon"></i>
              </a>
              <a class="edit icon right floated">
                <i class="edit black icon"></i>
              </a>
            </div>
            <div class="description">
              <div class="ui centered card">
                <div
                  class="content"
                  v-for="task in project.tasks"
                  v-bind:key="task.name"
                >
                  <div class="description">
                    <div class="ui left floated compact">
                      <div class="ui fitted checkbox">
                        <input type="checkbox" />
                        <label></label>
                      </div>
                    </div>
                    {{ task.name }}
                    <a class="trash right floated">
                      <i class="trash  red icon"></i>
                    </a>
                    <a class="edit icon right floated">
                      <i class="edit black icon"></i>
                    </a>
                    <a class="check circle right floated">
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
            <input type="text" value="" />
            <button class="ui positive right icon button">
              <i class="plus icon"></i>
              Add Task
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type = "text/javascript" >
export default {
  data() {
    return {
      projects: [
        [
          {
            name: 'Project A',
            tasks: [
              {
                name: 'Task A',
                description: 'My description of task A',
              },
              {
                name: 'Task B',
                description: 'My description of task A',
              },
              {
                name: 'Task C',
                description: 'My description of task A',
              },
            ],
          },
          {
            name: 'Project B',
            tasks: [
              {
                name: 'Task A',
                description: 'My description of task A',
              },
              {
                name: 'Task B',
                description: 'My description of task A',
              },
              {
                name: 'Task C',
                description: 'My description of task A',
              },
            ],
          },
          {
            name: 'Project C',
            tasks: [
              {
                name: 'Task A',
                description: 'My description of task A',
              },
              {
                name: 'Task B',
                description: 'My description of task A',
              },
              {
                name: 'Task C',
                description: 'My description of task A',
              },
            ],
          },
        ],
        [
          {
            name: 'Project D',
            tasks: [
              {
                name: 'Task A',
              },
              {
                name: 'Task B',
              },
              {
                name: 'Task C',
              },
            ],
          },
          {
            name: 'Project E',
            tasks: [
              {
                name: 'Task A',
              },
              {
                name: 'Task B',
              },
              {
                name: 'Task C',
              },
            ],
          },
          {
            name: 'Project F',
            tasks: [
              {
                name: 'Task A',
              },
              {
                name: 'Task B',
              },
              {
                name: 'Task C',
              },
            ],
          },
        ],
      ],
    };
  },
};
</script>
<style>
</style>
