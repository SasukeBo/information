import VueRouter from 'vue-router'
import { defaultRoutes } from './routes.js'

const router = new VueRouter({
  mode: 'history',
  routes: defaultRoutes()
})

export default router;
