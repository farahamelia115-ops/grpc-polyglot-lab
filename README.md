# Hands-on: Komunikasi Inter-service dengan gRPC (Polyglot)

## 1. Deskripsi Praktikum

Praktikum ini membahas implementasi komunikasi inter-service menggunakan gRPC (Google Remote Procedure Call) dengan pendekatan polyglot, yaitu penggunaan beberapa bahasa pemrograman dalam satu sistem.

Pada proyek ini, Go digunakan untuk membangun server gRPC, sedangkan Python dan Node.js digunakan sebagai client. Komunikasi dilakukan menggunakan Protocol Buffers (Protobuf) sebagai format definisi pesan dan service.

Server menyediakan layanan `GetUserProfile` yang menerima permintaan berdasarkan ID pengguna dan mengembalikan data profil berupa nama, email, serta nomor telepon.

## 2. Tujuan Praktikum

1. Memahami konsep komunikasi antar-service menggunakan gRPC.
2. Memahami penggunaan Protocol Buffers untuk mendefinisikan pesan dan service.
3. Mengimplementasikan server gRPC menggunakan bahasa Go.
4. Mengimplementasikan client gRPC menggunakan Python dan Node.js.
5. Memahami perbedaan static code generation dan dynamic runtime loading.
6. Menguji komunikasi antara server dan client yang menggunakan bahasa pemrograman berbeda.

## 3. Teknologi yang Digunakan

| Teknologi | Fungsi |
|---|---|
| Go | Implementasi server gRPC |
| Python | Implementasi client gRPC |
| Node.js | Implementasi client gRPC |
| gRPC | Komunikasi antar-service |
| Protocol Buffers | Definisi pesan dan service |
| `protoc` | Compiler untuk file `.proto` |
| `grpcio` | Library gRPC untuk Python |
| `grpcio-tools` | Pembuatan kode Protobuf Python |
| `@grpc/grpc-js` | Library gRPC untuk Node.js |
| `@grpc/proto-loader` | Memuat definisi Protobuf saat runtime |
| Git dan GitHub | Pengelolaan versi dan penyimpanan proyek |

## 4. Struktur Proyek

```text
grpc-polyglot-lab/
├── .gitignore
├── README.md
├── user.proto
├── go-server/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── pb/
│       ├── user.pb.go
│       └── user_grpc.pb.go
├── python-client/
│   ├── client.py
│   ├── user_pb2.py
│   ├── user_pb2_grpc.py
│   └── .venv/  (lingkungan virtual lokal, tidak diunggah)
└── node-client/
    ├── client.js
    ├── package.json
    └── package-lock.json
```

## 5. Implementasi Sistem

### 5.1 Protocol Buffers

File `user.proto` digunakan untuk mendefinisikan struktur permintaan, respons, dan service yang digunakan oleh server maupun client.

Pesan yang digunakan:

- `UserRequest`: menyimpan ID pengguna yang diminta.
- `UserResponse`: menyimpan nama, email, dan nomor telepon pengguna.
- `UserService`: mendefinisikan layanan `GetUserProfile`.

Field pada `UserResponse` terdiri atas:
- `name`: nama pengguna.
- `email`: alamat email pengguna.
- `phone`: nomor telepon pengguna.

### 5.2 Go Server

Go Server menyediakan service `GetUserProfile` dan berjalan pada port `50051`.

Server memeriksa ID pengguna yang diterima. Jika ID bernilai `1`, server mengembalikan data pengguna Budi Santoso. Untuk ID lain, server mengembalikan respons pengguna tidak ditemukan.

### 5.3 Python Client

Python Client menggunakan library `grpcio` serta file hasil generate dari Protobuf, yaitu `user_pb2.py` dan `user_pb2_grpc.py`.

Client mengirim permintaan dengan `user_id = 1` ke server, menerima respons, kemudian menampilkan nama, email, dan nomor telepon pengguna.

### 5.4 Node.js Client

Node.js Client menggunakan library `@grpc/grpc-js` dan `@grpc/proto-loader`.

File `user.proto` dimuat saat aplikasi berjalan. Client memanggil service `GetUserProfile` pada Go Server dan menampilkan data profil dari respons yang diterima.

## 6. Persiapan Lingkungan

Pastikan perangkat telah memiliki perangkat lunak berikut:

- Go.
- Python.
- Node.js dan npm.
- Protocol Buffers Compiler (`protoc`).
- Git.

Pastikan juga koneksi lokal pada port `50051` dapat digunakan oleh server gRPC.

## 7. Langkah Instalasi dan Menjalankan Program

Semua perintah berikut dijalankan dari terminal PowerShell di folder proyek utama, kecuali jika disebutkan lain.

### 7.1 Menjalankan Go Server

Buka terminal pertama:

