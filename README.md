# Voidspace

**Voidspace** Social is a cloud-first mini-microservices application. It is a mini Twitter-like social media app where users can create posts, follow others, send direct messages, and view their timeline. 

### Architecture & Tech Stack

-   **BFF (Backend for Frontend)**: The API Gateway acts as an orchestrator and aggregator, exposing a clean REST API while communicating with various internal microservices via gRPC.
-   **Distributed Workflows**: Uses **Temporal** for reliable service orchestration and complex long-running operations (e.g., account deletion workflows with compensation logic).
-   **Deployment**: 
    -   **Backend**: Deployed on **Google Cloud Run** via GitHub Actions.
    -   **Frontend**: Next.js application deployed on **Vercel**.
    -   **Persistence**: PostgreSQL databases managed on **Neon**.

---

[![Architecture of voidspace](/docs/img/diagram.png)](/docs/img/diagram.png)

---

## API Documentation

- [POSTMAN Collection v2](/docs/api/voidspace_v2_refined.postman_collection.json)
- [Live API Docs](https://voidspace-gateway-591941627936.asia-southeast2.run.app/docs) (OUTDATED)

---

## Services Overview

| Service         | Responsibilities                                       | Stack / Communication  |
| --------------- | ------------------------------------------------------ | ---------------------- |
| **Users**       | Authentication, user profiles, follow/unfollow         | gRPC (Go + PostgreSQL) |
| **Posts**       | Create posts, like posts, timeline feed                | gRPC (Go + PostgreSQL) |
| **Comments**    | Commenting on posts                                    | gRPC (Go + PostgreSQL) |
| **API Gateway** | **BFF** (Orchestrator & Aggregator)                    | Echo (Go, REST → gRPC) |

---

