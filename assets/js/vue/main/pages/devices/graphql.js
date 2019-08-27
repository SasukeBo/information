import gql from 'graphql-tag'

var apollo = {
  devices: {
    query: gql`
    query (
      $status: DeviceStatus
      $ownership: String
      $namePattern: String
    ) {
      devices: deviceList(
        status: $status
        ownership: $ownership
        namePattern: $namePattern
      ){
        uuid
        name
        type
        mac
        token
        user {
          uuid
          userExtend {
            name
          }
        }
        status
        description
        createdAt
        updatedAt
      }
    }
    `,
    variables() {
      var ownership = 'both'
      if (this.checkboxGroup.length === 1) ownership = this.checkboxGroup[0]

      return {
        ownership,
        namePattern: this.search
      }
    }
  }

}

export {
  apollo
}
