import Vue from 'vue'
import fetch from 'unfetch'
import { ApolloClient } from 'apollo-client'
import { HttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import { split } from 'apollo-link'
import { WebSocketLink } from 'apollo-link-ws'
import { getMainDefinition } from 'apollo-utilities'

import VueApollo from 'vue-apollo'

/*          users            */
const httpLink = new HttpLink({
  fetch: fetch,
  uri: '/graphql'
})

// 创建订阅的websocket连接
const wsLink = new WebSocketLink({
  uri: 'ws://localhost/websocket',
  options: {
    reconnect: true
  }
})

const link = split(
  ({ query }) => {
    const definition = getMainDefinition(query)
    return definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
  },
  wsLink,
  httpLink
)

// 创建 apollo 客户端
const defaultClient = new ApolloClient({
  link,
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
