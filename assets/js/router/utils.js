import router from './index'

function load(path) {
  const component = () => import(`js/vue/${path}`);
  return component;
};

/*      导航守卫
---------------------- */

// 如果用户已登录，拒绝访问目的页面
function denyIfLoggedIn() {
  return (to, from, next) => {
    var app = router.app;
    // 如果用户未登录，继续访问
    // 否则返回原页面

    if (app.$store.state.user.phone) {
      next(from);
    } else {
      app.$store.dispatch('user/clearUserData');
      next();
    }
  };
}

export {
  load,
  denyIfLoggedIn
}
