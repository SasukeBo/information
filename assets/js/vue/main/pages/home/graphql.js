import gql from 'graphql-tag';

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
