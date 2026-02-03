# 📅 14-Day Build Plan – AI Code Reviewer in Go

Goal:
Ship a real, working GitHub AI reviewer in 14 days with daily commits and LinkedIn content.

---

## 🔹 Day 1 – Project Bootstrap

**Tasks**

* Create repo & folder structure
* Setup Go module
* Logger & config

**Files**

* cmd/server/main.go
* internal/config/config.go
* internal/observability/logger.go

**Commit**

* “chore: bootstrap project structure”

**LinkedIn Post**

* Project intro + architecture diagram

**Outcome**

* Service starts with /health

---

## 🔹 Day 2 – GitHub App + Webhook

**Tasks**

* Create GitHub App
* Verify webhook signature
* Basic handler

**Files**

* internal/github/webhook.go
* internal/app/routes.go

**Commit**

* “feat: github webhook receiver”

**Outcome**

* Receive PR events locally

---

## 🔹 Day 3 – GitHub Client

**Tasks**

* Fetch PR details
* Get diff
* Auth with installation token

**Files**

* internal/github/client.go
* internal/github/pr.go

**Commit**

* “feat: github client and diff fetch”

**Outcome**

* Print diff on event

---

## 🔹 Day 4 – Diff Parser

**Tasks**

* Parse patch format
* Split by file
* Extract hunks

**Files**

* internal/diff/parser.go
* internal/diff/models.go

**Commit**

* “feat: diff parser with hunk support”

**Outcome**

* Structured diff object

---

## 🔹 Day 5 – Chunking Engine

**Tasks**

* Token estimation
* Smart chunking
* Context builder

**Files**

* internal/diff/chunk.go

**Commit**

* “feat: context aware chunker”

**Outcome**

* LLM-ready chunks

---

## 🔹 Day 6 – LLM Abstraction

**Tasks**

* Define interface
* Streaming types

**Files**

* internal/ai/llm.go
* internal/ai/stream.go

**Commit**

* “feat: llm abstraction”

**Outcome**

* Provider-agnostic design

---

## 🔹 Day 7 – OpenAI Provider

**Tasks**

* Implement client
* Streaming
* Retry

**Files**

* internal/ai/openai.go

**Commit**

* “feat: openai streaming client”

**Outcome**

* Send prompt → receive tokens

---

## 🔹 Day 8 – Prompt Engineering

**Tasks**

* Code review prompt
* Go best practices
* Security prompt

**Files**

* internal/ai/prompts/code_review.tmpl

**Commit**

* “feat: review prompts v1”

**Outcome**

* Meaningful AI response

---

## 🔹 Day 9 – Review Engine

**Tasks**

* Analyzer service
* Map AI → domain

**Files**

* internal/review/service.go
* internal/domain/review.go

**Commit**

* “feat: review engine”

**Outcome**

* Structured review result

---

## 🔹 Day 10 – Post Comments to GitHub

**Tasks**

* Line comments
* Summary comment

**Files**

* internal/github/comment.go

**Commit**

* “feat: post pr comments”

**Outcome**

* Bot comments on PR 🎉

---

## 🔹 Day 11 – Worker Queue

**Tasks**

* Async processing
* Retry with backoff

**Files**

* internal/worker/processor.go

**Commit**

* “feat: async worker”

---

## 🔹 Day 12 – Observability

**Tasks**

* Metrics
* Tracing
* Logs

**Files**

* internal/observability/metrics.go

**Commit**

* “feat: observability”

---

## 🔹 Day 13 – Tests

**Tasks**

* Table tests
* Mocks

**Commit**

* “test: add unit tests”

---

## 🔹 Day 14 – Docker + Polish

**Tasks**

* Dockerfile
* README
* Demo

**Commit**

* “chore: docker & docs”

---

# 📌 Daily LinkedIn Template

**Post Structure**

1. Problem
2. What I built today
3. Code snippet
4. Demo
5. Tomorrow plan

---

# 📦 Definition of Done

* Bot reviews real PR
* Line comments
* Summary
* Docker
* Tests
* README

---

# 🔮 After 14 Days

* Add vector memory
* Multi-language
* Slack
* Dashboard
