# xgo 🧭

**`xgo`** is a terminal-based CLI tool to track your progress through a 60-day Go learning roadmap — including fundamentals, web dev, microservices, DevOps, and real-time systems.

No external files needed after installation. Your progress is saved locally and can be reset anytime.

---
** Built on MANJARO-OS **

🧭 60-Day Go Learning Roadmap
📚 Phase 1: Go Fundamentals (Days 1–5)

  -- Learn Go syntax, data types, functions, arrays, and interfaces

🌐 Phase 2: Web Development with Go (Days 6–10)

  -- Build web servers, handle routing, templates, forms, and SQL

🚀 Phase 3: Advanced Web Development (Days 11–15)

  -- Middleware, auth, REST APIs, testing, and performance

🧼 Phase 4: Best Practices & Common Mistakes (Days 16–20)

  -- Go mistakes, error handling, memory, concurrency, structure

🧱 Phase 5: Microservices Concepts (Days 21–25)

  -- Learn service design, gRPC, database strategy, scaling

🛠️ Phase 6: Microservices Projects (Days 26–30)

  -- Build real microservices with discovery and deployment

📦 Phase 7: Infrastructure Essentials (Days 31–35)

  -- Docker, K8s, CI/CD, Prometheus, observability

🔒 Phase 8: Security & Observability (Days 36–40)

  -- Secure configs, JWT auth, tracing, logging, rate limiting

🔄 Phase 9: Communication Techniques (Days 41–45)

  -- RabbitMQ, Kafka, retry logic, backoff, API gateway

⚡ Phase 10: Real-Time & Serverless (Days 46–50)

  -- WebSockets, serverless (AWS Lambda), gRPC streaming

🧪 Phase 11: Project-Based Learning (Days 51–55)

   -- Build inventory, payment, chat, notification, analytics services

🚢 Phase 12: Final Polish & Deployment (Days 56–60)

  -- Helm, scaling, secrets, Grafana dashboards, final project

    


## 📦 Features

- ✅ View current step with `xgo -now`
- 🎯 Mark step as done with `xgo -done`
- 🔁 Reset progress with `xgo -reset`
- 📂 Stores progress in `~/.config/xgo/progress.json`
- 🔒 Self-contained — no external roadmap file needed after build

---

## 🚀 Installation

```bash
git clone https://github.com/Awatansh/xgo.git
cd xgo
go build -o xgo main.go
sudo mv xgo /usr/local/bin
