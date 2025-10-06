# finalproject-sb-go-70-Adrian



---

# 📝 Go Notes API — Simple Notes REST API with JWT Authentication

Proyek ini adalah **API sederhana** untuk mengelola catatan (*notes*) yang dibuat menggunakan **Golang (Go)** dengan dukungan **JWT Authentication**, **PostgreSQL**, dan dapat di-*deploy* ke **Railway**.
Kamu dapat melakukan **registrasi, login, membuat catatan, melihat daftar catatan, memperbarui, menandai favorit, dan menghapus catatan.**

---

## 🚀 Demo API (Railway)

API ini sudah bisa langsung diakses dari Railway:

```
https://finalproject-sb-go-70-adrian-production.up.railway.app
```

Contoh endpoint:

* **Register:** `POST /api/users/register`
* **Login:** `POST /api/users/login`
* **Get Notes:** `GET /api/notes`
* **Create Note:** `POST /api/notes`

---

## 📁 Struktur Proyek

```
go-notes-api/
├── auth/
│   └── jwt.go             # Logika pembuatan & validasi JWT
├── database/
│   └── db.go              # Koneksi ke database PostgreSQL
├── handler/
│   ├── user_handler.go    # Handler untuk user (register, login, profile)
│   └── note_handler.go    # Handler untuk notes (CRUD & favorit)
├── middleware/
│   └── auth.go            # Middleware untuk validasi JWT di endpoint privat
├── models/
│   ├── user.go            # Model User
│   └── note.go            # Model Note
├── migrations/            # File migrasi database
├── .env                   # Variabel lingkungan (JANGAN di-commit!)
├── dbconfig.yml
└── main.go                # Entry point aplikasi
```

---

## ⚙️ Setup Lokal

### 1. Clone Repository

```bash
git clone https://github.com/username/go-notes-api.git
cd go-notes-api
```

### 2. Setup Environment

Buat file `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=notesdb
JWT_SECRET=your_jwt_secret
```

### 3. Jalankan Aplikasi

```bash
go run main.go
```

API akan berjalan di:
👉 `http://localhost:8080`

---

## 🧪 Pengujian API Menggunakan Postman

Berikut alur pengujian yang **terstruktur dan mudah diikuti** menggunakan **Postman**.

---

### ⚡ Langkah 0: Persiapan

1. Jalankan `go run main.go`
2. Buka **Postman**
3. Pastikan API bisa diakses di `http://localhost:8080`

---

### 🌍 Langkah 1: Buat Environment di Postman

1. Klik ikon mata (👁️) → **Add Environment**
2. Tambahkan variabel:

   ```
   baseURL = http://localhost:8080
   jwt_token = (kosongkan dulu)
   ```
3. Simpan dan aktifkan environment.

---

### 👤 Langkah 2: Register User

* **Method:** `POST`
* **URL:** `{{baseURL}}/api/users/register`
* **Body (JSON):**

  ```json
  {
      "username": "adrian",
      "email": "adrian@example.com",
      "password": "password123"
  }
  ```

Respons:

```json
{
    "message": "User registered successfully"
}
```

---

### 🔐 Langkah 3: Login dan Simpan Token Otomatis

* **Method:** `POST`
* **URL:** `{{baseURL}}/api/users/login`
* **Body (JSON):**

  ```json
  {
      "email": "adrian@example.com",
      "password": "password123"
  }
  ```
* **Tests Tab:**

  ```javascript
  pm.environment.set("jwt_token", pm.response.json().token);
  ```

Setelah `Send`, token JWT akan tersimpan otomatis di Postman Environment.

---

### 📝 Langkah 4: Buat Catatan Baru

* **Method:** `POST`
* **URL:** `{{baseURL}}/api/notes`
* **Authorization:** Bearer Token → `{{jwt_token}}`
* **Body (JSON):**

  ```json
  {
      "title": "Catatan Pertamaku",
      "content": "Ini adalah isi dari catatan pertamaku di API Go."
  }
  ```

Respons:

```json
{
    "id": 1,
    "title": "Catatan Pertamaku",
    "content": "Ini adalah isi dari catatan pertamaku di API Go.",
    "favorite": false
}
```

---

### 📚 Langkah 5: Lihat Semua Catatan

* **Method:** `GET`
* **URL:** `{{baseURL}}/api/notes`
* **Authorization:** Bearer Token

---

### 🔍 Langkah 6: Lihat Catatan Spesifik

* **Method:** `GET`
* **URL:** `{{baseURL}}/api/notes/1`

---

### ✏️ Langkah 7: Update Catatan

* **Method:** `PUT`
* **URL:** `{{baseURL}}/api/notes/1`
* **Body:**

  ```json
  {
      "title": "Catatan Pertamaku (Update)",
      "content": "Isinya sudah diperbarui melalui Postman."
  }
  ```

---

### ⭐ Langkah 8: Ubah Status Favorit Catatan

* **Method:** `PUT`
* **URL:** `{{baseURL}}/api/notes/1/favorite`

Kirim sekali → akan menjadi favorit
Kirim lagi → favorit dibatalkan

---

### ❤️ Langkah 9: Lihat Semua Catatan Favorit

* **Method:** `GET`
* **URL:** `{{baseURL}}/api/notes/favorites`

---

### 👤 Langkah 10: Lihat Profil User

* **Method:** `GET`
* **URL:** `{{baseURL}}/api/users/profile`

---

### 🗑️ Langkah 11: Hapus Catatan

* **Method:** `DELETE`
* **URL:** `{{baseURL}}/api/notes/1`

---

## 🧱 Teknologi yang Digunakan

* **Go (Golang)**
* **Gin Gonic** — Web framework
* **GORM** — ORM untuk koneksi database
* **PostgreSQL**
* **JWT (JSON Web Token)** — Autentikasi
* **Railway.app** — Deployment platform

---

## 📦 Fitur Utama

✅ Register & Login (JWT Authentication)
✅ CRUD Notes
✅ Toggle Favorite Notes
✅ View Favorite Notes
✅ View User Profile
✅ Secure Endpoints via Middleware
✅ Siap untuk Deployment ke Railway

---

## 🧑‍💻 Author

**Adrian Syah Abidin**
🔗 [Railway Deployment Link](https://railway.com/project/5d4a250b-62f0-46e5-961d-fdeb868c41c7?environmentId=b2bc2b4e-e7f9-4623-b3da-33c3ef99f7ef)

---
