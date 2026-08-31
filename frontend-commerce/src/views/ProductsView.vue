<template>
    <div style="padding: 20px; font-family: sans-serif;">
        <div style="display: flex; justify-content: space-between; align-items: center;">
            <h2>Katalog Produk Go-Commerce</h2>
            <button @click="handleLogout" style="padding: 5px 10px; cursor: pointer; color: red;">Logout</button>
        </div>

        <p v-if="loading">Memuat produk...</p>
        <p v-if="errorMsg" style="color: red;">{{ errorMsg }}</p>

        <!-- Grid Produk -->
        <div style="display: flex; gap: 15px; flex-wrap: wrap; margin-top: 20px;">
            <div v-for="product in products" :key="product.id"
                style="border: 1px solid #ccc; padding: 15px; border-radius: 8px; width: 200px; box-shadow: 2px 2px 5px rgba(0,0,0,0.1);">
                <!-- Ubah Name menjadi name -->
                <h3 style="margin-top: 0;">{{ product.name }}</h3>

                <!-- Ubah CategoryID menjadi category_id -->
                <p style="color: gray; font-size: 14px;">Kategori ID: {{ product.category_id }}</p>

                <!-- (Opsional) Tambahkan harga karena muncul di log -->
                <p style="color: #42b883; font-weight: bold; margin-bottom: 15px;">Harga: Rp {{ product.price }}</p>

                <button
                    style="width: 100%; padding: 8px; cursor: pointer; background: #42b883; color: white; border: none; border-radius: 4px;">
                    Tambah ke Keranjang
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '../api';

const products = ref([]);
const loading = ref(true);
const errorMsg = ref('');
const router = useRouter();

const fetchProducts = async () => {
    try {
        const response = await api.get('/products');
        // Menangkap array data produk
        products.value = response.data.data || response.data;
        // TAMBAHKAN BARIS INI UNTUK MENGINTIP DATA ASLI:
        console.log("Struktur JSON dari Golang:", products.value);
    } catch (error) {
        console.error("Gagal mengambil data:", error);
        errorMsg.value = "Gagal memuat produk. Sesi mungkin telah berakhir.";

        // Auto-logout jika token ditolak (401 Unauthorized)
        if (error.response && error.response.status === 401) {
            handleLogout();
        }
    } finally {
        loading.value = false;
    }
};

const handleLogout = () => {
    localStorage.removeItem('token');
    router.push('/');
};

// Panggil fungsi segera setelah halaman dirender
onMounted(() => {
    fetchProducts();
});
</script>