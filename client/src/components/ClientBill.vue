<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Payment Bill</h1>
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
            <td>{{ new Date(bill.from).toLocaleString() }}</td>
            <td>{{ new Date(bill.until).toLocaleString() }}</td>
            <td>{{ bill.spend_amount }}</td>
            <td>{{ bill.payment_amount }}</td>
            <td>{{ bill.paid === true ? "Paid" : "Unpaid" }}</td>
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
  name: "ClientBill",
  data() {
    return {
      bill: {},
      ready: false,
      store,
    };
  },
  beforeMount() {
    fetch(
      `${process.env.VUE_APP_SERVER_API}/bills/${this.$route.params.billId}`,
      Security.requestOptionsWithoutBody()
    )
      .then((response) => response.json())
      .then((response) => {
        if (response.error) {
          this.$emit("error", response.error.message);
        } else {
          this.bill = response.data;
          this.ready = true;
        }
      })
      .catch((error) => {
        this.$emit("error", error.message);
      });
  },
};
</script>