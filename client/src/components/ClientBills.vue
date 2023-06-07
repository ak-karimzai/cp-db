<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Bills</h1>
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
            <tr v-for="bill in bills" :key="bill.id">
              <td>{{ new Date(bill.from).toLocaleString() }}</td>
              <td>{{ new Date(bill.until).toLocaleString() }}</td>
              <td>{{ bill.spend_amount }}</td>
              <td>{{ bill.payment_amount }}</td>
              <td>{{ bill.paid === true ? "Paid" : "Unpaid" }}</td>
              <td
                v-if="bill.paid === false"
                class="text-success"
                @click="pay(bill.id)"
              >
                Pay
              </td>
            </tr>
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
      `${process.env.VUE_APP_SERVER_API}/bills?aprId=${this.$route.params.aprId}`,
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
  methods: {
    pay(billId) {
      const payload = {
        bill_id: billId,
        user_id: this.store.user.id,
      };
      fetch(
        `${process.env.VUE_APP_SERVER_API}/payments`,
        Security.requestOptions(payload)
      ).then(response => response.json()).then(response => {
        if (response.error) {
          this.$emit('error', response.error.message);
        } else {
          this.$emit('success', 'payment successfull!');
        }
      }).catch(err => this.$emit('error', err));
    },
  },
};
</script>