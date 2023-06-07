<template>
  <div class="container">
    <div class="col">
      <div class="row">
        <h1 class="mt-3">Service</h1>
        <hr />
        <form-tag v-if="this.ready">
          <text-input
            label="Name"
            name="name"
            value="service.name"
            type="text"
            v-model="service.name"
            required="true"
          ></text-input>

          <label for="description">Description</label><br />
          <textarea
            class="form-control"
            name="description"
            id="description"
            rows="3"
            v-model="service.description"
          ></textarea>

          <text-input
            label="Measurment Amount"
            name="m_amount"
            value="service.m_amount"
            type="text"
            v-model="service.m_amount"
            required="true"
          ></text-input>

          <text-input
            label="Cost"
            name="cost"
            value="service.cost"
            type="number"
            v-model="service.cost"
            required="true"
          ></text-input>
          <hr />

          <div class="float-start">
            <input
              type="submit"
              class="btn btn-primary me-2"
              :value="this.service.id === 0 ? 'Create' : 'Update'"
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
import Security from "./security";
import TextInput from "./forms/TextInput.vue";

export default {
  name: "ServiceEdit",
  components: {
    TextInput,
  },
  data() {
    return {
      ready: false,
      service: {
        id: 0,
        name: "",
        description: "",
        m_amount: "",
        cost: 0,
      },
    };
  },
  beforeMount() {
    Security.isAdmin();
    if (this.$route.params.serviceId != 0) {
      fetch(
        `${process.env.VUE_APP_SERVER_API}/services/${this.$route.params.serviceId}`,
        Security.requestOptionsWithoutBody("GET")
      )
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            this.$emit("error", response.message);
          } else {
            this.service = response.data;
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