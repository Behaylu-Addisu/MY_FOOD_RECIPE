<template>
  <div class= "m-auto w-96 my-8 align-middle shadow-2xl">
    <h1 class="text-center font-bold my-6 animate-pulse"> Signup </h1>

    <form @submit.prevent="onsubmit">

      <!-- Username Field -->
      <div>
        <label class="block font-bold text-l p-2 mb-2">User Name</label>
        <input
          name="username"
          type="text"
          class="block bg-stone-200 ml-6 py-2 px-2 w-80 rounded"
          rules="required|min:4"
          placeholder="User Name"
        />
        <ErrorMessage name="username" class="text-red-500 text-sm" />
      </div>

      <!-- Email input -->
      <div>
        <label class="block font-bold text-l p-2 mb-2">Email</label>
        <input
          name="email"
          type="email"
          class="block bg-stone-200 ml-6 py-2 px-2 w-80 rounded"
          rules="required|email"
          placeholder="Email"
        />
        <ErrorMessage name="email" class="text-red-500 text-sm" />
      </div>

      <!-- Password input -->
      <div>
        <label class="block font-bold text-l p-2 mb-2">Password</label>
        <input
          name="password"
          type="password"
          class="block bg-stone-200 ml-6 py-2 px-2 w-80 rounded"
          rules="required|min:6"
          placeholder="Password"
        />
        <ErrorMessage name="password" class="text-red-500 text-sm" />
      </div>

      <!-- Confirm Password input -->
      <div>
        <label class="block font-bold text-l p-2 mb-2">Confirm Password</label>
        <input
          name="confirmPassword"
          type="password"
          class="block bg-stone-200 ml-6 py-2 px-2 w-80 rounded"
          rules="required|same:password"
          placeholder="Confirm Password"
        />
        <ErrorMessage name="confirmPassword" class="text-red-500 text-sm" />
      </div>

      <!-- Submit Button -->
      <button
        type="submit"
        class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        Signup
      </button>
    </form>
  </div>
</template>

<script>
import { useForm, ErrorMessage } from "vee-validate";

export default {
  setup() {
    const { handleSubmit, isSubmitting } = useForm();

    const onsubmit = handleSubmit((values) => {
      alert(`Form submitted successfully! \n\n${JSON.stringify(values)}`);
    });

    return {
      onsubmit,
      isSubmitting,
    };
  },
};
</script>

<style scoped>
/* Add any additional styles here */
</style>
