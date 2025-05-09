# 📚 Go REST API with Gin and MongoDB

This is a simple RESTful API built with [Golang](https://go.dev/), [Gin](https://github.com/gin-gonic/gin), and [MongoDB](https://www.mongodb.com/). It's designed as a beginner-friendly project to help you get started with building APIs in Go.

📝 **Based on this article:**  
👉 [Getting Started with REST APIs in Golang: A Practical Guide](#)

---

## 🚀 Features

- Create, Read, Update, Delete (CRUD) operations on books
- Clean project structure (models, routes, controllers)
- MongoDB connection using official Go driver
- Simple and easy to follow

---

## 📁 Project Structure

```
go-rest-api/
├── main.go
├── controllers/
│ └── book_controller.go
├── models/
│ └── book.go
├── routes/
│ └── book_routes.go
├── config/
│ └── db.go
```

---

## ⚙️ Requirements

- Go 1.20+
- MongoDB (local or Atlas)
- Git

---

## 📦 Installation

```bash
# Clone the repo
git clone https://github.com/gorogagah/go-rest-api.git
cd go-rest-api

# Initialize Go modules
go mod tidy

# Run the app
go run main.go
```

---

## 🔌 API Endpoints

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| GET    | /books      | Get all books     |
| GET    | /books/\:id | Get book by ID    |
| POST   | /books      | Create a new book |
| PUT    | /books/\:id | Update a book     |
| DELETE | /books/\:id | Delete a book     |

---

## 🛠 Sample Payload

```json
{
  "title": "The Go Programming Language",
  "author": "Alan A. A. Donovan"
}
```

---

## 🧑‍💻 Author

Made with ❤️ by [Your Name](https://github.com/gorogagah)

Follow me on [Medium](https://medium.com/@gorogagah) or [LinkedIn](https://www.linkedin.com/in/anggoro-gagah/) for more beginner Go tutorials.

---

## 📜 License
This project is open-source and available under the MIT [License](LICENSE).
