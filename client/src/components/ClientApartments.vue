<template>
  <div class="container">
    <div class="row">
      <div class="col"></div>
      <table v-if="this.ready" class="table table-compact table-scriped">
        <thead>
          <th>Size</th>
          <th>Room Number</th>
          <th>Tenant</th>
        </thead>
        <tbody>
          <tr v-for="apartment in apartments" :key="apartment.id">
            <td>
              {{ apartment.size }}
            </td>
            <td>{{ apartment.room_numbers }}</td>
            <td>
              {{ apartment.user.last_name }} ,
              {{ apartment.user.first_name }}
            </td>
            <td>
              <router-link
                class="btn btn-outline-secondary"
                :to="`/apartments/bills/${apartment.id}`"
              >
                Bills
              </router-link>
            </td>
            <td>
              <router-link
                class="btn btn-outline-secondary"
                :to="`/apartments/services/${apartment.id}`"
              >
                Services
              </router-link>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import Security from "./security";
import { store } from "./store.js";

export default {
  name: "ClientApartments",
  data() {
    return {
      ready: false,
      apartments: [],
      store,
    };
  },
  beforeMount() {
    fetch(
      `${process.env.VUE_APP_SERVER_API}/apartments?userId=${this.store.user.id}`,
      Security.requestOptionsWithoutBody()
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit("error", response.error.message);
        } else {
          this.apartments = response.data;
          this.ready = true;
        }
      })
      .catch((err) => this.$emit("error", err));
  },
};
</script>