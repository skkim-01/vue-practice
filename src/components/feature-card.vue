<!-- <template>
    <v-card>
        <v-icon left> {{ props.icon }}</v-icon>
        <v-card-subtitle>
            {{ props.subs }}
        </v-card-subtitle>
        <v-card-text>
            {{ props.title }}
        </v-card-text>
        <v-card-actions>
            {{ props.sups }}
        </v-card-actions>
    </v-card>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
    icon: {
        type: Object,
        required: false,
    },
    title: {
        type: Object,
        required: true,
    },
    subs: {
        type: Object,
        required: false,
    },
    sups: {
        type: Object,
        required: false,
    }
})
</script> -->

<template>
    <v-card>
      <v-container pa-1>
        <v-layout row>
          <v-flex xs7>
            <v-card-title primary-title>
              <div>
                <div class="grey--text">{{title}}</div>
                <h3 class="headline">{{number}}</h3>
              </div>
            </v-card-title>
          </v-flex>
          <v-flex xs5>
            <v-card-title primary-title>
              <v-avatar>
                <v-icon medium :color="tIconColor">{{tIcon}}</v-icon>
              </v-avatar>
            </v-card-title>
          </v-flex>
        </v-layout>
        <v-divider light></v-divider>
        <v-card-actions class="pa-2">
          <v-icon small :color="bIconColor">{{bIcon}}</v-icon>
          <span class="grey--text caption font-italic">&nbsp;{{bText}}</span>
          <v-spacer></v-spacer>
        </v-card-actions>
      </v-container>
    </v-card>
  </template>
  <script>
  export default {
    props: ['title', 'number', 'tIcon', 'tIconColor', 'bIcon', 'bIconColor', 'bText']
  }
  </script>