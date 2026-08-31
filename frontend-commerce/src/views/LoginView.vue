<template>
    <div class="login-container">
        <h2>Login Go-Commerce</h2>
        <form @submit.prevent="handleLogin">
            <input v-model="username" type="text" placeholder="Username" required />
            <input v-model="password" type="password" placeholder="Password" required />
            <button type="submit">Masuk</button>
        </form>
        <p v-if="errorMsg" style="color: red;">{{ errorMsg }}</p>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../api';

const username = ref('');
const password = ref('');
const errorMsg = ref('');
const router = useRouter();

const handleLogin = async () => {
    try {
        const response = await api.post('/login', {
            username: username.value,
            password: password.value
        });
        localStorage.setItem('token', response.data.token);
        router.push('/products');
    } catch (error) {
        errorMsg.value = "Username atau password salah!";
    }
};
</script>