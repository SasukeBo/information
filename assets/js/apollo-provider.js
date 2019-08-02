import Vue from 'vue'
import fetch from 'unfetch'
import { ApolloClient } from 'apollo-client'
import { HttpLink } from 'apollo-link-http'
// import { setContext } from 'apollo-link-context'
import { InMemoryCache } from 'apollo-cache-inmemory'
import VueApollo from 'vue-apollo'

/*          users            */
const httpLink = new HttpLink({
  fetch: fetch,
  uri: '/graphql'
})

// const authLink = setContext(({ variables }, { headers }) => {
  // var api = variables.graphql ;

  // if (api) return { headers: { ...headers, api } };
  // else return { headers };
// })

// 创建 apollo 客户端
const defaultClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
  connectToDevTools: true
})

const apolloProvider = new VueApollo({
  defaultClient,
  defaultOptions: {
    $query: {
      fetchPolicy: 'no-catch'
    }
  }
})

Vue.use(VueApollo)

export default apolloProvider
