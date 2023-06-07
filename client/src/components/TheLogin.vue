<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <form-tag @myevent="submitHandler" name="myform" event="myevent">
          <text-input
            v-model="username"
            label="Username"
            type="text"
            name="username"
            required="true"
          ></text-input>
          <text-input
            v-model="password"
            label="Password"
            type="password"
            name="password"
            required="true"
          ></text-input>
          <hr />
          <input type="submit" class="btn btn-primary" value="Login" />
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import FormTag from "./forms/FormTag.vue";
import TextInput from "./forms/TextInput.vue";
import { store } from "./store.js";
import Security from "./security.js";
import router from "@/router";

export default {
  name: "TheLogin",
  components: {
    TextInput,
    FormTag,
  },
  data() {
    return {
      username: "",
      password: "",
      store,
    };
  },
  methods: {
    submitHandler() {
      const payload = {
        username: this.username,
        user_password: this.password,
      };
      fetch(`${process.env.VUE_APP_SERVER_API}/users/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      })
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            this.$emit("error", response.error.message);
          } else {
            this.store.token = response.data.token;
            this.store.user = response.data.user;
            Security.saveToken(response.data);
            router.push("/");
          }
        })
        .catch((err) => this.$emit("error", err));
    },
  },
};
</script>