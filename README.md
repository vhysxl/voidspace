# Voidspace

**Voidspace** Social is a cloud-first mini-microservices application. It is a mini Twitter-like social media app where users can create posts, follow others, send direct messages, and view their timeline. This project is deployed using Google Cloud products (Cloud Run, Artifact Registry, Cloud Storage), Railway (MySQL and PostgreSQL), and GitHub Actions for CI/CD automation. The application runs entirely on Cloud Run.

---

[![Architecture of voidspace](/docs/img/diagram.png)](/docs/img/diagram.png)

---

## API Documentation

- [API Docs](http://localhost:5000/docs)

> **Disclaimer**: If you encounter **500/503 errors**, it may be due to Railway databases going into sleep mode. Please wait a few moments and try again.

---

## Services Overview

| Service         | Responsibilities                                       | Stack / Communication  |
| --------------- | ------------------------------------------------------ | ---------------------- |
| **Users**       | Authentication, user profiles, follow/unfollow         | gRPC (Go + MySQL)      |
| **Posts**       | Create posts, like posts, timeline feed                | gRPC (Go + PostgreSQL) |
| **Comments**    | Commenting on posts                                    | gRPC (Go + MySQL)      |
| **API Gateway** | Orchestrator & aggregator; exposes REST API to clients | Echo (Go, REST → gRPC) |

---

## To Do

- [ ] Finish UI implementation
- [ ] Add Redis caching
- [ ] Refactor to full microservices pattern (Saga / 2PC)
- [ ] Add tracing and monitoring (jaeger, prometheus, grafana, otel)
