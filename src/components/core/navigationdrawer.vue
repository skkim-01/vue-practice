<template>
    <v-navigation-drawer
        width="200"
        v-model= "drawer"
        :rail="rail"
        permanent
    >        
        <!-- <v-list-item>
            <template v-slot:append>
                <v-btn
                    position="relative"
                    variant="text"
                    icon
                    @click.stop="rail=!rail"
                >
                    <v-icon>
                        {{ rail ? "mdi-chevron-right" : "mdi-chevron-left" }}
                    </v-icon>
                </v-btn>
            </template>
        </v-list-item> -->

        <v-list>
            <v-list-item>
                <v-btn
                    variant="text"
                    icon
                    @click.stop="rail=!rail"
                >
                    <v-icon>
                        {{ rail ? "mdi-chevron-right" : "mdi-chevron-left" }}
                    </v-icon>
                </v-btn>
            </v-list-item>
            <v-divider />

            <v-list-item
                v-for="(item, i) in items"
                :key="i"
                :value="item"
                active-color="primary"
                @click=onClickedItem(i)
            >
                <template v-slot:prepend>
                    <v-icon>{{ item.icon }}</v-icon>
                </template>

                <v-list-item-title v-text="item.text"></v-list-item-title>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>

<script>

export default {
    name: 'navigation',
   
    data: () => ({
      drawer: true,
      rail: true,
      items: [
        { text: 'Dashboard', icon: 'mdi-chart-bar-stacked' },
        { text: 'Support Case', icon: 'mdi-format-list-bulleted' },
        { text: 'Support Plan', icon: 'mdi-flag' },
        { text: 'Case Open', icon: 'mdi-open-in-new' },
      ],
    }),
    methods: {
        onClickedItem(item) {
            switch(item) {
                case 0:
                    this.$router.push('/dashboard');
                    break;
                case 1:
                    this.$router.push('/support/case');
                    break;
                case 2:
                    this.$router.push('/support/plan');
                    break;
                case 3:
                    this.$router.push('/case/open');
                    break;
                default:
                    console.log("[WARN]\tUndefined Router");
            }
        }
    }
};
</script>

<style>
.v-navigation-drawer {
    height: 100%;
    width: 90px;
    background-color: #555;
    position: fixed; /* Fixed Sidebar (stay in place on scroll) */
    /*z-index: 1;*/ /* Stay on top */
    top: 0; /* Stay at the top */
    left: 0;
    overflow-x: hidden; /* Disable horizontal scroll */
    padding-top: 0;
}
</style>