import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core';
import { provideApolloClient } from '@vue/apollo-composable';

const httpLink = new HttpLink({
  uri: 'http://localhost:8081/v1/graphql',
  // headers: {
  //   'x-hasura-admin-secret': 'getye',
  // },
});

const cache = new InMemoryCache();

const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});

provideApolloClient(apolloClient);

// export default defineNuxtPlugin((_nuxtApp) => {
//   provideApolloClient(apolloClient);
// });




// import { ApolloClient, InMemoryCache } from '@apollo/client/core';
// import { provideApolloClient } from '@vue/apollo-composable';

// const cache = new InMemoryCache();

// const apolloClient = new ApolloClient({
//   uri: 'https://your-graphql-endpoint.com/graphql',
//   cache,
// });

// provideApolloClient(apolloClient);
