import { createApp } from 'vue';
import App from '../app.vue';
import router from './src/plugins/router';
import { provideApolloClient } from '@vue/apollo-composable';
import { apolloClient } from './plugins/apollo';

const app = createApp(App);

app.use(router);

app.mount('#app');

provideApolloClient(apolloClient);
