import { createMemoryHistory, createRouter } from 'vue-router';
import Signup from '../pages/auth/Signup.vue';
import Login from '../pages/auth/Login.vue';
import Index from '../pages/index.vue';

const routes = [
  { path: '/', component: Index },
  { path: '/signup', component: Signup },
  { path: '/login', component: Login },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;
