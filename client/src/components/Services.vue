<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Services</h1>
        <hr />

        <table v-if="this.ready" class="table table-compact table-scriped">
          <thead>
            <th>Service Name</th>
            <th>Measurment unit</th>
            <th>Cost</th>
          </thead>
          <tbody>
            <tr v-for="service in services" :key="service.id">
              <td>
                <router-link :to="`/admin/services/${service.id}`">
                  {{ service.name }}
                </router-link>
              </td>
              <td>{{ service.m_amount }}</td>
              <td>{{ (service.cost) ?  service.cost : 0 }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Security from './security';

export default {
  name: "AdminServices",
  data() {
    return {
      services: [],
      ready: false,
    };
  },
  beforeMount() {
    Security.isAdmin();

    fetch(
      `${process.env.VUE_APP_SERVER_API}/services`,
      Security.requestOptionsWithoutBody('GET')
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit("error", response.message);
        } else {
          this.services = response.data;
          this.ready = true;
        }
      })
      .catch((err) => this.$emit("error", err));
  },
};
</script>