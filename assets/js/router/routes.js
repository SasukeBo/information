import { load, denyIfLoggedIn } from './utils'

function defaultRoutes() {
  return [
    {
      path: '/',
      name: 'index',
      component: load('main'),
      redirect: '/home',
      children: [
        {
          path: 'home',
          name: 'home',
          component: load('main/home')
        },
        {
          path: 'product',
          component: load('main/product'),
          children: [
            {
              path: 'new',
              name: 'product-new',
              component: load('main/product/new')
            },
            {
              path: ':id/show',
              name: 'product-show',
              props: true,
              component: load('main/product/show')
            },
            {
              path: 'list',
              name: 'product-list',
              component: load('main/product/list')
            }
          ]
        },
        {
          path: 'device',
          component: load('main/device'),
          children: [
            {
              path: 'list',
              name: 'device-list',
              component: load('main/device/list')
            },
            {
              path: ':id/show',
              name: 'device-show',
              props: true,
              component: load('main/device/show')
            },
            {
              path: 'new',
              name: 'device-new',
              component: load('main/device/new')
            }
          ]
        }
      ]
    },
    {
      path: '/auth',
      component: load('auth'),
      beforeEnter: denyIfLoggedIn(),
      children: [
        {
          path: 'register',
          alias: '/register',
          name: 'register',
          component: load('auth/register')
        },
        {
          path: 'login',
          alias: '/login',
          name: 'login',
          component: load('auth/login')
        },
        {
          path: 'reset_password',
          alias: '/reset_password',
          name: 'reset_password',
          component: load('auth/reset')
        }
      ]
    },
    {
      path: '/admin',
      component: load('admin'),
      children: [
        {
          path: 'settings',
          name: 'system-conf',
          component: load('admin/settings')
        }
      ]
    },
    {
      path: '/*',
      name: 'error-page',
      component: load('errors/404.vue')
    }
  ]
}

export {
  defaultRoutes
}
