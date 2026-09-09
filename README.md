# Go-Commerce: Fullstack E-Commerce Monorepo

Go-Commerce adalah aplikasi web *e-commerce* tingkat produksi (*production-grade*) yang dibangun menggunakan arsitektur *monorepo*. Proyek ini mendemonstrasikan implementasi siklus pengembangan *fullstack*, mulai dari antarmuka reaktif, autentikasi ganda, hingga pemrosesan transaksi database yang mematuhi standar ACID.

## 🛠️ Tumpukan Teknologi (Tech Stack)

**Frontend:**
* Vue.js 3 (Composition API) & Vite
* Vue Router (Navigation Guards)
* Axios (HTTP Client & Interceptors)

**Backend:**
* Golang (Go) & Gin Web Framework
* GORM (Object Relational Mapping)
* JWT (JSON Web Tokens) & Bcrypt (Password Hashing)
* CORS Middleware

**Database:**
* MySQL

## ✨ Fitur Unggulan

* **Transaksi ACID & Snapshot Harga:** Logika *checkout* di backend menggunakan *Database Transaction* untuk mencegah kebocoran data, serta merekam harga *snapshot* agar riwayat pesanan kebal dari perubahan harga produk di masa depan.
* **Sistem Autentikasi & Keamanan:** Fitur Registrasi dan Login yang dilindungi oleh enkripsi sandi Bcrypt dan sesi berbasis JWT. Akses peramban dikunci secara ketat menggunakan *Vue Navigation Guards*.
* **Keranjang Belanja Reaktif:** Antarmuka laci (*Drawer*) reaktif dengan deteksi duplikasi barang otomatis (*Upsert*) dan sistem umpan balik *Toast Notifications*.
* **Jejak Riwayat Belanja:** Visualisasi data relasional kompleks (*Nested Preload* GORM) dari tabel entitas transaksi, rincian item, dan katalog produk.

## 📂 Struktur Direktori

```text
go-commerce/
├── backend-api/         # Logika server, model database, proteksi route (Golang)
├── frontend-commerce/   # Antarmuka klien reaktif, state management (Vue 3)
└── README.md

```

## 🚀 Panduan Instalasi dan Menjalankan Proyek

Pastikan **Go**, **Node.js**, dan **MySQL** telah terinstal di sistem Anda.

### 1. Setup Database (MySQL)

* Buat database kosong bernama `ecomm`.
* *(Seluruh tabel akan dicetak secara otomatis oleh AutoMigrate GORM)*.

### 2. Menjalankan Backend (Golang)

Buka terminal dan arahkan ke folder backend:

```bash
cd backend-api
go mod tidy
go run main.go

```

*Catatan: Anda dapat mengaktifkan sementara `seeders.Seed(db)` di `main.go` pada saat eksekusi pertama untuk menyuntikkan data katalog produk dummy.*

### 3. Menjalankan Frontend (Vue 3)

Buka tab terminal baru dan arahkan ke folder frontend:

```bash
cd frontend-commerce
npm install
npm run dev

```

*Aplikasi web dapat diakses melalui `http://localhost:5173*`
