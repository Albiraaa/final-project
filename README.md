# 📡 Inventaris Barang Radio Komunikasi (HT)

Project ini adalah **REST API Inventaris HT (Handy Talky)** menggunakan **Golang**, **JWT Authentication**, **PostgreSQL**, dan **GORM Auto Migration**. Project ini dibuat sebagai **Final Project Bootcamp Golang** dan sudah siap dijalankan secara **local maupun deploy di Railway**.

---

## 🚀 Fitur Utama

- Autentikasi User (Login) menggunakan **JWT**
- CRUD **Kategori**
- CRUD **Lokasi**
- CRUD **Inventaris HT**
- Relasi data (HT → Kategori & Lokasi)
- Auto migration database (GORM)
- Bisa diakses via **Postman** dan **Web Client (HTML)**

---

## 🛠️ Teknologi yang Digunakan

- Golang
- Gin Framework
- GORM
- PostgreSQL
- JWT (JSON Web Token)
- Railway (Deployment)

---


## ▶️ Menjalankan Project

Server berjalan di:
```
final-project-production-2cff.up.railway.app
```
Sudah terdeploy di Railway
---

## 🔐 Autentikasi (Login)

### Endpoint
```
POST /api/users/login
```

### Body (JSON)
```json
{
  "email": "admin@mail.com",
  "password": "admin"
}
```

### Response
```json
{
  "token": "JWT_TOKEN_DISINI"
}
```

> 🔥 **Token WAJIB** dikirim pada setiap request selanjutnya melalui header:
```
Authorization: Bearer <token>
```

---

# 📦 API ENDPOINTS

## 📁 Kategori

| Method | Endpoint | Keterangan |
|------|--------|------------|
| GET | /api/kategori | Ambil semua kategori |
| GET | /api/kategori/:id | Ambil kategori by ID |
| POST | /api/kategori | Tambah kategori |
| PUT | /api/kategori/:id | Update kategori |
| DELETE | /api/kategori/:id | Hapus kategori |

### Body POST / PUT
```json
{
  "nama": "Radio HT"
}
```

---

## 📍 Lokasi

| Method | Endpoint | Keterangan |
|------|--------|------------|
| GET | /api/lokasi | Ambil semua lokasi |
| GET | /api/lokasi/:id | Ambil lokasi by ID |
| POST | /api/lokasi | Tambah lokasi |
| PUT | /api/lokasi/:id | Update lokasi |
| DELETE | /api/lokasi/:id | Hapus lokasi |

### Body POST / PUT
```json
{
  "nama": "Gudang Utama"
}
```

---

## 📻 Inventaris HT

| Method | Endpoint | Keterangan |
|------|--------|------------|
| GET | /api/ht | Ambil semua HT |
| GET | /api/ht/:id | Ambil HT by ID |
| POST | /api/ht | Tambah HT |
| PUT | /api/ht/:id | Update HT |
| DELETE | /api/ht/:id | Hapus HT |

### Body POST / PUT
```json
{
  "serial": "HT-001",
  "merk": "Motorola",
  "status": "Aktif",
  "kategori_id": 1,
  "lokasi_id": 1
}
```

---

# 🧪 Cara Menggunakan API

## ✅ Cara 1: Menggunakan Postman

1. Login ke endpoint `/api/users/login`
2. Copy **JWT Token** dari response
3. Masukkan ke Header Postman:
   ```
   Authorization: Bearer <token>
   ```
4. Gunakan endpoint GET / POST / PUT / DELETE sesuai kebutuhan

> 🔹 POST & PUT → **Body → raw → JSON**

---

## 🌐 Cara 2: Menggunakan Web Client (HTML)

Project ini dilengkapi **client HTML** untuk testing API.

### Cara Pakai:

1. Buka file HTML di browser
2. Isi **Base URL**:
   ```
   https://final-project-production-2cff.up.railway.app
   ```
3. Login menggunakan email & password
4. Token otomatis tersimpan
5. Gunakan:
   - Dropdown list (HT / Kategori / Lokasi)
   - Form POST / PUT / DELETE
   - Data otomatis terisi saat memilih ID

> 💡 Cocok untuk demo tanpa Postman

---

## 🗄️ Database & Migration

- Database menggunakan **PostgreSQL**
- Tabel dibuat otomatis oleh **GORM Auto Migration**
- Tidak perlu manual create table

Tabel utama:
- users
- kategori
- lokasi
- ht

---

## 👤 Author

Nama: **Aldo Bintang Rhamadhan**  
Project: Final Project Bootcamp Golang

---

## ✅ Status Project

✔ CRUD lengkap
✔ JWT Auth
✔ Auto Migration
✔ Railway Deploy
✔ Web Client

🚀 **Project siap dipresentasikan & dinilai**

