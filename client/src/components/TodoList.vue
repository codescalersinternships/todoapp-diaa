<script setup lang="ts">
defineEmits(['updateTodo', 'deleteTodo', 'toggleItemFinish', 'toggleEditFlag'])

import TodoItem from '@/components/TodoItem.vue'
import { type Task } from '@/App.vue'

const props = defineProps<{ todoList: Task[] }>()
</script>

<template>
  <div class="list">
    <ul v-for="todo in todoList" :key="todo.id">
      <TodoItem
        :todo="todo"
        @updateTodo="$emit('updateTodo', todo.name, todo)"
        @deleteTodo="$emit('deleteTodo', todo)"
        @toggleItemFinish="$emit('toggleItemFinish', todo)"
        @toggleEditFlag="$emit('toggleEditFlag', todo)"
      />
    </ul>
  </div>
</template>
<style>
ul {
  margin: 0;
  padding: 0;
}
</style>
