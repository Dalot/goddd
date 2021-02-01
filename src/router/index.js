import axios from 'axios';
import Vue from 'vue';
import Router from 'vue-router';
import HelloWorld from '@/components/HelloWorld';
import Home from '@/components/Home';
import Login from '@/components/Login';
import SignUp from '@/components/SignUp';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: '/home',
    },
    {
      path: '/hello',
      name: 'HelloWorld',
      component: HelloWorld,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignUp,
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (localStorage.getItem('token') == null) {
      next({
        path: '/login',
        params: { nextUrl: to.fullPath },
      });
    } else {
      const token = localStorage.getItem('token');
      axios
      .get('http://localhost:8080/api', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
          console.log('ola');
          const data = response.data;

          if (data.status === 'error') {
            console.error('response:', response);
            console.error('UNAUTHORIZED');
            next({
              path: '/login',
              params: { nextUrl: to.fullPath },
            });
          } else if (data.status === 'ok') {
            next();
          }
        });
    }
  } else {
    next();
  }
});

export default router;
