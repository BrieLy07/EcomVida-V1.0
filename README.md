# 🛒 ECOMVIDA – eCommerce Platform (Microservices-Based)

## 📌 Overview

**ECOMVIDA** is a modular eCommerce platform designed with a **microservices architecture**, enabling high scalability, independent deployments, and a clear separation of concerns across six business domains.

---

## 🎯 Objectives

- Design a distributed and decoupled platform.
- Enable independent deployment of each domain.
- Provide a unified frontend experience.
- Apply DevOps best practices and secure architecture.
- Use optimal database technologies for each domain.

---

## 🧱 General Architecture

- **Frontend:** React (modular interface)
- **Backend:** Microservices in Node.js, Python, Go, Rust, Kotlin
- **Cloud Infrastructure:** AWS (EC2, RDS, S3, Redis, DynamoDB, Elasticsearch)
- **Orchestration:** Docker + GitHub Actions (CI/CD)
- **Security:** Bastion Host, IAM, WAF
- **Ingress:** Cloudflare + API Gateway
- **Load Balancing:** 1 Load Balancer per domain

---

## 📦 Domains & Databases

| Domain                  | Microservices | Databases Used                              |
|-------------------------|---------------|----------------------------------------------|
| Users & Auth            | 5             | PostgreSQL, MongoDB, MySQL, DynamoDB         |
| Products                | 5             | PostgreSQL, MongoDB, Redis, S3               |
| Orders & Payments       | 5             | PostgreSQL, Redis, MongoDB, DynamoDB         |
| Inventory & Logistics   | 5             | MySQL, PostgreSQL, MongoDB, Elasticsearch    |
| User Experience         | 5             | MongoDB, Redis, DynamoDB, Elasticsearch      |
| Admin & Audit           | 5             | PostgreSQL, Redis, DynamoDB                  |

---

## ⚙️ DevOps Workflow

- **GitHub Repositories** with `main` and `test` branches
- **CI/CD Pipelines:** GitHub Actions with `deploy.yml`
- **Dockerized Microservices** (`Dockerfile`, `docker-compose.yml`)
- **Secrets Handling:** via GitHub Secrets (.env)
- **Testing:** Unit tests and Postman collections
- **Deployment:** EC2 instances per domain with Auto Scaling Groups and Load Balancers

---

## 🛡️ Disaster Recovery Plan

- **RTO:** 1–3 hours | **RPO:** up to 24 hours
- **Backups:** Daily (RDS, MongoDB, DynamoDB)
- **Mitigation Strategies:** Auto Scaling, snapshots, WAF, MFA, firewalls
- **Response:** SNS alerts + automatic redeploy with GitHub Actions

---

## 💵 Annual Cost Estimate (AWS)

| Resource                     | Quantity | Monthly Cost | Annual Total |
|------------------------------|----------|---------------|---------------|
| EC2 t3.micro (6 domains)     | 6        | $9            | $648          |
| RDS PostgreSQL / MySQL       | 4        | $15           | $720          |
| DynamoDB + MongoDB Atlas     | 2        | $10           | $240          |
| S3 (20 GB)                   | 1        | $2            | $24           |
| Redis (Elasticache)          | 2 nodes  | $15           | $360          |
| Outbound Traffic (600 GB/y)  | -        | $0.09/GB      | $54           |
| CloudWatch (Basic)           | -        | $5            | $60           |
| **Estimated Annual Total**   | -        | -             | **≈ $2,106**  |

> This estimate assumes Free Tier-optimized resources. Costs may drop with Spot Instances or service sharing.

---

## 📊 System Diagrams

- Use Case Diagrams per domain
- State Diagrams: Product & Support Ticket
- DevOps Flow: GitHub Actions → EC2
- Infrastructure Flow: Cloudflare → Bastion → API Gateway → Load Balancer → Domains

> See `/docs/diagrams/*.puml` for PlantUML source files

---

## ✅ Project Status

- ✅ Microservice design completed
- ✅ Dockerization and database scripts ready
- ✅ AWS infrastructure architecture defined
- 🔄 Deployment in progress (per domain)
- 🔜 React frontend integration pending

---

## 📌 Recommendations

- Add centralized monitoring (CloudWatch, Prometheus)
- JWT authentication and role-based access control
- Daily backup strategy + replication per domain
- Observability stack (ELK or OpenTelemetry)

---