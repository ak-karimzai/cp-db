<template>
  <the-header></the-header>
  <div>
    <router-view
    :key="this.componentKey"
      @success="success"
      @error="error"
      @warning="warning"
      @forceUpdate="forceUpdate"
    ></router-view>
  </div>
  <the-footer></the-footer>
</template>

<script>
import TheHeader from "./components/TheHeader.vue";
import TheFooter from "./components/TheFooter.vue";
import notie from "notie";
import { store } from "./components/store";

const getCookie = (name) => {
  return document.cookie.split("; ").reduce((r, v) => {
    const parts = v.split("=");
    return parts[0] === name ? decodeURIComponent(parts[1]) : r;
  }, "");
};

export default {
  components: {
    TheHeader,
    TheFooter,
  },
  data() {
    return {
      store,
      componentKey: 0,
    };
  },
  methods: {
    success(message) {
      notie.alert({
        type: "success",
        text: message,
      });
    },
    error(message) {
      notie.alert({
        type: "error",
        text: message,
      });
    },
    warning(message) {
      notie.alert({
        type: "warning",
        text: message,
      });
    },
    forceUpdate() {
      this.componentKey++;
    }
  },
  beforeMount() {
    let data = getCookie("token");

    if (data !== "") {
      let cookieData = JSON.parse(data);
      store.token = cookieData.token;
      store.user = cookieData.user;
    }
  },
};
</script>

<style>
</style>