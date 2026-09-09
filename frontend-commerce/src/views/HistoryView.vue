<template>
    <div style="padding: 20px; font-family: sans-serif; max-width: 800px; margin: 0 auto;">
        <div
            style="display: flex; justify-content: space-between; align-items: center; border-bottom: 2px solid #eee; padding-bottom: 10px;">
            <h2>Riwayat Transaksi</h2>
            <button @click="router.push('/products')"
                style="padding: 8px 15px; cursor: pointer; background: white; border: 1px solid #333; border-radius: 4px;">
                &larr; Kembali ke Katalog
            </button>
        </div>

        <p v-if="loading" style="text-align: center; color: gray;">Memuat riwayat perjalanan belanjamu...</p>
        <div v-else-if="history.length === 0" style="text-align: center; color: gray; margin-top: 50px;">
            Belum ada transaksi. Ayo mulai belanja!
        </div>

        <div v-else style="margin-top: 20px; display: flex; flex-direction: column; gap: 15px;">
            <div v-for="trx in history" :key="trx.id"
                style="border: 1px solid #ddd; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.05);">
                <div style="display: flex; justify-content: space-between; margin-bottom: 15px;">
                    <h3 style="margin: 0; color: #333;">Order ID: #{{ trx.id }}</h3>
                    <span
                        style="background: #42b883; color: white; padding: 3px 10px; border-radius: 12px; font-size: 14px;">Sukses</span>
                </div>
                <p style="font-weight: bold; font-size: 18px; margin: 0 0 15px 0;">Total Pembayaran: Rp {{ trx.amount }}
                </p>

                <div style="background: #f9f9f9; padding: 10px; border-radius: 4px;">
                    <p style="margin: 0 0 10px 0; font-size: 14px; font-weight: bold; color: #666;">Barang yang dibeli:
                    </p>
                    <ul style="margin: 0; padding-left: 20px; font-size: 14px;">
                        <li v-for="item in trx.items" :key="item.id" style="margin-bottom: 5px;">
                            {{ item.product?.name || `Produk ID ${item.product_id}` }} &times; {{ item.quantity }}
                            <span style="color: gray;">(Harga saat beli: Rp {{ item.price }})</span>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '../api';

const history = ref([]);
const loading = ref(true);
const router = useRouter();

onMounted(async () => {
    try {
        const response = await api.get('/history');
        history.value = response.data.data.reverse();
    } catch (error) {
        console.error("Gagal memuat riwayat", error);
    } finally {
        loading.value = false;
    }
});
</script>