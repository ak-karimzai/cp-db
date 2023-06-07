<template>
  <nav class="navbar navbar-expand navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">CP-DB</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item">
            <router-link
              v-if="store.token !== ''"
              class="nav-link active"
              aria-current="page"
              to="/apartments"
              >Apartments</router-link
            >
          </li>
          <li class="nav-item">
            <router-link
              v-if="store.token !== ''"
              class="nav-link active"
              aria-current="page"
              to="/payments"
              >Payments</router-link
            >
          </li>
          <li v-if="store.user.user_role === 'admin'" class="nav-item dropdown">
            <a
              href="#"
              class="nav-link dropdown-toggle"
              id="navBarDropDown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
              >Admin</a
            >
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <router-link class="dropdown-item" to="/admin/users"
                  >Manage Users</router-link
                >
              </li>
              <li>
                <router-link
                  class="dropdown-item"
                  :to="{ name: 'UserEdit', params: { userId: 0 } }"
                  >Add User</router-link
                >
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/services"
                  >Manage Services</router-link
                >
              </li>
              <li>
                <router-link
                  class="dropdown-item"
                  :to="{ name: 'SerivceEdit', params: { serviceId: 0 } }"
                  >Add Service</router-link
                >
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/apartments"
                  >Manage Apartments</router-link
                >
              </li>
              <li>
                <router-link
                  class="dropdown-item"
                  :to="{ name: 'ApartmentEdit', params: { apartmentId: 0 } }"
                  >Add Apartment</router-link
                >
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/bills"
                  >Manage Bills</router-link
                >
              </li>
              <li>
                <router-link class="dropdown-item" :to="{ name: 'BillEdit' , params: { billId: 0 } }"
                  >Add new Bill</router-link
                >
              </li>
            </ul>
          </li>
          <li class="nav-item">
            <router-link v-if="store.token === ''" class="nav-link" to="/login"
              >Login</router-link
            >
            <a href="javascript:void(0)" v-else class="nav-link" @click="logout"
              >Logout</a
            >
          </li>
        </ul>
        <span class="navbar-text">
          {{ store.user.first_name ?? "" }}
        </span>
      </div>
    </div>
  </nav>
</template>

<script>
import { store } from "./store";
import router from "@/router";
import Security from "./security";

export default {
  data() {
    return {
      store,
    };
  },
  methods: {
    logout() {
      Security.setTokenNull();
      this.store.token = "";
      this.store.user = {};
      this.$emit("forceUpdate");
      router.push("/login");
    },
  },
};
</script>

<style>
</style>