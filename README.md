# golang-grpc-api

A simple gRPC API built with Golang for efficient microservice communication.

## 🚀 Features
- Uses **gRPC** for efficient microservice communication.
- Implements **service and message definitions** in Protocol Buffers.
- High-performance and optimized for large-scale applications.

## 📂 Project Structure
```
golang-grpc-api/
│── src/
│   ├── server.go  # Golang gRPC server
│   ├── proto/item.proto  # gRPC Protocol Buffers definition
│── go.mod  # Go module file
│── README.md
```

## 🛠 Installation & Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/tktarun03/golang-grpc-api.git
   cd golang-grpc-api/src
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Generate gRPC code (if needed):
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/item.proto
   ```

4. Start the gRPC server:
   ```bash
   go run server.go
   ```

5. Use a gRPC client to send requests to `localhost:50051`.

## 👨‍💻 About the Author

🚀 Created by [Arunkumar Thamilarasu](https://github.com/tktarun03) | UI Technical Architect | Golang & gRPC Expert

## ⭐ Contribute & Support
- Fork & Star this repository! ⭐
- Submit Issues and PRs to improve the gRPC API.

---
🎯 **Follow me on GitHub**: [@tktarun03](https://github.com/tktarun03)
