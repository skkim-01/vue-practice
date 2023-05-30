<template>
  <v-row class="title"><h2>CASE OPEN</h2></v-row>
  <v-row class="title"><h5>Let me know tpye, service, category and severity of your case.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="ma-1" />

  <!-- Main Contents -->
  <v-card style="height:500px">
    <v-card-title>[TODO: Remove] Case Open Contents</v-card-title>
  </v-card>

  <v-container class="ma-1" />
  <!-- Action buttons -->
  <v-row justify="center">
    <v-btn class="ma-1" @click="onClickedNext"> 다음 </v-btn>
  </v-row>
</template>

<script setup>
import { useRouter } from 'vue-router';
const router = useRouter()

const onClickedNext = () => {
  router.push({ path: '/case/open/detail' });
}
</script>
  
  
<style lang="scss" scoped>
.title {
  margin: 0;
}
</style>

