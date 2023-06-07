<template>
  <div class="container">
    <div class="col">
      <div class="row">
        <h1 class="mt-3">
          {{ this.apartment.id === 0 ? "Create" : "Update" }} Apartment
        </h1>
        <hr />
        <form-tag v-if="this.ready">
          <text-input
            label="Size"
            name="size"
            value="apartment.size"
            type="number"
            v-model="apartment.size"
            required="true"
          ></text-input>

          <text-input
            label="Room Numbers"
            name="room_numbers"
            value="apartment.room_numbers"
            type="number"
            v-model="apartment.room_numbers"
            required="true"
          ></text-input>

          <div class="mb-3">
            <label for="User" class="form-label">User</label>
            <select
              ref="apruser"
              class="form-select"
              id="User"
              name="user"
              :value="this.apartment.user.last_name"
              required
            >
              <option
                v-for="user in this.users"
                :value="user.last_name"
                :key="user.id"
              >
                {{ user.last_name }}
              </option>
            </select>
          </div>

          <div class="float-start">
            <input
              type="submit"
              class="btn btn-primary me-2"
              @click="addOrEdit"
              :value="this.apartment.id === 0 ? 'Create' : 'Update'"
            />
            <router-link
              to="/admin/apartments"
              class="btn btn-outline-secondary"
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
import Security from "./security";
import TextInput from "./forms/TextInput.vue";
import router from "@/router";

export default {
  name: "ApartmentEdit",
  components: {
    TextInput,
  },
  data() {
    return {
      ready: false,
      apartment: {
        id: 0,
        size: "",
        room_numbers: "",
        user: {
          id: "",
          first_name: "",
          last_name: "",
        },
      },
      users: [],
    };
  },
  beforeMount() {
    Security.isAdmin();
    if (this.$route.params.apartmentId != 0) {
      fetch(
        `${process.env.VUE_APP_SERVER_API}/apartments/${this.$route.params.apartmentId}`,
        Security.requestOptionsWithoutBody("GET")
      )
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            this.$emit("error", response.message);
          } else {
            this.apartment = response.data;
            this.ready = true;
          }
        })
        .catch((err) => this.$emit("error", err));
    } else {
      this.ready = true;
    }
    fetch(
      `${process.env.VUE_APP_SERVER_API}/users`,
      Security.requestOptionsWithoutBody("GET")
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit("error", response.message);
        } else {
          this.users = response.data;
        }
      })
      .catch((err) => this.$emit("error", err));
  },
  methods: {
    addOrEdit() {
      const payload = {
        size: this.apartment.size,
        room_numbers: this.apartment.room_numbers,
        user_id: this.apartment.user.id,
      };
      payload.id = this.apartment.id === 0 ? "" : this.apartment.id;

      fetch(
        `${process.env.VUE_APP_SERVER_API}/apartments`,
        Security.requestOptions(payload)
      )
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            this.$emit("error", response.message);
          } else {
            this.$emit("success", "updated successfully!");
            router.push("/admin/apartments");
          }
        })
        .catch((err) => this.$emit("error", err));
    },
  },
};
</script>