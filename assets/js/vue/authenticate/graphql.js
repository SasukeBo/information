import gql from 'graphql-tag';

function login(app) {
  app.$apollo.mutate({
    mutation: gql`
      mutation loginByPassword(
        $phone: String!
        $password: String!
        $remember: Boolean
      ) {
        loginByPassword(
          phone: $phone
          password: $password
          remember: $remember
        )
      }
    `,
    variables: app.loginForm,
  }).then(({ data: { loginByPassword: r } }) => {
    app.$router.push({ name: 'home' });
    console.log("login success: ", r)
  }).catch(e => {
    app.message = e.message
    console.log(e);
  })
}

function sendSmsCode(app) {
  app.$apollo
    .mutate({
      mutation: gql`
            mutation sendSmsCode($phone: String!) {
              sendSmsCode(phone: $phone) {
                message
                code
              }
            }
          `,
      variables: {
        phone: app.$refs['form'].form.phone
      }
    })
    .then(({ data: { sendSmsCode: res } }) => {
      if (res.message !== 'OK') console.log(res.message);
    })
    .catch(e => {
      console.log(e);
    });
};

function register(app) {
  app.$apollo.mutate({
    mutation: gql`
      mutation register(
        $phone: String!
        $password: String!
        $smsCode: String!
      ) {
        register(
          phone: $phone
          password: $password
          smsCode: $smsCode
        ) {
          uuid
        }
      }
    `,
    variables: app.$refs['form'].form,
  }).then(({ data: { register: r } }) => {
    app.$message({
      message: '恭喜你，注册成功，请登录',
      type: 'success'
    });
    app.$router.push({ name: 'login' });
    console.log("uuid is", r.uuid)
  }).catch(e => {
    app.message = e.message
    console.log(e);
  })
}

function resetPassword(app) {
  app.$apollo.mutate({
    mutation: gql`
    mutation resetPassword(
      $phone: String!
      $smsCode: String!
      $password: String!
    ){
      resetPassword(
        phone: $phone
        smsCode: $smsCode
        password: $password
      ){
        uuid
      }
    }
    `,
    variables: app.$refs['form'].form,
  }).then(({ data: { resetPassword: r } }) => {
    app.$message({
      message: '重置密码成功，请登录',
      type: 'success'
    });
    app.$router.push({ name: 'login' });
    console.log("uuid is", r.uuid)
  }).catch(e => {
    app.message = e.message
    console.error(e);
  })
}

export default {
  login,
  register,
  sendSmsCode,
  resetPassword
};
