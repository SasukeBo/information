import gql from 'graphql-tag';
// import { parseGQLError } from '../../utils'

var apollo = {
  lastLogin: {
    query: gql`
    query {
      lastLogin: getLastLogin {
        remoteIP
        userAgent
        createdAt
      }
    }`
  },
  thisLogin: {
    query: gql`
    query {
      thisLogin: userLoginGet {
        remoteIP
        userAgent
        createdAt
      }
    }`
  }
}

export {
  apollo
}
