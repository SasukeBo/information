import gql from 'graphql-tag'

var paramQuery = `
        id
        name
        sign
        type
        author {
          uuid
          userExtend {name}
        }
        createdAt
`

var apollo = {
  params: {
    query: gql`
    query(
      $deviceUUID: String!
      $namePattern: String
      $signPattern: String
      $type: DeviceParamValueType
      $userUUID: String
    ) {
      params: deviceParamList(
        type: $type
        deviceUUID: $deviceUUID
        namePattern: $namePattern
        signPattern: $signPattern
        userUUID: $userUUID
      ) {
        ${paramQuery}
      }
    }
    `,
    variables() {
      return {
        deviceUUID: this.uuid,
        namePattern: this.namePattern
      }
    }
  }
}

var deviceParamUpdate = app => {
  app.saving = true;
  return app.$apollo.mutate({
    mutation: gql`
    mutation(
      $id: Int!
      $name: String
      $sign: String
      $type: DeviceParamValueType
    ) {
      deviceParamUpdate(
        id: $id
        name: $name
        sign: $sign
        type: $type
      ) {
        ${paramQuery}
      }
    }`,
    variables: {
      ...app.form
    }
  })
}

var deviceParamCreate = app => {
  app.saving = true;
  return app.$apollo.mutate({
    mutation: gql`
    mutation(
      $deviceID: Int!
      $name: String!
      $type: DeviceParamValueType!
      $sign: String!
    ) {
      deviceParamCreate(
        name: $name
        sign: $sign
        type: $type
        deviceID: $deviceID
      ) {
        ${paramQuery}
      }
    }`,
    variables: {
      deviceID: app.deviceID,
      ...app.form
    }
  })
}

var deviceParamDelete = app => {
  app.deleting = true;
  return app.$apollo.mutate({
    mutation: gql`
    mutation($id: Int!) {
      deviceParamDelete(id: $id)
    }
    `,
    variables: { ...app.form }
  })
}

export {
  apollo,
  deviceParamUpdate,
  deviceParamCreate,
  deviceParamDelete
}
