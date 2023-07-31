<script setup lang="ts">
import { onMounted, ref } from 'vue'

import TodoList from './components/TodoList.vue'
import { axiosInstance } from './config/axios'
import PageHeader from './components/PageHeader.vue'




export interface Todo {
  id: number
  name: string
  finished: boolean
  editFlag: boolean
}
const todoList = ref<Todo[]>([])

async function addTodo(text: any) {
  try {
    const res = await axiosInstance.post('/todo', { name: text })
    res.data.editFlag = false
    todoList.value.unshift(res.data)
  } catch (error) {
    console.log(error)
  }
}
async function getTodoList() {
  try {
    const res = await axiosInstance.get('/todo')

    if (!res.data) return

    console.log(res.data)
    res.data.forEach((element: Todo) => {
      let todo: Todo = {
        id: element.id,
        name: element.name,
        finished: element.finished ? true : false,
        editFlag: false
      }
      todoList.value.unshift(todo)
    })
    console.log('todolist', todoList.value[0])
  } catch (error) {
    console.log(error)
  }
}
async function deleteTodo(todo: Todo) {
  try {
    const res = await axiosInstance.delete(`/todo/${todo.id}`)
    todoList.value = todoList.value.filter((x) => x !== todo)
  } catch (error) {
    console.log(error)
  }
}
async function updateTodo(newName: string, todo: Todo) {
  try {
    console.log(todo.name)
    const res = await axiosInstance.put('/todo', {
      name: newName,
      id: todo.id,
      finished: todo.finished ? 1 : 0
    })

    console.log(res)
    todo.editFlag = false
  } catch (error) {
    console.log(error)
  }
}
function toggleEditFlag(todo: Todo) {
  todo.editFlag = !todo.editFlag
}
function toggleItemFinish(todo: Todo) {
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
