import gql from 'graphql-tag'

var chartApollo = {
  params: {
    query: gql`
    query ($uuid: String) {
      params: deviceParamList(deviceUUID: $uuid) { name sign type }
    }`,
    variables() {
      return {
        uuid: this.device.uuid
      }
    }
  }
}

var apollo = {
  device: {
    query: gql`
    query ($uuid: String!){
      device: deviceGet(uuid: $uuid) {
        id
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
  apollo,
  chartApollo
}
