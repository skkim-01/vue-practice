<template>
  <v-row class="title"><h2>CASE OPEN</h2></v-row>
  <v-row class="title"><h5>Let me know tpye, service, category and severity of your case.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="ma-1" />

  <!-- Main Contents -->
  <v-card>
    <v-container class="ma-1">
      <v-col cols="12">
        <v-row>
          <v-col cols=6>
            <v-card>
              <v-card-title>Account And Billing</v-card-title>
              <v-card-text>Assistance for your account, such as billing, pricing, and reserved instances.</v-card-text>
            </v-card>
          </v-col>

          <v-col cols=6>
            <v-card>
              <v-card-title>Technical</v-card-title>
              <v-card-text>Support for service-related technical issues, such as Amazon EC2, Amazon S3 and more.</v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-col>

      <v-container class="pa-1" />
      <v-divider />
      <v-container class="pa-1" />

      <v-col cols="12">
        <v-card>
          <v-card-text>Service</v-card-text>
          <v-card-text>
            <v-select
              :items="['get', 'items', 'from', 'api']"
              variant="outlined"
            />
          </v-card-text>
        </v-card>
      </v-col>
      
      <v-col cols="12">
        <v-card>
          <v-card-text>Category</v-card-text>
          <v-card-text>
            <v-select
              :items="['get', 'items', 'from', 'api']"
              variant="outlined"
            />
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12">
        <v-card>
          <v-card-text>Severity</v-card-text>
          <v-card-text>
            <v-select
              :items="['get', 'items', 'from', 'api']"
              variant="outlined"
            />
          </v-card-text>
        </v-card>
      </v-col>
    </v-container>
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

