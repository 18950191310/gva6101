import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
    path: '/',
    redirect: '/login'
},
    {
        path: '/login',
        name: 'login',
        component: () =>
            import ('@/view/login/login.vue')
    } ,
    {
        path: '/chome',
        name: 'chome',
        component: () =>
            import ('@/view/home/chome.vue')
    },
    {
        path: '/cart',
        name: 'cart',
        component: () =>
            import ('@/view/components/cart.vue')
    },
    {
        path: '/cuser',
        name: 'cuser',
        component: () =>
            import ('@/view/components/cuser/cuser.vue')
    },
    {
        path: '/cuser-detail',
        name: 'cuser-detail',
        component: () =>
            import ('@/view/components/cuser/cuser-detail.vue')
    },
    {
        path: '/order',
        name: 'order',
        component: () =>
            import ('@/view/components/order/order.vue')
    },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router