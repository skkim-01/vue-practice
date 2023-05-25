<template>
  <div id="app">
    <appLayout />
  </div>
</template>

<script setup>
import appLayout from './layouts/appLayout.vue';

</script>

<style lang="scss">
#app {
  font-family: Spoqa Han Sans Neo;
  font-style: normal;
  font-weight: 400;
  letter-spacing: 0em;
  text-align: left;

  -webkit-font-smoothing: antialiased;
  -webkit-transform: translateZ(0);
  color: var(--mainFontColor);
  overflow: hidden;
}

body {
  white-space: nowrap;
  margin: 0;
  padding: 0;
}
</style>
