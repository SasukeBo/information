import gql from 'graphql-tag'

var apollo = {
  device: {
    query: gql`
    query ($uuid: String!){
      device: deviceGet(uuid: $uuid) {
        uuid
        mac
        type
        name
        description
        status
        token
        createdAt
        updatedAt
        register: user {
          phone
          userExtend { name email }
        }
      }
    }
    `,
    variables() {
      return {
        uuid: this.uuid
      }
    }
  }
}

export {
  apollo
}
