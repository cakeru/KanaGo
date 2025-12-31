# Kanago: Japanese Learning & Word Tracker

Kanago is a full-stack Japanese learning and word tracking application. It helps users master kana, kanji, and vocabulary, and allows them to track and review words and sentences they are learning. The project is designed as a showcase of modern Go backend and React frontend development.

---

## ✨ Features

- **Static Kana & Kanji Library:** Fast, static serving of all kana and kanji characters.
- **User Vocabulary & Sentence Tracker:** Users can add, edit, and review their own vocabulary and sentences.
- **Authentication:** Secure registration, login, and JWT-based authentication.
- **Progress Tracking:** Track learning stats and review history.
- **RESTful API:** Well-documented endpoints for all features.
- **Modern Frontend:** Responsive React UI for learning, reviewing, and tracking.
- **Extensible:** Easily add new features or content.

---

## 🗂️ Project Structure

```
/backend         # Go API server, database, migrations
/frontend        # React app (user interface)
Japanese/        # Static data: kana, kanji, vocab (JSON/TS files)
dev-doc/         # Developer documentation and guides
README.md        # This file
```

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/cakeru/kanago.git
cd kanago
```

### 2. Backend Setup

- **Requirements:** Go 1.20+, PostgreSQL

```bash
cd backend
cp .env.example .env      # Edit DB credentials as needed
go mod tidy
go run ./cmd/server/main.go
```

- **Database:**
  - Run the SQL migrations in `dev-doc/go-backend-setup-guide.md` to create tables.
  - Seed kana/kanji/vocab tables using scripts or import from `Japanese/`.

### 3. Frontend Setup

- **Requirements:** Node.js 18+, npm

```bash
cd frontend
npm install
npm start
```

- The React app will run on [http://localhost:3000](http://localhost:3000) by default.

---

## 📚 Documentation

- **API Reference:** [`dev-doc/api-specification.md`](dev-doc/api-specification.md)
- **Backend Setup:** [`dev-doc/go-backend-setup-guide.md`](dev-doc/go-backend-setup-guide.md)
- **Frontend Setup:** [`dev-doc/react-frontend-setup-guide.md`](dev-doc/react-frontend-setup-guide.md)
- **Architecture:** [`dev-doc/VISUAL-ARCHITECTURE.md`](dev-doc/VISUAL-ARCHITECTURE.md)
- **Implementation Checklist:** [`dev-doc/IMPLEMENTATION-CHECKLIST.md`](dev-doc/IMPLEMENTATION-CHECKLIST.md)

---

## 🏗️ Key Technologies

- **Backend:** Go, Gin, PostgreSQL, JWT, GORM/SQLX
- **Frontend:** React, TypeScript, Material UI/Chakra UI, Fetch/Axios
- **DevOps:** Docker (optional), .env config, REST API

---

## 📝 Example API Endpoints

- `POST /api/register` — Register a new user
- `POST /api/login` — Login and get JWT
- `GET /api/kana` — Get all kana (static file)
- `GET /api/kanji` — Get all kanji (static file)
- `GET /api/vocab` — Get user’s vocabulary (requires auth)
- `POST /api/vocab` — Add a new vocabulary word (requires auth)
- `GET /api/sentences` — Get user’s sentences (requires auth)

See [`dev-doc/api-specification.md`](dev-doc/api-specification.md) for full details.

---

## 🖼️ Screenshots

<!-- Add screenshots of the UI here when available -->

---

## 🤝 Contributing

Pull requests and issues are welcome!  
See [`dev-doc/IMPLEMENTATION-CHECKLIST.md`](dev-doc/IMPLEMENTATION-CHECKLIST.md) for open tasks.

---

## 📄 License

MIT License

---

## 🙏 Acknowledgements

- Major Thanks to [KanaDojo](https://github.com/lingdojo/kana-dojo) this project is heavily inspired by this great project.
- [KanaDojo](https://github.com/lingdojo/kana-dojo) for static kana/kanji/vocab data inspiration.
- Open source Japanese datasets and the language learning community.

---

## 📬 Contact

For questions or feedback, open an issue or contact [your email/contact info].
