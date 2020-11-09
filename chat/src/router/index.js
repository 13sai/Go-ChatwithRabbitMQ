import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
  },
  {
    path: '/',
    redirect: '/chat/list' 
  },
  {
    path: '/chat/detail/:id',
    name: 'chatDetail',
    component: () => import('@/views/chat/detail'),
  },
  {
    path: '/chat/list',
    name: 'chatDetail',
    component: () => import('@/views/chat/list'),
  }
]


const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}


export default router

