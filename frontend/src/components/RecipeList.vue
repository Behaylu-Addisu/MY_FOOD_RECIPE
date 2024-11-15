<template>
    <div>
      <h1>Recipes</h1>
      <ul>
        <li v-for="recipe in recipes" :key="recipe.recipe_id">
          <p class="text-2xl">{{ recipe.title }}</p>
          <p class="font-bold">{{ recipe.category }}</p>
          <p class="text-normal">{{ recipe.description }}</p>
        </li>
      </ul>
      <div v-if="loading">Loading...</div>
      <div v-if="error">Error: {{ error.message }}</div>
    </div>
  </template>
  
  <script>
  import { useQuery, gql } from '@vue/apollo-composable';
  
  const GET_RECIPES = gql`
    query GetRecipes {
      recipes {
        category
        description
        title
        created_at
        updated_at
        recipe_id
        user_id
      }
    }
  `;
  
  export default {
    setup() {
      const { result, loading, error } = useQuery(GET_RECIPES);
      return {
        recipes: result,
        loading,
        error,
      };
    },
  };
  </script>
  