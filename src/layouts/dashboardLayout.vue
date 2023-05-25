<template>
  <v-app>
      <toolbar app primary fixed />
      <navigation app dark fixed />
      <v-main>
        <v-container>
          <slot />
        </v-container>
        <page-footer />
      </v-main>
  </v-app>
</template>

<script setup>
</script>
