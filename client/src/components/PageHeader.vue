<script setup lang="ts">
import { ref } from 'vue'
const text = ref('')
defineEmits(['addTodo'])
</script>

<template>
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
    <input @keydown.enter="$emit('addTodo', text)"  id="title" type="title" placeholder="Title..." v-model="text" />
    <button type="button" class="add-btn" @click="$emit('addTodo', text)">Add</button>
  </div>
</template>

<style>
.header {
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
</style>
