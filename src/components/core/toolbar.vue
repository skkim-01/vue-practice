<template>
    <v-toolbar
        class="v-app-bar"
        app
    >
        <v-toolbar-title>
            <v-app-bar-nav-icon @click="toggleNavigationBar"></v-app-bar-nav-icon>
        </v-toolbar-title>

        <template v-slot:append>
            <v-menu>
                <template v-slot:activator="{ props }">
                    <v-btn v-bind="props">
                        <v-text>UserName</v-text>
                    </v-btn>
                </template>
                <v-list>
                    <v-list-item
                        v-for="(item, index) in items"
                            :key="index"
                            :value="index"
                        @click=onClickedContext(index)
                    >
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>
          <v-btn icon="mdi-logout" @click=onClickedLogout></v-btn>
        </template>
    </v-toolbar>
</template>

<script>
export default {
    name: 'toolbar',
    data: () => ({
      items: [
        { title: 'Profile' },
        { title: 'My Issues' },
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
            case 0:
                this.$router.push('/user/profile');
                break;
            case 1:
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


<style lang="scss">
.v-app-bar {
    padding: {
      left: 32px;
      right: 32px;
    }
    position: relative;
    z-index: 1000;
    display: flex;
    align-items: center;
    justify-content: space-between;
  
    background-color: #333333;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.25);
}

.toolbar-menu-item {
    padding-left: 5px;
}

</style>