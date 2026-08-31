import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import LoginView from './views/LoginView.vue'
import ProductsView from './views/ProductsView.vue'

// Definisikan rute halaman
const routes = [
    { path: '/', component: LoginView },
    { path: '/products', component: ProductsView }
]

// Inisialisasi router
const router = createRouter({
    history: createWebHistory(),
    routes
})

// Pasang router ke dalam aplikasi Vue
createApp(App).use(router).mount('#app')