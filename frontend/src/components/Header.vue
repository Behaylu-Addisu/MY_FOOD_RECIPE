<template>
  <header class="header bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 text-white shadow-md py-10">
    <div class="container mx-auto flex flex-col items-center">
      <div class="flex justify-between items-center w-full">
        <div>
          <RouterLink to="/" class="container">
            <img src="../assets/images/mylogo.png" alt="logo" class="header-logo">
          </RouterLink>
          <p class="text-4xl text-white font-bold shadow-2xl ml-3"> ETHIOBEAN RECIPES </p>
        </div>
        <div class="flex space-x-4 items-center">
          <div class="search-bar relative">
            <input type="text" placeholder="Search..." class="bg-gray-200 text-rose-950 px-4 py-2 rounded-full focus:outline-none">
            <button class="absolute right-2 top-2 text-gray-400 hover:text-white">
              🔍
            </button>
          </div>
          <RouterLink to="/" class="bg-transparent border border-white text-white hover:bg-white hover:text-blue-500 px-4 py-2 rounded-full font-semibold">
            Login
          </RouterLink>
          <RouterLink to="/" class="bg-white text-blue-500 hover:bg-blue-500 hover:text-white px-4 py-2 rounded-full font-semibold"> Sign Up </RouterLink>
        </div>
      </div>
      <nav class="nav-links flex space-x-7 mt-4">

        <router-link to="/../pages/auth/Signup.vue" class="hover:text-gray-400 mt-3 mx-1">Home</router-link>
        <RouterLink to="/" class="hover:text-gray-400 mt-3 mx-1">Recipes</RouterLink>
        <RouterLink to="/" class="hover:text-gray-400 mt-3 mx-1">About</RouterLink>
      </nav>
    </div>
  </header>
</template>

<script>
export default {
}
</script>

<style scoped>
.header {
background-image: url("../assets/images/footerBackground.jpg");
background-repeat: no-repeat;
background-size: cover;
background-origin: border-box;
opacity: 0.7;
}
.header-logo {
width: 100px;
height: 100px;
border-radius: 50px;
position: absolute;
left: 5px;
top: 5px;
margin-right: 40px;
}

.nav-links {
justify-content: space-between;
}
.nav-links a {
text-decoration: none;
color: inherit;
}
</style>
