import { defineNuxtConfig } from "nuxt/config";

export default defineNuxtConfig({
  plugins: ["~/plugins/cookies.ts"],
  modules: ["@nuxtjs/axios"],
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:7777",
    },
  },
});
