<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="login">
      <div>
        <label for="username">Username:</label>
        <input v-model="username" id="username" type="text" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input v-model="password" id="password" type="password" required />
      </div>
      <button type="submit">Login</button>
      <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import { useCookies } from "~/plugins/cookies";

const username = ref("");
const password = ref("");
const errorMessage = ref("");

const router = useRouter();
const cookies = useCookies();

const login = async () => {
  try {
    const response = await axios.post("/api/login", {
      username: username.value,
      password: password.value,
    });
    const { token } = response.data;

    // Save token in cookies
    cookies.set("token", token, { maxAge: 3 * 60 }); // 3 minutes expiration

    // Redirect to homepage or another page
    router.push("/");
  } catch (error) {
    errorMessage.value = "Invalid username or password";
  }
};
</script>

<style scoped>
.error {
  color: red;
}
</style>
