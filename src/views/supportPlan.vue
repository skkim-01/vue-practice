<template>
  <v-row class="title"><h2>SUPPORT PLAN</h2></v-row>
  <v-row class="title"><h5>List of support plan.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="pa-1" />
  <v-row class="support-plan-body" justify="space-between">
    <v-col cols="12">
      <!-- Search Field -->
      <v-container class="pa-1">
        <v-row>
          <v-col cols="2" align-self="center">
            <div class="text-caption">SEARCH</div>
          </v-col>
          <v-col cols="9">
            <v-text-field class="pt-5"></v-text-field>  
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="2" align-self="center">
            <div class="text-caption">FILTER</div>
          </v-col>

          <v-col cols="9" align-self="center">
            <v-row>
                <v-col cols="4" align-self="center">
                    <v-select align-self="center" class="pt-5"
                        :items="['filter name']"
                        variant="outlined"
                    />
                </v-col>
                <v-col cols="8" align-self="center">
                    <div class="text-caption">
                        <v-select align-self="center" class="pt-5"
                            :items="['filter actions']"
                            variant="outlined"
                        />
                    </div>
                </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-container>
    
      <!-- divider -->
      <v-divider />
      <v-container class="pa-2" />
      <v-row justify="center">
        <v-btn class="ma-1" @click="onClickedFind"> 찾기 </v-btn>
      </v-row>

      <v-container class="pa-2" />
      <v-divider />

      <!-- Space -->
      <v-container class="ma-2" />
    
      <!-- Table -->
      <TableCard />

    </v-col>
  </v-row>
</template>
  
  
<script setup>
const onClickedFind = async () => {
    console.log("#DBG\tonClickedFind");
}
</script>


<style lang="scss" scoped>
.title {
  margin: 0;
}
</style>
