<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Manage Bills</h1>
        <hr />

        <table v-if="this.ready" class="table table-compact table-scriped">
          <thead>
            <th>From</th>
            <th>Until</th>
            <th>Spend Amount</th>
            <th>Payment Amount</th>
            <th>Status</th>
          </thead>
          <tbody>
            <tr
              v-for="bill in bills"
              :key="bill.id"
              @click="$router.push(`/admin/bills/${bill.id}`)"
            >
              <td>{{ new Date(bill.from).toLocaleString() }}</td>
              <td>{{ new Date(bill.until).toLocaleString() }}</td>
              <td>{{ bill.spend_amount }}</td>
              <td>{{ bill.payment_amount }}</td>
              <td>{{ bill.paid === true ? "Paid" : "Unpaid" }}</td>
            </tr>
          </tbody>
        </table>
        <p v-else>Loading...</p>
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
      bills: [],
      ready: false,
      store,
    };
  },
  beforeMount() {
    fetch(
      `${process.env.VUE_APP_SERVER_API}/bills`,
      Security.requestOptionsWithoutBody()
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit("error", response.error.message);
        } else {
          this.bills = response.data;
          this.ready = true;
        }
      })
      .catch((error) => {
        this.$emit("error", error.message);
      });
  },
};
</script>