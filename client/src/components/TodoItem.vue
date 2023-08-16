<script setup lang="ts">
defineEmits(['updateTodo', 'deleteTodo', 'toggleItemFinish', 'toggleEditFlag'])

import { type Task } from '@/App.vue'
defineProps<{ todo: Task }>()
</script>
<template>
  <li :class="`${todo.finished && 'checked'} ${todo.id % 2 == 1 && 'odd'}`">
    <span @click="$emit('deleteTodo', todo)" class="delete bi bi-trash blue-color"></span>
    <input
      type="checkbox"
      name=""
      id="item-checkbox"
      @click="$emit('toggleItemFinish', todo)"
      :checked="todo.finished"
    />
    <span class="item-title" v-show="!todo.editFlag"> {{ todo.name }}</span>
    <span
      id="item-edit"
      v-show="!todo.editFlag"
      @click="$emit('toggleEditFlag', todo)"
      class="edit bi bi-pen"
    ></span>
    <span
      v-show="todo.editFlag"
      @click="$emit('updateTodo', todo)"
      class="save bi bi-check2"
    ></span>
    <input
      v-show="todo.editFlag"
      @keydown.enter="$emit('updateTodo', todo.name, todo)"
      type="text"
      name=""
      id="input"
      v-model="todo.name"
    />
  </li>
</template>

<style>
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
