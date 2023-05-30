// https://router.vuejs.org/guide/migration/

import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/views/login';

import Dashboard from '@/views/dashboard';
import supportCase from '@/views/supportCase';
import supportPlan from '@/views/supportPlan';
import caseOpen from '@/views/caseOpen';
import caseOpenDetail from '@/views/caseOpenDetail';

import userProfile from '@/views/userProfile';
import userIssues from '@/views/userIssues';

const routes = [
  // TODO: /
  //  if sessioned : dashboard
  //  else : login
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      layout: 'defaultLayout'
    }
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Dashboard,
    meta: {
      layout: 'dashboardLayout'
    }
  },
  {
    path: '/support/case',
    name: 'supportCase',
    component: supportCase,
    meta: {
      layout: 'dashboardLayout'
    }
  },
  {
    path: '/support/plan',
    name: 'supportPlan',
    component: supportPlan,
    meta: {
      layout: 'dashboardLayout'
    }
  },
  {
    path: '/case/open',
    name: 'caseOpen',
    component: caseOpen,
    meta: {
      layout: 'dashboardLayout'
    }
  },
  {
    path: '/case/open/detail',
    name: 'caseOpenDetail',
    component: caseOpenDetail,
    meta: {
      layout: 'dashboardLayout'
    }
  },
  {
    path: '/user/profile',
    name: 'userProfile',
    component: userProfile,
    meta: {
      layout: 'dashboardLayout'
    }
  },
  {
    path: '/user/issues',
    name: 'userIssues',
    component: userIssues,
    meta: {
      layout: 'dashboardLayout'
    }
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;