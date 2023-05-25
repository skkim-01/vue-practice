<template>
  <v-app id="inspire">
    <div class="app-container">
      <toolbar @toggleNavigationBar="drawer=!drawer; console.log(drawer)"/>
      <navigation :rail="drawer"/>
      <v-content>
          <router-view> </router-view>
          <page-footer />
        <!-- <page-footer /> -->
      </v-content>
      
    </div>
  </v-app>
</template>

<script setup>
</script>
