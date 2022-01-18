import Vue from 'vue';
import Router from 'vue-router';
import routes from './router';
import iView from 'iview';
import config from '@/config';
const { homeName } = config;
const LOGIN_PAGE_NAME = 'login';


Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});


router.beforeEach((to: any, from: any, next: any) => {
  (iView as any).LoadingBar.start();
  const userName = sessionStorage.getItem('userName');
  if (!userName && to.name !== LOGIN_PAGE_NAME) { // 未登录且要跳转的页面不是登录页
    next({
      name: LOGIN_PAGE_NAME,
    });
  } else if (!userName && to.name === LOGIN_PAGE_NAME) { // 未登陆且要跳转的页面是登录页
    next();
  } else if (userName && to.name === LOGIN_PAGE_NAME) { // 已登录且要跳转的页面是登录页
    next({
      name: homeName, // 跳转到homeName页
    });
  } else {
    next();
  }
});


router.afterEach((route: any) => {
  (iView as any).LoadingBar.finish();
});

export default router;

