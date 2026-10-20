
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");

const PROTO_PATH = path.join(__dirname, "..", "user.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: false,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const proto = grpc.loadPackageDefinition(packageDefinition).user;

const client = new proto.UserService(
  "localhost:50051",
  grpc.credentials.createInsecure()
);

client.GetUserProfile({ userId: 1 }, (error, response) => {
  if (error) {
    console.error("Gagal mengambil profil:", error.message);
    client.close();
    process.exitCode = 1;
    return;
  }

  console.log("=== Profil Pengguna ===");
  console.log("Nama  :", response.name);
  console.log("Email :", response.email);
  console.log("Telepon:", response.phone);

  client.close();
});
