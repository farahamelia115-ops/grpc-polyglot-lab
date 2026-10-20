# Hands-on: Komunikasi Inter-service dengan gRPC (Polyglot)

## 1. Deskripsi
Praktikum ini mengimplementasikan komunikasi antar-service menggunakan gRPC dan Protocol Buffers (Protobuf). Proyek menggunakan Go sebagai server, Python sebagai client pertama, dan Node.js sebagai client kedua.

## 2. Struktur Proyek
```text
grpc-polyglot-lab/
├── user.proto
├── README.md
├── .gitignore
├── go-server/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── pb/
├── python-client/
│   ├── client.py
│   ├── user_pb2.py
│   ├── user_pb2_grpc.py
│   └── .venv/
└── node-client/
    ├── client.js
    └── package.json
```

## 3. Implementasi
- **Go Server:** Menyediakan layanan `GetUserProfile` pada port `50051`.
- **Python Client:** Memanggil layanan gRPC menggunakan kode hasil generate dari Protobuf.
- **Node.js Client:** Memanggil layanan gRPC menggunakan `@grpc/grpc-js` dan `@grpc/proto-loader`.
- **Protocol Buffers:** Mendefinisikan `UserRequest`, `UserResponse`, dan `UserService` pada file `user.proto`.

## 4. Hasil Pengujian
Kedua client berhasil menerima respons dari Go Server dengan data berikut:

- Nama: Budi Santoso
- Email: budi@universitas.ac.id
- Telepon: 081234567890

## 5. Analisis Static Code Generation dan Dynamic Runtime Loading

### Static Code Generation
Static code generation menghasilkan kode sumber dari file `.proto` sebelum aplikasi dijalankan. Pada proyek ini, Python menggunakan `grpc_tools.protoc` untuk menghasilkan file `user_pb2.py` dan `user_pb2_grpc.py`. Go juga menggunakan plugin Protobuf untuk menghasilkan kode di folder `go-server/pb/`.

Kelebihannya adalah struktur pesan dan service tersedia sebagai kode yang dapat diperiksa sebelum aplikasi dijalankan. Kekurangannya, kode perlu dihasilkan ulang ketika definisi `.proto` berubah.

### Dynamic Runtime Loading
Dynamic runtime loading memuat definisi `.proto` ketika aplikasi berjalan. Pada proyek ini, Node.js menggunakan `@grpc/proto-loader` untuk membaca file `user.proto` secara langsung saat client dijalankan.

Kelebihannya adalah tidak perlu membuat file kode client hasil generate secara terpisah. Kekurangannya, file `.proto` harus tersedia dan dapat dibaca ketika aplikasi dijalankan.

### Kesimpulan
Static code generation cocok ketika ingin menghasilkan kode bertipe dari definisi Protobuf sebelum runtime. Dynamic runtime loading memberikan cara yang lebih fleksibel untuk memuat definisi service saat runtime. Keduanya dapat digunakan untuk berkomunikasi dengan Go Server selama definisi service dan pesan yang digunakan sesuai.

## 6. Cara Menjalankan

### Go Server
```powershell
cd go-server
go run main.go
```

### Python Client
```powershell
cd python-client
.\.venv\Scripts\python.exe client.py
```

### Node.js Client
```powershell
cd node-client
node client.js
```

Jalankan Go Server terlebih dahulu dan biarkan terminalnya tetap terbuka sebelum menjalankan client.
