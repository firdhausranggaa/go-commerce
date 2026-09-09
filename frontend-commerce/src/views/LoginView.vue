<template>
    <div
        style="max-width: 400px; margin: 100px auto; padding: 30px; border: 1px solid #ccc; border-radius: 8px; text-align: center; font-family: sans-serif; box-shadow: 0 4px 10px rgba(0,0,0,0.1);">
        <h2>{{ isRegister ? 'Daftar Akun Baru' : 'Login Go-Commerce' }}</h2>

        <form @submit.prevent="handleSubmit"
            style="display: flex; flex-direction: column; gap: 15px; margin-top: 20px;">
            <input v-model="username" type="text" placeholder="Username" required
                style="padding: 10px; border: 1px solid #ccc; border-radius: 4px;" />

            <input v-if="isRegister" v-model="email" type="email" placeholder="Email" required
                style="padding: 10px; border: 1px solid #ccc; border-radius: 4px;" />

            <input v-model="password" type="password" placeholder="Password" required
                style="padding: 10px; border: 1px solid #ccc; border-radius: 4px;" />

            <button type="submit"
                style="padding: 12px; background: #333; color: white; border: none; border-radius: 4px; cursor: pointer; font-weight: bold;">
                {{ isRegister ? 'Daftar Sekarang' : 'Masuk' }}
            </button>
        </form>

        <p v-if="errorMsg" style="color: red; margin-top: 15px;">{{ errorMsg }}</p>
        <p v-if="successMsg" style="color: green; margin-top: 15px;">{{ successMsg }}</p>

        <p style="margin-top: 20px; font-size: 14px; color: #666;">
            {{ isRegister ? 'Sudah punya akun?' : 'Belum punya akun?' }}
            <span @click="toggleMode"
                style="color: #42b883; cursor: pointer; font-weight: bold; text-decoration: underline;">
                {{ isRegister ? 'Login di sini' : 'Daftar di sini' }}
            </span>
        </p>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../api';

const isRegister = ref(false);
const username = ref('');
const email = ref('');
const password = ref('');
const errorMsg = ref('');
const successMsg = ref('');
const router = useRouter();

const toggleMode = () => {
    isRegister.value = !isRegister.value;
    errorMsg.value = '';
    successMsg.value = '';
};

const handleSubmit = async () => {
    errorMsg.value = '';
    successMsg.value = '';

    try {
        if (isRegister.value) {
            const response = await api.post('/register', {
                username: username.value,
                email: email.value,
                password: password.value
            });
            successMsg.value = response.data.message || "Registrasi berhasil! Silakan login.";
            isRegister.value = false;
            password.value = '';
        } else {
            const response = await api.post('/login', {
                username: username.value,
                password: password.value
            });
            localStorage.setItem('token', response.data.token);
            router.push('/products');
        }
    } catch (error) {
        errorMsg.value = error.response?.data?.error || error.response?.data?.message || "Kredensial salah atau terjadi kesalahan sistem.";
    }
};
</script>