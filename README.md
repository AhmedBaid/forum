# 🗨️ Forum 

## 📚 Description

This project is a full-stack web forum application built with Go, HTML/CSS, and SQLite. It allows users to communicate through posts and comments, associate posts with categories, interact via likes/dislikes, and filter posts. It includes user authentication with session management using cookies. The backend uses the SQLite database and is containerized using Docker.

---

## 🚀 Features

### 👥 Authentication
- User registration with:
  - **Email**
  - **Username**
  - **Password** (encrypted using `bcrypt` - _bonus_)
- Login and session creation using **cookies**.
- Session expiration and management (only one session per user).
- Checks for unique email on registration.
- Secure password validation.

### 💬 Communication
- Registered users can:
  - Create posts
  - Comment on posts
- Posts and comments are visible to all users (even unregistered).
- Posts can have one or more **categories** (subforums).

### 👍 Likes & Dislikes
- Registered users can like/dislike **posts** and **comments**.
- Total like/dislike count is visible to all users.

### 🔎 Filtering
- Filter posts by:
  - **Categories** (public)
  - **User’s created posts** (logged-in only)
  - **User’s liked posts** (logged-in only)

### 🗃️ SQLite Database
- Stores:
  - Users
  - Posts
  - Comments
  - Categories
  - Likes/Dislikes
  - Sessions
- Includes at least:
  - One `CREATE` query
  - One `INSERT` query
  - One `SELECT` query
- Entity-relationship diagram recommended for schema planning.

### 🐳 Docker
- Fully containerized application using **Docker**.
- Easy setup and deployment.

---

## 🛠️ Tech Stack

- **Backend**: Go
- **Database**: SQLite3
- **Auth & Encryption**: `bcrypt`, `uuid`
- **Session Management**: Cookies with expiration
- **Containerization**: Docker

---

## 📁 Project Structure

