# 🚀 AI Code Reviewer – Project Plan

## 1. Vision

Build a **production-grade AI Code Review Bot in Go** that:

* Reviews GitHub Pull Requests automatically
* Provides line-level comments
* Suggests performance, security, and Go best practices
* Generates tests
* Streams AI feedback
* Works with OpenAI / Ollama (pluggable)

Target audience:
Teams who want automated, consistent, and intelligent PR reviews.

---

## 2. Goals

### Functional

* GitHub App integration with webhooks
* Parse PR diff and context
* AI analysis per file
* Post comments on PR
* Summarize review
* Generate unit tests suggestions
* Security scan via AI

### Non-Functional

* Clean architecture
* Testable core
* Streaming responses
* Retry + rate limiting
* Observability
* Pluggable LLM
* Docker ready

---

## 3. Architecture

### High Level Flow

```
GitHub Webhook
      ↓
Webhook Handler (Go)
      ↓
Diff Parser
      ↓
Chunker + Context Builder
      ↓
AI Provider (OpenAI / Ollama)
      ↓
Review Engine
      ↓
GitHub Comments
      ↓
Storage + Metrics
```

---

## 4. Tech Stack

* Go 1.22+
* Gin / Fiber
* PostgreSQL
* Redis (queue)
* OpenAI / Ollama
* GitHub App
* Docker
* Prometheus

---

## 5. Folder Structure

```
ai-code-reviewer/
├── cmd/server
├── internal
│   ├── app
│   ├── domain
│   ├── github
│   ├── ai
│   ├── diff
│   ├── review
│   ├── worker
│   ├── storage
│   ├── vector
│   └── observability
├── pkg
├── migrations
└── docker
```

---

## 6. Milestones

### Phase 1 – Foundation

* [ ] Project skeleton
* [ ] GitHub App setup
* [ ] Webhook receiver
* [ ] Config management
* [ ] Logger & metrics

### Phase 2 – Core Logic

* [ ] Diff parser
* [ ] Chunker
* [ ] Prompt templates
* [ ] LLM interface
* [ ] OpenAI implementation

### Phase 3 – Review Engine

* [ ] Analyzer service
* [ ] Go best practices rules
* [ ] Security prompts
* [ ] Test generator

### Phase 4 – Integration

* [ ] Post PR comments
* [ ] Summary comment
* [ ] Streaming support

### Phase 5 – Production

* [ ] Worker queue
* [ ] Retry
* [ ] Rate limit
* [ ] Docker
* [ ] Tests

---

## 7. Core Domains

### Review

* File review
* Issues
* Suggestions
* Score

### Diff

* Parse
* Chunk
* Token limit

### AI

* Provider
* Streaming
* Prompts

### GitHub

* Webhook
* Comments
* PR client

---

## 8. API Endpoints

```
POST /webhook/github
GET  /health
GET  /metrics
```

---

## 9. Prompt Strategy

* Code quality
* Performance
* Security
* Idiomatic Go
* Test coverage

---

## 10. Risks

* Token limits
* Large diffs
* Rate limits
* Hallucinations

Mitigation:

* Chunking
* Validation
* Deterministic prompts

---

## 11. Definition of Done

* Bot reviews real PR
* Posts comments
* Streams response
* Dockerized
* Tests passing
* README ready

---

## 12. Future

* Multi-language
* Vector memory
* Custom policies
* Slack integration
* Dashboard
