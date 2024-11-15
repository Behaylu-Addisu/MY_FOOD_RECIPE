<template>
    <div>
      <h1 class="text-2xl font-bold mb-4">Log In</h1>
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label for="email" class="block text-sm font-medium">Email</label>
          <input type="email" id="email" v-model="email" class="mt-1 block w-full border rounded-md">
        </div>
        <div class="mb-4">
          <label for="password" class="block text-sm font-medium">Password</label>
          <input type="password" id="password" v-model="password" class="mt-1 block w-full border rounded-md">
        </div>
        <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md">Log In</button>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const email = ref('');
  const password = ref('');
  
  const handleLogin = async () => {
    try {
      // Implement login logic here
      alert('Login successful!');
    } catch (error) {
      console.error(error);
      alert('Login failed!');
    }
  };
  </script>
  
  <style scoped>
  /* Add specific styles for the login page */
  </style>
  