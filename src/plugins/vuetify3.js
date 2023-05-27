// Vuetify
import "@mdi/font/css/materialdesignicons.css";
import "vuetify/styles";
import { createVuetify } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

export const vuetify3 = createVuetify({
  ssr: true,
  components,
  directives,
  // https://vuetifyjs.com/en/features/theme/
  theme: {
    // defaultTheme: "dark",
    themes: {
      light: {
        dark: true,
        colors: {
          primary: "#BB86FC",
          primaryVariant: "#3700B3",
          accent: "#22f1bf",
          error: "#CF6679",
          textPrimary: "#FFF",
          textSecondary: "",
          dark: "#121212",
          border: "#434343",
          "24dp": "#383838",
          "16dp": "#363636",
          "12dp": "#333",
          "8dp": "#2e2e2e",
          "6dp": "#2c2c2c",
          "4dp": "#272727",
          "3dp": "#252525",
          "2dp": "#232323",
          "1dp": "#1e1e1e",
        },
      },
    },
  },
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
});
