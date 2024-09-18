import { defineNuxtPlugin } from "#app";
import Cookie from "cookie-universal";

export default defineNuxtPlugin((nuxtApp) => {
  const cookies = Cookie();

  // Return as a plugin for use
  nuxtApp.provide("cookies", cookies);
});
