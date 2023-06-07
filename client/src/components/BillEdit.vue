<template>
  <div class="container">
    <div class="col">
      <div class="row">
        <h1 class="mt-3">
          {{ this.bill.id === 0 ? "Create" : "Update" }} Bill
        </h1>
        <hr />
        <form @submit.prevent="submitHandler" v-if="this.ready" event="myevent">
          <label for="from" class="form-label">From:</label>
          <input
            v-model="bill.from"
            type="date"
            id="from"
            class="form-control"
            required
          />

          <label for="until" class="form-label">Until:</label>
          <input
            v-model="bill.until"
            type="date"
            id="until"
            class="form-control"
            required
          />

          <label for="spend_amount" class="form-label">Spend Amount</label>
          <input
            type="number"
            id="spend_amount"
            class="form-control"
            v-model="bill.spend_amount"
            step="0.01"
          />

          <div v-if="this.$route.params.billId != 0">
            <label for="payment_amount" class="form-label"
              >Payment Amount</label
            >
            <label id="payment_amount" class="form-control">{{
              bill.payment_amount
            }}</label>
          </div>
          <div v-else>
            <label for="apartment" class="form-label">Apartment</label>
            <select v-model="apartment_id" id="apartment" class="form-select">
              <option
                v-for="apartment in apartments"
                :key="apartment.id"
                ref="apartment"
              >
                {{ apartment.size }}, {{ apartment.room_numbers }},
                {{ apartment.user.last_name }}
              </option>
            </select>
          </div>

          <hr />
          <div class="float-start">
            <input
              type="submit"
              class="btn btn-primary me-2"
              @click="submitHandler"
              :value="this.bill.id === 0 ? 'Create' : 'Update'"
            />
            <router-link to="/admin/bills" class="btn btn-outline-secondary"
              >Cancel</router-link
            >
          </div>
        </form>
        <p v-else>Loading...</p>
      </div>
    </div>
  </div>
</template>

<script>
import Security from "./security";
export default {
  name: "BillEdit",
  data() {
    return {
      ready: false,
      bill: {
        id: 0,
        from: "",
        until: "",
        spend_amount: 0.0,
        payment_amount: 0.0,
        paid: "false",
        apartment_id: "",
      },
      apartments: [],
      services: [],
    };
  },
  beforeMount() {
    if (this.$route.params.billId != 0) {
      fetch(
        `${process.env.VUE_APP_SERVER_API}/bills/${this.$route.params.billId}`,
        Security.requestOptionsWithoutBody()
      )
        .then((res) => res.json())
        .then((res) => {
          if (res.error) {
            this.$emit("error", res.error.message);
          } else {
            this.bill = res.data;
            this.bill.from = this.getDate(this.bill.from);
            this.bill.until = this.getDate(this.bill.until);
            this.ready = true;
          }
        })
        .catch((err) => {
          this.$emit("error", err);
        });
    } else {
      this.ready = true;
    }
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
  methods: {
    getDate(dateString) {
      const date = new Date(dateString);

      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, "0");
      const day = String(date.getDate()).padStart(2, "0");
      return `${year}-${month}-${day}`;
    },
    submitHandler() {
      if (this.from == "" || this.until == "") {
        this.$emit("error", "All fields are required");
        return;
      }
      const payload = {
        id: this.bill.id,
        from: new Date(this.bill.from).toISOString(),
        until: new Date(this.bill.until).toISOString(),
        spend_amount: this.bill.spend_amount,
        apartment_id: this.bill.apartment_id,
        service_id: this.bill.service_id,
      };
      console.log(this.user.apartment_id)
      console.log(payload);
      // fetch(
      //   `${process.env.VUE_APP_SERVER_API}/bills`,
      //   Security.requestOptions(payload)
      // )
      //   .then((res) => res.json())
      //   .then((res) => {
      //     if (res.error) {
      //       this.$emit("error", res.error.message);
      //     } else {
      //       this.$emit("success", "Changed successfully");
      //       this.$emit("forceUpdate");
      //     }
      //   })
      //   .catch((err) => this.$emit("error", err));
    },
    getApartmentServices(aprId) {
      console.log("aprId: ", aprId);
    },
  },
};
</script>
