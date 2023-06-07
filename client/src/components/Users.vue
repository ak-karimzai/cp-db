<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Users</h1>
        <hr />

        <table v-if="this.ready" class="table table-compact table-scriped">
          <thead>
            <th>Last Name</th>
            <th>First Name</th>
            <th>Username</th>
            <th>Role</th>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id">
              <td>
                <router-link :to="`/admin/users/${user.id}`">
                  {{ user.last_name }}
                </router-link>
              </td>
              <td>{{ user.first_name }}</td>
              <td>{{ user.username }}</td>
              <td>{{ user.user_role }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Security from "./security.js";
import { store } from "./store.js";
export default {
  name: "AdminUsers",
  data() {
    return {
      users: [],
      ready: false,
      store,
    };
  },
  beforeMount() {
    Security.isAdmin();
    fetch(
      `${process.env.VUE_APP_SERVER_API}/users`,
      Security.requestOptionsWithoutBody()
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit('error', response.error.message);
        } else {
          this.users = response.data;
          this.ready = true;
        }
      }).catch(error => {
        console.log(error);
        this.$emit('error', error.message);
      });
  },
};
</script>