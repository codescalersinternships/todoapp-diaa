<script setup lang="ts">
import { onMounted, ref } from 'vue'

import TodoList from './components/TodoList.vue'
import { axiosInstance } from './config/axios'
import PageHeader from './components/PageHeader.vue'
import ErrorComponent from './components/ErrorComponent.vue'

export interface Task {
  id: number
  name: string
  finished: boolean
  editFlag: boolean
}
const apiError = ref('')

const todoList = ref<Task[]>([])

async function addTodo(text: any) {
  try {
    const res = await axiosInstance.post('/todo', { name: text })
    res.data.editFlag = false
    todoList.value.unshift(res.data)
  } catch (error) {
    apiError.value = 'error adding todo'
    console.log(error)
  }
}
async function getTodoList() {
  try {
    const res = await axiosInstance.get('/todo')

    if (!res.data) return

    res.data.forEach((element: Task) => {
      let todo: Task = {
        id: element.id,
        name: element.name,
        finished: element.finished ? true : false,
        editFlag: false
      }
      todoList.value.unshift(todo)
    })
  } catch (error) {
    apiError.value = 'error getting todo list'
    console.log(error)
  }
}
async function deleteTodo(todo: Task) {
  try {
    const res = await axiosInstance.delete(`/todo/${todo.id}`)
    todoList.value = todoList.value.filter((x) => x !== todo)
  } catch (error) {
    apiError.value = 'error deleting todo'
    console.log(error)
  }
}
async function updateTodo(newName: string, todo: Task) {
  try {
    const res = await axiosInstance.put('/todo', {
      name: newName,
      id: todo.id,
      finished: todo.finished ? 1 : 0
    })

    todo.editFlag = false
  } catch (error) {
    apiError.value = 'error updating todo'
    console.log(error)
  }
}
function toggleEditFlag(todo: Task) {
  todo.editFlag = !todo.editFlag
}
function toggleItemFinish(todo: Task) {
  todo.finished = !todo.finished
  updateTodo(todo.name, todo)
}
onMounted(() => {
  getTodoList()
})
</script>
<template>
  <div class="app">
    <PageHeader @addTodo="addTodo" />

    <ErrorComponent :error="apiError" />

    <TodoList
      :todoList="todoList"
      @deleteTodo="deleteTodo"
      @updateTodo="updateTodo"
      @toggleEditFlag="toggleEditFlag"
      @toggleItemFinish="toggleItemFinish"
    />
  </div>
</template>

<style scoped>
.app {
  width: 40%;
  margin: 20px auto;
  margin-top: 0px;
}
</style>
