<template>
  <v-app-bar color="primary">
        <v-toolbar-title>
            vuetify console
        </v-toolbar-title>

        <template v-slot:append>
            <v-menu>
                <template v-slot:activator="{ props }">
                    <v-btn v-bind="props">
                        <div class="text-body-2">사용자</div>
                    </v-btn>
                </template>
                <v-list>
                    <v-list-item
                        v-for="(item, index) in items"
                            :key="index"
                            :value="index"
                        @click=onClickedContext(index)
                    >
                    <v-list-item-title>
                        <div class="text-caption">
                            {{ item.title }}
                        </div>
                    </v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>
          <v-btn icon="mdi-logout" @click=onClickedLogout></v-btn>
        </template>
  </v-app-bar>
</template>

<script>
export default {
    name: 'toolbar',
    data: () => ({
      items: [
        { title: '내 관심 이슈' },
      ],
    }),
    methods: {
        toggleNavigationBar() {
            console.log("#DBG\ttoogleNavigationBar");
            const vm = this;
            vm.$emit('toggleNavigationBar');
        },
        onClickedContext(index) {
            switch(index) {
            // case 0:
            //     this.$router.push('/user/profile');
            //     break;
            case 0:
                this.$router.push('/user/issues');
                break;
            default:
            console.log("[WARN]\tUndefined Context");
            }
        },
        onClickedLogout() {
            // TODO: Logout - session clear
            this.$router.replace('/login');
        },
    },
};
</script>
