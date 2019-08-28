import gql from 'graphql-tag'

var newApollo = {
  users: {
    query: gql`
      query($namePattern: String) {
        users: userList(namePattern: $namePattern) {
          uuid
          userExtend {
            name
          }
        }
      }
    `,
    variables() {
      return {
        namePattern: this.queryString
      }
    }
  }
}

export { newApollo }
