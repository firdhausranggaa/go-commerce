import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import LoginView from './views/LoginView.vue'
import ProductsView from './views/ProductsView.vue'
import HistoryView from './views/HistoryView.vue'

const routes = [
    { path: '/', component: LoginView, meta: { requiresGuest: true } },
    { path: '/products', component: ProductsView, meta: { requiresAuth: true } },
    { path: '/history', component: HistoryView, meta: { requiresAuth: true } }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const isAuthenticated = !!localStorage.getItem('token');

    if (to.meta.requiresAuth && !isAuthenticated) {
        next('/');
    } else if (to.meta.requiresGuest && isAuthenticated) {
        next('/products');
    } else {
        next();
    }
})

createApp(App).use(router).mount('#app')