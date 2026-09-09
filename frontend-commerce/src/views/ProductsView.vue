<template>
    <div style="padding: 20px; font-family: sans-serif;">

        <div v-if="toastMsg"
            style="position: fixed; top: 20px; left: 50%; transform: translateX(-50%); background: #333; color: white; padding: 12px 24px; border-radius: 30px; z-index: 9999; box-shadow: 0 4px 12px rgba(0,0,0,0.2); transition: all 0.3s ease;">
            {{ toastMsg }}
        </div>

        <div style="display: flex; justify-content: space-between; align-items: center;">
            <h2>Katalog Produk Go-Commerce</h2>
            <div>
                <button @click="router.push('/history')"
                    style="padding: 5px 15px; cursor: pointer; margin-right: 10px; background: white; color: #333; border: 1px solid #333; border-radius: 4px;">
                    Riwayat Pesanan
                </button>

                <button @click="openCart"
                    style="padding: 5px 15px; cursor: pointer; margin-right: 10px; background: #333; color: white; border: none; border-radius: 4px;">
                    Lihat Keranjang
                </button>
                <button @click="handleLogout"
                    style="padding: 5px 10px; cursor: pointer; color: red; background: none; border: 1px solid red; border-radius: 4px;">
                    Logout
                </button>
            </div>
        </div>

        <p v-if="loading">Memuat produk...</p>
        <p v-if="errorMsg" style="color: red;">{{ errorMsg }}</p>

        <div style="display: flex; gap: 15px; flex-wrap: wrap; margin-top: 20px;">
            <div v-for="product in products" :key="product.id"
                style="border: 1px solid #ccc; padding: 15px; border-radius: 8px; width: 200px; box-shadow: 2px 2px 5px rgba(0,0,0,0.1);">
                <h3 style="margin-top: 0;">{{ product.name }}</h3>
                <p style="color: gray; font-size: 14px;">Kategori ID: {{ product.category_id }}</p>
                <p style="color: #42b883; font-weight: bold; margin-bottom: 15px;">Harga: Rp {{ product.price }}</p>
                <button @click="addToCart(product.id)"
                    style="width: 100%; padding: 8px; cursor: pointer; background: #42b883; color: white; border: none; border-radius: 4px;">
                    Tambah ke Keranjang
                </button>
            </div>
        </div>

        <div v-if="isCartOpen"
            style="position: fixed; top: 0; left: 0; width: 100vw; height: 100vh; background: rgba(0,0,0,0.5); display: flex; justify-content: flex-end; z-index: 1000;">
            <div
                style="width: 350px; background: white; height: 100%; padding: 20px; box-shadow: -2px 0 10px rgba(0,0,0,0.2); overflow-y: auto;">

                <div
                    style="display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #ccc; padding-bottom: 10px; margin-bottom: 20px;">
                    <h3 style="margin: 0;">Keranjang Belanja</h3>
                    <button @click="isCartOpen = false"
                        style="background: none; border: none; font-size: 24px; cursor: pointer; color: #888;">&times;</button>
                </div>

                <p v-if="cartLoading">Mengambil data keranjang...</p>
                <div v-else-if="cartItems.length === 0" style="color: gray; text-align: center; margin-top: 50px;">
                    Keranjang kamu masih kosong.
                </div>

                <div v-else>
                    <div v-for="item in cartItems" :key="item.id"
                        style="border-bottom: 1px solid #eee; padding-bottom: 10px; margin-bottom: 10px;">
                        <h4 style="margin: 0 0 5px 0;">{{ item.product.name }}</h4>
                        <div style="display: flex; justify-content: space-between; font-size: 14px;">
                            <span>Jumlah: {{ item.quantity }}</span>
                            <span style="font-weight: bold; color: #42b883;">Rp {{ item.product.price * item.quantity
                            }}</span>
                        </div>
                    </div>

                    <div style="margin-top: 20px; text-align: right; font-weight: bold; font-size: 18px;">
                        Total: Rp {{ cartTotal }}
                    </div>
                    <button @click="processCheckout"
                        style="width: 100%; padding: 12px; margin-top: 20px; cursor: pointer; background: #333; color: white; border: none; border-radius: 4px; font-weight: bold;">
                        Checkout Sekarang
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import api from '../api';

const products = ref([]);
const loading = ref(true);
const errorMsg = ref('');
const router = useRouter();

const isCartOpen = ref(false);
const cartItems = ref([]);
const cartLoading = ref(false);

const toastMsg = ref('');
const showToast = (msg) => {
    toastMsg.value = msg;
    setTimeout(() => { toastMsg.value = ''; }, 3000);
};

const fetchProducts = async () => {
    try {
        const response = await api.get('/products');
        products.value = response.data.data || response.data;
    } catch (error) {
        console.error("Gagal mengambil data:", error);
        errorMsg.value = "Gagal memuat produk. Sesi mungkin telah berakhir.";
        if (error.response && error.response.status === 401) handleLogout();
    } finally {
        loading.value = false;
    }
};

const addToCart = async (productId) => {
    try {
        const response = await api.post('/cart', { product_id: productId });
        showToast(response.data.message);
        fetchCart();
    } catch (error) {
        showToast('Gagal menambahkan ke keranjang.');
    }
};

const fetchCart = async () => {
    cartLoading.value = true;
    try {
        const response = await api.get('/cart');
        cartItems.value = response.data.data || [];
    } catch (error) {
        console.error("Gagal mengambil keranjang:", error);
    } finally {
        cartLoading.value = false;
    }
};

const openCart = () => {
    isCartOpen.value = true;
    fetchCart();
};

const cartTotal = computed(() => {
    return cartItems.value.reduce((total, item) => {
        if (item.product) {
            return total + (item.product.price * item.quantity);
        }
        return total;
    }, 0);
});

const processCheckout = async () => {
    if (cartItems.value.length === 0) {
        showToast("Pilih barang dulu sebelum checkout!");
        return;
    }
    try {
        const response = await api.post('/checkout');
        showToast(response.data.message);
        fetchCart();
        isCartOpen.value = false;
    } catch (error) {
        showToast(error.response?.data?.error || "Gagal memproses pembayaran");
    }
};

const handleLogout = () => {
    localStorage.removeItem('token');
    router.push('/');
};

onMounted(() => {
    fetchProducts();
});
</script>