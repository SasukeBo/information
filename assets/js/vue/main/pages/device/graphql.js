import gql from 'graphql-tag'

var apollo = {
  device: {
    query: gql`
    query ($uuid: String!) {
      device: deviceGet(uuid: $uuid) {
        name
        type
        description
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
