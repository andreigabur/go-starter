# 🚀 GO-STARTER Project Template

Welcome to the **GO-STARTER** project template! This repository is designed to kickstart new projects by providing a structured foundation with both backend service examples (using Go and Hono.js) and a modern frontend setup (using Astro and React).

## 📂 Project Structure

The project is divided into core backend services (`services`) and a frontend application (`web`).

| Directory | Description |
| :--- | :--- |
| `database` | Contains database-related files (e.g., schemas, migrations). |
| `proto` | Houses Protocol Buffers definitions (`.proto` files) for gRPC communication. |
| `services` | **Backend services**. |
| `web` | **Frontend applications**. |
| `Makefile` | Utility commands for common tasks (e.g., building, cleaning, running tests). |

---

## 💻 Backend Services (`/services`)

The `services` directory contains various backend implementations, each exposing a simple endpoint to return a **list of users** at the `/users` address for comparison and testing.

* **`app` (Standard Go)**
    * A standard Go application using the built-in `net/http` package.
    * Exposes a **REST API** endpoint.
    * Also includes an example of a **gRPC service**.
* **`gin` (Go with Gin Framework)**
    * A Go application leveraging the **Gin** framework for fast and robust web services.
    * Exposes the same **REST API** endpoint.
* **`honojs` (Hono.js)**
    * A service built with **Hono.js** (a lightweight web framework for the edge) running on Bun.
    * Exposes the same **REST API** endpoint.

---

## 🌐 Frontend (`/web/astro`)

This directory contains the project for the frontend application(s):

* **astro**
    * Uses **Astro** for high-performance static sites and content-driven websites.
    * Integrated with **React** for dynamic components.
    * Uses **Tailwind CSS** for utility-first styling.

---

## 🧪 Benchmarking Tests

You can run simple benchmarking tests to compare the performance of the different REST API implementations. Ensure you have the necessary environment (Go, Bun) set up before running.

### 1. **Go Standard (`/services/app`)**

These benchmarks test the performance of the standard Go HTTP service.

| Step | Command | Description |
| :--- | :--- | :--- |
| **Start Service** | `go run services/app/cmd/main.go` | Starts the service. |
| **Run Benchmark** | `go test services/app/test -run TestGoStandardRESTUsersEndpoint -v` | Executes the dedicated benchmark test for the `/users` REST endpoint. |

### 2. **Go Gin (`/services/gin`)**

These benchmarks test the performance of the service built with the Gin framework.

| Step | Command | Description |
| :--- | :--- | :--- |
| **Start Service** | `go run services/gin/main.go` | Starts the service. |
| **Run Benchmark** | `go test services/gin/test -run TestGoGinRESTUsersEndpoint -v` | Executes the dedicated benchmark test for the `/users` REST endpoint. |

### 3. **Hono.js (`/services/honojs`)**

These benchmarks test the performance of the Hono.js service, which requires the **Bun runtime**.

| Step | Command | Description |
| :--- | :--- | :--- |
| **Start Service** | `bun run --cwd services/honojs dev` | Starts the Hono.js service using Bun. |
| **Run Benchmark** | `bun run --cwd services/honojs test:benchmark` | Executes the benchmark script defined in the `package.json` for the REST endpoint. |

> **Note:** Remember to **stop a running service** before starting the next one if they use the same port, or adjust the port configurations as needed.