<template>
    <div v-if="error" class="error">
      <p>An error occurred: {{ error }}</p>
    </div>
  </template>
  
  <script setup lang="ts">
  
  const props = defineProps({
    error: String
  })
  
  
  </script>
  
  <style scoped>
  .error {
    background-color: #f8d7da;
    border: 1px solid #f5c6cb;
    color: #721c24;
    padding: 10px;
    margin-top: 10px;
  }
  </style>
  