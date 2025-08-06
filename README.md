# 🗺️ GeoGraphDB - Hierarchical Geographical Graph Backend

GeoGraphDB is a scalable, modular backend system for managing multi-level geographical graphs. It supports nested regional graphs (e.g., a country made of states made of districts), integrates with PostgreSQL for persistent storage, Redis for caching, and uses the GraphHopper API for real-time edge routing information.

---

## ✨ Features

- ⚙️ **Graph CRUD APIs** — Add/Delete Nodes and Edges.
- 🧠 **Recursive Graph Structure** — Nodes can contain sub-graphs (hierarchical modeling).
- 🚀 **GraphHopper Integration** — Auto-generates edge metadata (distance, time, etc.) from real-world data.
- ⚡ **Redis Caching** — Avoid redundant GraphHopper API calls by caching edge lookups.
- 🛢️ **PostgreSQL Persistence** — Durable node and edge storage.
- 🔌 **REST API (Gin)** — Clean modular API structure using `internal/api`, `internal/db`, `internal/graph`.

---

## 📁 Project Structure

```

geographdb/
│
├── cmd/
│   └── main.go              # Entry point
│
├── internal/
│   ├── api/                 # HTTP handlers and routes
│   ├── db/                  # PostgreSQL and Redis interaction
│   ├── graph/               # Graph objects, types, and logic
│   ├── gh/                  # GraphHopper API wrapper
│   └── config/              # Config loading
│
├── go.mod / go.sum          # Dependencies
├── .env                     # Environment variables
└── README.md

````

---

## 🛠️ Getting Started

### ✅ Prerequisites

- Go ≥ 1.20
- PostgreSQL
- Redis
- GraphHopper API key

### 🔧 Installation

```bash
git clone https://github.com/yourusername/geographdb.git
cd geographdb
go mod tidy
````

### ⚙️ Environment Variables

Create a `.env` file at the root:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_password
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=geograph

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

GRAPHHOPPER_API_KEY=your_api_key
```


## 🧪 API Endpoints

### ➕ Create Node

```
POST /nodes
{
  "id": "node-123",
  "name": "Bangalore",
  "latitude": 12.9716,
  "longitude": 77.5946,
  "parent_id": "region-1"
}
```

### ➖ Delete Node

```
DELETE /nodes/:id
```

### ➕ Create Edge (Auto fetches metadata from GraphHopper if not cached)

```
POST /edges
{
  "from_node_id": "node-1",
  "to_node_id": "node-2"
}
```

### 🔍 Get Edge

```
GET /edges/:from/:to
```

---

## 🧩 Technologies Used

* [Go](https://golang.org/)
* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [pgx PostgreSQL Driver](https://github.com/jackc/pgx)
* [go-redis](https://github.com/redis/go-redis)
* [GraphHopper API](https://graphhopper.com)

---
