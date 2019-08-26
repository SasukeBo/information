import gql from 'graphql-tag'

var apollo = {
  devicePrivs: {
    query: gql`
    query ($privType: PrivType){
      devicePrivs: privilegeList(privType: $privType) {
        key: id
        label: name
      }
    }`,
    variables() {
      return {
        privType: 'device'
      }
    }
  }
}

export {
  apollo
}
