<template>
  <v-row class="title"><h2>SUPPORT CASE</h2></v-row>
  <v-row class="title pb-2"><h5>List of support case.</h5></v-row>
  <v-row>
    <v-divider />
  </v-row>
  <v-row>
    <v-container>
      <FormItem label="SEARCH">
        <v-col cols="8">
          <!-- 현재 search 크기는 density="compact"로 지정했지만 직접 크기 수정을 원하실 경우 
        class에 form-item작성 후 하단의 style에서 직접 수정 필요-->
          <v-text-field density="compact" hide-details></v-text-field>
        </v-col>
      </FormItem>
      <FormItem label="FILTER">
        <v-col class="pa-0" cols="8" align-self="center">
          <v-row>
            <v-col cols="4" align-self="center">
              <!-- select도 커스텀 하고 싶으면 density='compact'삭제 후 class='form-item'추가 및 하단의 style에서 padding값 변경 -->
              <v-select
                class="form-item pa-0"
                placeholder="select"
                v-model="state.filter1Val"
                align-self="center"
                :items="state.filterItems"
                variant="outlined"
                density="compact"
                hide-details
              />
            </v-col>
            <v-col cols="8" align-self="center">
              <v-select
                align-self="center"
                class=""
                :items="['filter actions']"
                variant="outlined"
                density="compact"
                hide-details
              />
            </v-col>
          </v-row>
        </v-col>
      </FormItem>
    </v-container>
  </v-row>
  <!-- divider -->
  <v-row>
    <v-divider />
  </v-row>
  <v-row align="center" justify="center">
    <v-col class="d-flex justify-center">
      <v-btn class="ma-1" @click="onClickedFind"> 찾기 </v-btn>
    </v-col>
  </v-row>

  <v-row>
    <v-divider />
  </v-row>
  <v-row>
    <!-- Table -->
    <TableCard class="ma-2 mt-8 supportcase-table" />
  </v-row>
</template>

<script>
import { reactive } from "vue";
import FormItem from "../components/FormItem.vue";
export default {
  name: "supportPlan",
  components: { FormItem },
  setup() {
    const state = reactive({
      filter1Val: "",
      filter2Val: "",
      filterItems: ["filter name", "filter name1", "filter name2"],
    });
    return { state };
  },
};
</script>

<style lang="scss">
.title {
  margin: 0;
}
.form-item {
  .v-field__field {
    input.v-field__input {
      // search 크기 조정(padding값으로 조정)
      padding: 15px;
    }
    div.v-field__input {
      // search 크기 조정(padding값으로 조정)
      input {
        padding: 15px;
      }
    }
  }
}
.supportcase-table {
  width: 100%;
}
</style>
