# Gator CLI
My solution to [https://www.boot.dev/courses/build-blog-aggregator-golang](https://www.boot.dev/courses/build-blog-aggregator-golang)


A command-line tool built in Go that interacts with a PostgreSQL database to manage users, feeds, and follow functionality. This tool provides commands for authentication, feed management, and aggregation.

---

## 📦 Prerequisites

Before using Gator, ensure the following are installed on your system:

- **Go** (v1.20+ recommended): [https://golang.org/dl/](https://golang.org/dl/)
- **PostgreSQL**: [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
- **Goose** (Optional) (DB migration tool) — [https://github.com/pressly/goose](https://github.com/pressly/goose)

Go Dependancies
- require (
    github.com/google/uuid v1.6.0  // indirect
    github.com/lib/pq v1.10.9      // PostgreSQL driver
)

---

## 🚀 Installation
Install the Gator CLI using the `go install` command:

```bash
go install github.com/your-username/gator@late
```
## ⚙️ Configuration

Gator expects a configuration file named .gatorconfig.json located either in your home directory or as a fallback in the current project root.
✅ Sample .gatorconfig.json:

{
  "DB_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "Current_user_name": "b"
}

Place this file at one of the following locations:

    $HOME/config.json

    ./config.json (project root)

Database Gator is required and can be setup by running make up with goose installed or manually running sql queries found in the schema folders

| Command     | Description                               | Auth Required |
| ----------- | ----------------------------------------- | ------------- |
| `login`     | Log in to the CLI using your credentials  | ❌             |
| `register`  | Register a new user                       | ❌             |
| `reset`     | Reset the database or state               | ❌             |
| `users`     | List all registered users                 | ❌             |
| `agg`       | Run the feed aggregator                   | ❌             |
| `addfeed`   | Add a new feed                            | ✅             |
| `feeds`     | List all available feeds                  | ❌             |
| `follow`    | Follow a feed                             | ✅             |
| `following` | List feeds the user is following          | ✅             |
| `unfollow`  | Unfollow a feed                           | ✅             |
| `browse`    | View aggregated posts from followed feeds | ✅             |
