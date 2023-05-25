<template>
  <div class="d-flex align-center justify-center" style="height: 100vh">
    <v-sheet width="400" class="mx-auto">
      <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Login</v-toolbar-title>
          </v-toolbar>
        </v-card>

      <v-form>
        <v-text-field
          name="login"
          label="Login"
          type="text"
        ></v-text-field>
        <v-text-field
          id="password"
          name="password"
          label="Password"
          type="password"
        ></v-text-field>
          <a href="#" class="text-body-2 font-weight-regular">Forgot Password?</a>

          <v-btn type="submit" color="primary" block class="mt-2" @click="onClickedSignIn">Sign in</v-btn>

      </v-form>
      <div class="mt-2">
          <p class="text-body-2">Don't have an account? <a href="#">Sign Up</a></p>
      </div>
    </v-sheet>
  </div>
</template>
  
<script>
  export default {
  name: 'login',
  data() {
  return {
      user: '',
      password: '',
    };
  },
  methods: {
    async onClickedSignIn() {
      console.log("onclicked" + this.user + ", " + this.password);
      this.$router.push({ path: '/dashboard' })
    },
  },
};
</script>

<style>
p {
  margin: 20;
}
</style>