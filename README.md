# Go-Commerce: Fullstack E-Commerce Monorepo

Go-Commerce adalah aplikasi web *e-commerce* fungsional yang dibangun menggunakan arsitektur modern pemisahan *frontend* dan *backend* di dalam satu repositori tunggal (*monorepo*). Proyek ini mendemonstrasikan penguasaan siklus pengembangan perangkat lunak secara utuh (*end-to-end*), mulai dari manajemen antarmuka reaktif, autentikasi sesi, hingga manipulasi database relasional dan API.

## 🛠️ Teknologi yang Digunakan

**Frontend:**
* Vue.js 3 (Composition API & `<script setup>`)
* Vite (Build Tool)
* Vue Router (Navigasi Klien)
* Axios (HTTP Client)

**Backend:**
* Golang (Go)
* Gin Web Framework (RESTful API & Routing)
* GORM (Object Relational Mapping)
* JWT (JSON Web Tokens untuk Autentikasi)
* CORS Middleware

**Database:**
* MySQL

## ✨ Fitur Utama

* **Sistem Autentikasi Keamanan:** Registrasi dan Login menggunakan token JWT.
* **Katalog Produk Dinamis:** Pengambilan data produk secara asinkron dari REST API.
* **Keranjang Belanja Reaktif:** Antarmuka *Drawer* (laci luncung) modern yang memperbarui *state* dan kalkulasi harga secara *real-time* tanpa memuat ulang halaman.
* **Manajemen Transaksi (Checkout):** Integrasi relasi database (*Preload/JOIN*) untuk memproses perpindahan data keranjang dan pembersihan sesi transaksi pengguna secara otomatis.

## 📂 Struktur Direktori

Repositori ini menggunakan arsitektur *monorepo* simetris:

```text
go-commerce/
├── backend-api/         # Logika server, model database, dan routing REST API (Golang)
├── frontend-commerce/   # Antarmuka pengguna reaktif dan state management (Vue 3)
└── README.md

```

## 🚀 Panduan Instalasi dan Menjalankan Proyek

Pastikan Anda telah menginstal **Go**, **Node.js**, dan **MySQL** di sistem Anda.

### 1. Kloning Repositori

```bash
git clone [https://github.com/firdhausranggaa/go-commerce.git](https://github.com/firdhausranggaa/go-commerce.git)
cd go-commerce

```

### 2. Setup Database (MySQL)

* Buka MySQL (melalui DBeaver, phpMyAdmin, atau CLI).
* Buat database baru bernama `ecomm`.
* *(Tabel akan secara otomatis dibuat oleh fitur AutoMigrate GORM saat server Golang dijalankan).*

### 3. Menjalankan Backend (Golang)

Buka terminal dan arahkan ke folder backend:

```bash
cd backend-api
go mod tidy
go run main.go

```

*Server akan berjalan di `http://localhost:5000*`

### 4. Menjalankan Frontend (Vue 3)

Buka tab terminal baru dan arahkan ke folder frontend:

```bash
cd frontend-commerce
npm install
npm run dev

```

*Aplikasi web dapat diakses melalui `http://localhost:5173*`