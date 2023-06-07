<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">User payments</h1>
        <hr />

        <table v-if="this.ready" class="table table-compact table-scriped">
          <thead>
            <th>Amount</th>
            <th>Bill</th>
            <th>Payment Date</th>
          </thead>
          <tbody>
            <tr v-for="payment in payments" :key="payment.id">
              <td>{{ payment.amount }}</td>
              <td>
                <router-link :to="`/bills/${payment.bill_id}`">
                  Bill
                </router-link>
              </td>
              <td>{{ new Date(payment.created_at).toLocaleString() }}</td>
            </tr>
          </tbody>
        </table>
        <h3 v-else>Loading...</h3>
      </div>
    </div>
  </div>
</template>

<script>
import Security from "./security.js";
import { store } from "./store.js";
export default {
  name: "ClientBills",
  data() {
    return {
      payments: [],
      ready: false,
      store,
    };
  },
  beforeMount() {
    fetch(
      `${process.env.VUE_APP_SERVER_API}/payments?userId=${this.store.user.id}`,
      Security.requestOptionsWithoutBody()
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit("error", response.error.message);
        } else {
          this.payments = response.data;
          this.ready = true;
        }
      })
      .catch((error) => {
        this.$emit("error", error.message);
      });
  },
};
</script>