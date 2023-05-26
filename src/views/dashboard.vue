<template>
  <h1>dashboard</h1>
  <vue-echarts :option="chart_options" style="height: 500px" ref="chart" />
  <feature-card
    icon="mdi-chart-bar-stacked" 
    subs="issue"
    title="13"
    sups="Issue Opened"
  />
  <v-card>
    <v-card-text>

    </v-card-text>
  </v-card>
</template>

<script>
import { VueEcharts } from 'vue3-echarts';

export default {
    // data,
    // methods,
    data() {
      return {
        chart_options: {
          title: {
    text: 'World Population'
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {},
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  yAxis: {
    type: 'value',
    boundaryGap: [0, 0.01]
  },
  xAxis: {
    type: 'category',
    data: ['Brazil', 'Indonesia', 'USA', 'India', 'China', 'World']
  },
  series: [
    {
      name: '2011',
      type: 'bar',
      data: [18203, 23489, 29034, 104970, 131744, 630230]
    },
    {
      name: '2012',
      type: 'bar',
      data: [19325, 23438, 31000, 121594, 134141, 681807]
    }
  ]
        }
      }
    },
    components: {
        VueEcharts,
    },
};
</script>

<style>
</style>
