<template>
  <v-row class="title"><h2>CASE DETAIL</h2></v-row>
  <v-row class="title"><h5>Let me know details of your case.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="ma-1" />

  <!-- Main Contents -->
  <v-card style="height:500px">
    <v-card-title>[TODO: Remove] Case Open Detail Contents</v-card-title>
  </v-card>

  <v-container class="ma-1" />
  <!-- Action buttons -->
  <v-row justify="center">
    <v-btn class="ma-1" @click="onClickedBefore"> 이전 </v-btn>
    <v-btn class="ma-1" @click="onClickedCreateCase"> 만들기 </v-btn>
  </v-row>
</template>

<script setup>
import { useRouter } from 'vue-router';
const router = useRouter()

const onClickedBefore = () => {
  router.push({ path: '/case/open' });
}

const onClickedCreateCase = async () => {
  // TODO: API CALL
  console.log("#DBG\tonClickedCreateCase()");
  setTimeout(() => {
    router.push({ path: '/dashboard' });
  }, 2000);
} 
</script>
  
  
<style lang="scss" scoped>
.title {
  margin: 0;
}
</style>

