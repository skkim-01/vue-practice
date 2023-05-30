<template>
  <v-row class="title"><h2>CASE DETAIL</h2></v-row>
  <v-row class="title"><h5>Let me know details of your case.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="ma-1" />

  <!-- Main Contents -->
  <v-card>
    <v-container class="ma-1">
      <v-col cols="12">
        <v-card>
          <v-card-text>Subject</v-card-text>
          <v-card-subtitle>Briefly summarize your question or issue. 250 characters.</v-card-subtitle>
          <v-card-text>
            <v-text-field></v-text-field>
          </v-card-text>
        </v-card>
      </v-col>
      
      <v-col cols="12">
        <v-card>
          <v-card-text>Description</v-card-text>
          <v-card-subtitle>Describe your question or issue in detail. 8000 characters.</v-card-subtitle>
          <v-card-text>
            <v-text-field></v-text-field>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12">
        <v-card>
          <v-card-text>Attach files</v-card-text>
          <v-card-subtitle>You can attach up to 3 files. Each file can be up to 5 MB.</v-card-subtitle>
          <v-card-text>
            <v-text-field></v-text-field>
          </v-card-text>
        </v-card>
      </v-col>
    </v-container>
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

