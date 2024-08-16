import { Routes } from '@/types/Routes';
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
          name: Routes.NotOpened,
          component: () => import('@/components/NoFileOpened.vue'),
        },
        {
          path: 'flowchart/:path',
          name: 'flowchart',
          component: () => import('@/components/FlowChart.vue'),
        },
        {
          path: 'image/:path',
          name: Routes.Image,
          component: () => import('@/components/AppImage.vue'),
        },
        {
          path: 'video/:path',
          name: Routes.Video,
          component: () => import('@/components/AppVideo.vue'),
        },
      ],
    },
  ],
});