```powershell
cd go-server
go run main.go
```

Jika berhasil, server akan menampilkan pesan bahwa gRPC Server berjalan pada port `50051`.

Biarkan terminal server tetap terbuka.

### 7.2 Menjalankan Python Client

Buka terminal kedua dari folder utama proyek:

```powershell
cd python-client
.\.venv\Scripts\python.exe client.py
```

Client akan mengirim permintaan ke Go Server dan menampilkan respons profil pengguna.

Jika lingkungan virtual belum dibuat atau library belum terpasang, siapkan lingkungan Python dan pasang dependensi gRPC terlebih dahulu.

### 7.3 Menjalankan Node.js Client

Buka terminal ketiga dari folder utama proyek:

```powershell
cd node-client
npm install
node client.js
```

Perintah `npm install` digunakan untuk memasang dependensi Node.js berdasarkan `package.json`.

## 8. Hasil Pengujian

Pengujian dilakukan dengan menjalankan Go Server terlebih dahulu, kemudian menjalankan Python Client dan Node.js Client.

### 8.1 Hasil Python Client

```text
=== Profil Pengguna ===
Nama  : Budi Santoso
Email : budi@universitas.ac.id
Telepon: 081234567890
```

### 8.2 Hasil Node.js Client

```text
=== Profil Pengguna ===
Nama  : Budi Santoso
Email : budi@universitas.ac.id
Telepon: 081234567890
```

Hasil pengujian menunjukkan bahwa kedua client berhasil menerima data profil pengguna dari Go Server.

## 9. Analisis Static Code Generation dan Dynamic Runtime Loading

### 9.1 Static Code Generation

Static code generation merupakan proses menghasilkan kode sumber dari definisi Protobuf sebelum aplikasi dijalankan.

Pada proyek ini, Python menggunakan `grpc_tools.protoc` untuk menghasilkan `user_pb2.py` dan `user_pb2_grpc.py`. Go juga menggunakan plugin Protobuf untuk menghasilkan kode di dalam folder `go-server/pb/`.

Kelebihan:
- Kode hasil generate dapat digunakan langsung oleh aplikasi.
- Struktur pesan dan service tersedia dalam bentuk kode.
- Kesalahan terkait definisi pesan tertentu dapat diketahui ketika kode dibuat atau dikompilasi.

Kekurangan:
- Kode perlu dihasilkan ulang ketika definisi `.proto` berubah.
- Proses generate membutuhkan compiler dan plugin yang sesuai.

### 9.2 Dynamic Runtime Loading

Dynamic runtime loading merupakan pendekatan memuat definisi Protobuf ketika aplikasi dijalankan.

Pada proyek ini, Node.js menggunakan `@grpc/proto-loader` untuk membaca file `user.proto` saat runtime.

Kelebihan:
- Tidak perlu menghasilkan file kode client Protobuf secara terpisah.
- Definisi service dapat dimuat langsung dari file `.proto`.

Kekurangan:
- File `.proto` harus tersedia dan dapat dibaca saat aplikasi berjalan.
- Kesalahan definisi atau pemuatan file dapat muncul ketika runtime.

### 9.3 Perbandingan

| Aspek | Static Code Generation | Dynamic Runtime Loading |
|---|---|---|
| Waktu pemrosesan definisi | Sebelum aplikasi berjalan | Saat aplikasi berjalan |
| Contoh pada proyek | Go dan Python | Node.js |
| File hasil generate | Digunakan | Tidak diperlukan untuk pemuatan client ini |
| Kebutuhan file `.proto` saat runtime | Bergantung pada implementasi | Dibutuhkan oleh client ini |

### 9.4 Kesimpulan Analisis

Static code generation sesuai ketika kode Protobuf ingin dihasilkan sebelum aplikasi dijalankan. Dynamic runtime loading memungkinkan definisi service dimuat saat runtime tanpa menghasilkan kode client secara terpisah.

Kedua pendekatan dapat digunakan untuk berkomunikasi dengan Go Server selama definisi pesan dan service sesuai.

## 10. Kesimpulan Praktikum

Praktikum ini berhasil mengimplementasikan komunikasi inter-service menggunakan gRPC dengan Go sebagai server serta Python dan Node.js sebagai client.

Definisi pesan dan service dibuat menggunakan Protocol Buffers pada file `user.proto`. Hasil pengujian menunjukkan bahwa kedua client berhasil memanggil layanan `GetUserProfile` dan menerima data nama, email, serta nomor telepon pengguna.

Praktikum ini juga menunjukkan perbedaan pendekatan static code generation pada Go dan Python dengan dynamic runtime loading pada Node.js.

## 11. Repository GitHub

Source code dan dokumentasi proyek tersedia pada repository berikut:

https://github.com/farahamelia115-ops/grpc-polyglot-lab
