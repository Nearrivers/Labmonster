import { createRouter, createWebHashHistory } from 'vue-router';

export const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@/views/HomeView.vue'),
      children: [
        {
          path: '',
          name: 'NotOpened',
          component: () => import('@/components/NoFileOpened.vue'),
        },
        {
          path: 'flowchart/:path',
          name: 'flowchart',
          component: () => import('@/components/FlowChart.vue'),
        },
      ],
    },
  ],
});
