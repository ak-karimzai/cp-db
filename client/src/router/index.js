import { createRouter, createWebHistory } from "vue-router";

import TheBody from "./../components/TheBody.vue";
import TheLogin from "./../components/TheLogin.vue";
import Users from "./../components/Users.vue"
import UserEdit from "./../components/UserEdit.vue"
import Services from "./../components/Services.vue"
import ServiceEdit from "./../components/ServiceEdit.vue"
import Apartments from "./../components/Apartments.vue"
import ApartmentEdit from "./../components/ApartmentEdit.vue"
import ClientApartments from "./../components/ClientApartments.vue"
import ClientBills from "./../components/ClientBills.vue"
import ApartmentServices from "./../components/ApartmentServices.vue"
import ClientPayments from "./../components/ClientPayments.vue"
import ClientBill from "./../components/ClientBill.vue"
import AdminBills from "./../components/AdminBills.vue"
import BillEdit from "./../components/BillEdit.vue"

const routes = [
  {
    path: "/",
    name: "Home",
    component: TheBody
  },
  {
    path: "/login",
    name: "Login",
    component: TheLogin
  },
  {
    path: "/admin/users",
    name: 'AdminUsers',
    component: Users
  },
  {
    path: "/admin/users/:userId",
    name: 'UserEdit',
    component: UserEdit
  },
  {
    path: "/admin/services",
    name: 'Services',
    component: Services
  },
  {
    path: "/admin/services/:serviceId",
    name: 'SerivceEdit',
    component: ServiceEdit
  },
  {
    path: "/admin/apartments",
    name: 'AdminApartment',
    component: Apartments
  },
  {
    path: "/admin/apartments/:apartmentId",
    name: 'ApartmentEdit',
    component: ApartmentEdit
  },
  {
    path: '/apartments',
    name: 'ClientApartments',
    component: ClientApartments
  },
  {
    path: '/apartments/bills/:aprId',
    name: 'AparmentBills',
    component: ClientBills
  },
  {
    path: '/apartments/services/:aprId',
    name: 'AparmentServices',
    component: ApartmentServices
  },
  {
    path: '/payments',
    name: 'ClientPayments',
    component: ClientPayments
  },
  {
    path: '/bills/:billId',
    name: 'ClientBill',
    component: ClientBill,
  },
  {
    path: '/admin/bills',
    name: 'AdminBills',
    component: AdminBills
  },
  {
    path: '/admin/bills/:billId',
    name: 'BillEdit',
    component: BillEdit
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

export default router;
