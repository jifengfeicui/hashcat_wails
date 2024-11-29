<script setup>
import {NMessageProvider, NNotificationProvider} from "naive-ui";
import Layout from "./components/layout.vue";
</script>

<template>
  <div id="app">
    <n-modal-provider>
      <n-message-provider>
        <n-notification-provider>
          <layout></layout>
        </n-notification-provider>
      </n-message-provider>
    </n-modal-provider>
  </div>
</template>

<style>
#app {
  display: flex;
  height: 100vh;
  overflow: hidden;
  width: 100%;
}


</style>
