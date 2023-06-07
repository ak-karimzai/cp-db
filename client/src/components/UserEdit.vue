<template>
  <div class="container">
    <div class="col">
      <div class="row">
        <h1 class="mt-3"> {{ this.user.id === 0 ? 'Create' : 'Update' }} User</h1>
        <hr />
        <form-tag v-if="this.ready" :key="componetId">
          <text-input
            label="Last Name"
            name="last_name"
            value="user.last_name"
            type="text"
            v-model="user.last_name"
            required="true"
          ></text-input>

          <text-input
            label="First Name"
            name="first_name"
            value="user.first_name"
            type="text"
            v-model="user.first_name"
            required="true"
          ></text-input>

          <text-input
            label="Username"
            name="username"
            value="user.username"
            type="text"
            v-model="user.username"
            required="true"
          ></text-input>

          <text-input
            v-if="user.id === 0"
            label="Password"
            name="user_password"
            value="user.user_password"
            type="password"
            v-model="user.user_password"
            required="true"
          ></text-input>

          <text-input
            v-else
            label="Password"
            name="user_password"
            value="user.user_password"
            type="password"
            v-model="user.user_password"
            required="false"
          ></text-input>

          <div class="mb-3">
            <label for="role" class="form-label">Role</label>
            <select
              ref="userRole"
              class="form-select"
              id="role"
              name="role"
              :value="user.user_role"
              required
            >
              <option
                v-for="role in this.user_roles"
                :value="role"
                :key="role"
              >
                {{ role }}
              </option>
            </select>
          </div>

          <hr />
          <div class="float-start">
            <input
              type="submit"
              class="btn btn-primary me-2"
              :value="this.user.id === 0 ? 'Create' : 'Update'"
            />
            <router-link to="/admin/services" class="btn btn-outline-secondary"
              >Cancel</router-link
            >
          </div>
        </form-tag>
        <h1 v-else class="mt-3">Loading...</h1>
      </div>
    </div>
  </div>
</template>

<script>
import FormTag from "./forms/FormTag.vue";
import TextInput from "./forms/TextInput.vue";
import Security from "./security";
import { store } from "./store";

export default {
  name: "UserEdit",
  components: {
    FormTag,
    TextInput,
  },
  data() {
    return {
      ready: false,
      store,
      user: {
        id: 0,
        first_name: "",
        last_name: "",
        username: "",
        user_password: "",
        user_role: "",
      },
      componetId: 0,
      user_roles: ['admin', 'user'],
    };
  },
  beforeMount() {
    if (this.$route.params.userId != 0) {
      fetch(
        `${process.env.VUE_APP_SERVER_API}/users/${this.$route.params.userId}`,
        Security.requestOptionsWithoutBody()
      )
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            this.$emit("error", response.error.message);
          } else {
            this.user = response.data;
            this.ready = true;
          }
        })
        .catch((err) => this.$emit("error", err));
    } else {
      this.ready = true;
    }
  },
};
</script>