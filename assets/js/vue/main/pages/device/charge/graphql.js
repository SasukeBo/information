import gql from 'graphql-tag'

var apollo = {
  charges: {
    query: gql`
        query($uuid: String) {
          charges: deviceChargeList(deviceUUID: $uuid) {
            id
            createdAt
            updatedAt
            user {
              uuid
              phone
              avatarURL
              userExtend {
                name
              }
            }
            deviceChargeAbilities{
              privilege {
                name
                sign
              }
            }
          }
        }
      `,
    variables() {
      return {
        uuid: this.uuid
      };
    }
  }
}

export {
  apollo
}
