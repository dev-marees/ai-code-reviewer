# 🤖 AI Code Reviewer (Go)

An **enterprise-grade AI Code Review Bot** built with Go that automatically reviews GitHub Pull Requests, provides line-level suggestions, security insights, performance tips, and generates test recommendations.

---

## 🎯 Goal of This Project

Modern teams spend huge time reviewing pull requests.
Human reviews are:

* inconsistent
* slow
* depend on individual experience
* often miss security & performance issues

### This project aims to:

✅ Automate first-level code review
✅ Enforce Go best practices
✅ Detect security & performance problems
✅ Suggest unit tests
✅ Provide learning feedback for developers
✅ Reduce reviewer workload

---

## 🚀 What It Does

When a Pull Request is opened:

1. GitHub sends a webhook
2. Service fetches PR diff
3. Diff is parsed & chunked
4. AI analyzes each file
5. Bot posts:

   * line comments
   * summary review
   * test suggestions
   * security notes

---

## 🧠 Features

### Core

* GitHub App integration
* PR diff parsing
* Context-aware chunking
* Streaming AI responses
* Line-level comments
* Summary generation

### AI Capabilities

* Go idiomatic review
* Performance improvements
* Security scanning
* Test generation
* Code smell detection

### Engineering

* Clean architecture
* Pluggable LLM (OpenAI / Ollama)
* Async worker
* Retry & rate limit
* Observability
* Docker ready

---

## 🏗 Architecture

```
GitHub Webhook
      ↓
Webhook Handler (Go)
      ↓
Diff Parser → Chunker
      ↓
AI Provider
      ↓
Review Engine
      ↓
GitHub Comments
      ↓
Storage + Metrics
```

### Design Principles

* Interface-driven design
* Provider abstraction
* Streaming first
* Testability
* Separation of concerns

---

## 📁 Project Structure

```
ai-code-reviewer/
├── cmd/server        → entrypoint
├── internal
│   ├── app           → http server
│   ├── domain        → core models
│   ├── github        → github client
│   ├── ai            → llm providers
│   ├── diff          → patch parser
│   ├── review        → analysis engine
│   ├── worker        → async jobs
│   ├── storage       → db layer
│   └── observability → logs & metrics
```

---

## 🛠 Tech Stack

* **Go 1.22+**
* GitHub App
* OpenAI / Ollama
* PostgreSQL
* Redis Queue
* Docker
* Prometheus

---

## ⚙️ Setup

### 1. Clone

```bash
git clone <repo>
cd ai-code-reviewer
```

### 2. Env

Copy:

```
.env.example → .env
```

Add:

* GitHub App credentials
* OpenAI key

### 3. Run

```bash
go run cmd/server/main.go
```

### 4. Expose (local dev)

Use ngrok for webhook testing.

---

## 🔍 Review Criteria

The AI reviews code based on:

### 1. Correctness

* Bugs
* Edge cases
* Nil checks

### 2. Go Best Practices

* Error handling
* Naming
* Concurrency
* Context usage

### 3. Performance

* Allocations
* Goroutines
* Complexity

### 4. Security

* Injection
* Secrets
* Validation

### 5. Tests

* Coverage
* Table tests
* Mocks

---

## 🧩 Example Output

**Comment on line 42**

> Avoid creating a goroutine inside a loop without capturing variable.
> Suggestion: pass value as parameter.

**Summary**

* 2 security issues
* 3 performance tips
* Test cases missing

---

## 🧪 Development Plan

* Day 1 – Bootstrap
* Day 2 – Webhooks
* Day 3 – GitHub client
* Day 4 – Diff parser
* Day 5 – Chunker
* Day 6 – LLM abstraction
* Day 7 – OpenAI
* Day 8 – Prompts
* Day 9 – Engine
* Day 10 – Comments
* Day 11 – Worker
* Day 12 – Observability
* Day 13 – Tests
* Day 14 – Docker

---

## 🎯 Who Is This For?

* Go teams
* Solo developers
* Open source maintainers
* Startups without big review bandwidth

---

## 📌 Roadmap

* Vector memory
* Multi-language
* Slack alerts
* Dashboard
* Custom policies

---

## 🤝 Contributing

PRs welcome!
Follow clean architecture & tests.

---

## 🧑‍💻 Author

Built as part of **#BuildInPublic – AI with Go** journey.
