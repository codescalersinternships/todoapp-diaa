<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
axios.defaults.baseURL = import.meta.env.VITE_BACKEND_BASE_URL
interface Todo {
  id: number
  name: string
  finished: boolean
  editFlag: boolean
}
const todoList = ref<Todo[]>([])
const text = ref('')
async function addTodo() {
  try {
    const res = await axios.post(
      '/todo',
      { name: text.value },
      {
        headers: {
          'Content-Type': 'application/json'
        }
      }
    )
    res.data.editFlag = false
    todoList.value.unshift(res.data)
  } catch (error) {
    console.log(error)
  }
}
async function getTodoList() {
  try {
    const res = await axios.get('/todo')

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
  } catch (error) {
    console.log(error)
  }
}
async function deleteTodo(todo: Todo) {
  try {
    const res = await axios.delete(`/todo/${todo.id}`)
    todoList.value = todoList.value.filter((x) => x !== todo)
  } catch (error) {
    console.log(error)
  }
}
async function updateTodo(todo: Todo) {
  try {
    console.log(todo.name)
    const res = await axios.put('/todo', {
      name: todo.name,
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
  updateTodo(todo)
}
onMounted(() => {
  getTodoList()
})
</script>
<template>
  <div class="app">
    <div class="header">
      <link
        rel="stylesheet"
        href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.7.2/font/bootstrap-icons.css"
      />
      <link
        href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css"
        rel="stylesheet"
        integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC"
        crossorigin="anonymous"
      />
      <h2>My To Do List</h2>
      <input type="title" placeholder="Title..." v-model="text" />
      <button type="button" class="add-btn" @click="addTodo">Add</button>
    </div>

    <div class="list">
      <ul v-for="todo in todoList" :key="todo.id">
        <li :class="`${todo.finished && 'checked'} ${todo.id % 2 == 1 && 'odd'}`">
          <span @click="deleteTodo(todo)" class="delete bi bi-trash blue-color"></span>
          <input
            type="checkbox"
            name=""
            id=""
            @click="toggleItemFinish(todo)"
            v-model="todo.finished"
          />
          <span v-show="!todo.editFlag"> {{ todo.name }}</span>
          <span v-show="!todo.editFlag" @click="toggleEditFlag(todo)" class="edit bi bi-pen"></span>
          <span v-show="todo.editFlag" @click="updateTodo(todo)" class="save bi bi-check2"></span>
          <input
            v-show="todo.editFlag"
            @keydown.enter="updateTodo(todo)"
            type="text"
            name=""
            id=""
            v-model="todo.name"
          />
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
ul {
  margin: 0;
  padding: 0;
}

li i {
  left: 13px;
  position: absolute;
}

ul li {
  cursor: pointer;

  position: relative;
  padding: 12px 8px 12px 40px;
  list-style-type: none;
  background: #eee;
  font-size: 18px;
  transition: 0.2s;
  color: black;
}

ul li:hover {
  background: #ddd;
}
.odd {
  background-color: rgb(183, 183, 183);
}

ul li.checked {
  background: #888;
  color: #fff;
  text-decoration: line-through;
}

.delete {
  position: absolute;
  right: 0;
  top: 0;
  padding: 12px 16px 12px 16px;
}

.delete:hover {
  background-color: #f44336;
  color: white;
}

.edit {
  position: absolute;
  right: 40px;
  top: 0;
  padding: 12px 16px 12px 16px;
}

.edit:hover {
  background-color: rgb(222, 125, 35);
  color: white;
}

.header {
  /* background-color: darkblue; */
  padding: 30px 40px;
  color: darkblue;
  text-align: center;
}

.add-btn {
  cursor: pointer;
  margin-left: 10px;
  height: 35px;
  text-align: center;
  background-color: rgb(255, 255, 255);
  width: 100px;
}

.add-btn:hover {
  background-color: #0065e1;
  color: white;
}

.save {
  position: absolute;
  right: 40px;
  top: 0;
  padding: 12px 16px 12px 16px;
}

.save:hover {
  background-color: rgb(8, 140, 70);
  color: white;
}

input[type='checkbox'] {
  margin-right: 10px;
  margin-top: 6px;
  width: 17px;
  height: 17px;
  cursor: pointer;
}

.app {
  width: 40%;
  margin: 20px auto;
  margin-top: 0px;
}
</style>
