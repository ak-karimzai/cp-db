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
              <router-link class="text-decoration-none" :to="`/admin/apartments/${apartment.id}`">
                {{ apartment.size }}
              </router-link>
            </td>
            <td>{{ apartment.room_numbers }}</td>
            <td>
                {{ apartment.user.last_name }} ,
                {{ apartment.user.first_name }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import Security from "./security";

export default {
  name: "AdminApartments",
  data() {
    return {
      ready: false,
      apartments: [],
    };
  },
  beforeMount() {
    Security.isAdmin();
    fetch(
      `${process.env.VUE_APP_SERVER_API}/apartments`,
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