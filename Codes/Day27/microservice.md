```markdown
# Microservices in Go 
---
##  What are Microservices?

Microservices = breaking a large application into **small independent services**

Each service:
- Has **one responsibility**
- Runs independently
- Communicates with other services

---

##  Service Boundaries

Service boundaries define **what each service does**

Rule:
- One service = one business responsibility  
- Should be independently deployable  

### Good Example
- Auth Service → authentication  
- User Service → user data  
- Notification Service → emails/SMS  

### Bad Example
- One service handling everything (auth + users + notifications)

---

##  Inter-Service Communication

Services need to talk to each other

### 1. Synchronous Communication
- Direct request-response  
- Uses HTTP (REST) or gRPC  
- Simple but tightly coupled  

### 2. Asynchronous Communication
- Uses message brokers  
- Services communicate via events  
- More scalable and resilient  

### Trade-off
- Sync → simple but can fail if one service is down  
- Async → complex but more reliable  

---

##  Service Discovery

Problem:
- Services run on different machines (dynamic IPs)

Solution:
- Use a way to **find services dynamically**

### Approaches
- Client-side discovery  
- Server-side discovery (load balancer)  
- DNS-based discovery  

Benefit:
- No hardcoding of service addresses  
- Easy scaling  

---

##  Config & Secrets

### Configuration
Non-sensitive data:
- Ports  
- URLs  
- Feature flags  

Stored in:
- Environment variables  
- Config files  

---

### Secrets
Sensitive data:
- Passwords  
- API keys  
- Tokens  

Best practices:
- Never hardcode  
- Use secure storage  
- Load at runtime  

---

##  Observability Basics

Observability = understanding system behavior

### 3 Main Parts

#### 1. Logging
- Records events  
- Helps debugging  

#### 2. Metrics
- Numerical data  
- Example: request count, latency  

#### 3. Tracing
- Tracks request across services  
- Helps find bottlenecks  

---

##  Rate Limiting

Rate limiting controls **how many requests are allowed**

Why?
- Prevent overload  
- Protect system  
- Ensure fair usage  

---

##  What is a Burst?

A **burst** = many requests coming **at the same time**

Example:
- Normal → 10 requests/sec  
- Burst → 100 requests in 1 second  

Bursts can overload systems if not controlled  

---

#  Rate Limiting Approaches
---

## 1. Token Bucket

Idea:
You need a **token to enter**

- Tokens are added slowly over time  
- Each request uses 1 token  
- No token → no entry  

Simple understanding:
- You can **save tokens and use later**  
- Allows short bursts  

✔ Good for: APIs where bursts are okay  

---

## 2. Leaky Bucket

Idea:
Requests go into a **queue and come out slowly**

- Input can be fast  
- Output is always steady  

Simple understanding:
- Like water leaking at constant speed  
- No sudden spikes  

✔ Good for: smooth traffic systems  

---

## 3. Fixed Window

Idea:
Count requests in a **time block**

- Example: 100 requests per minute  
- After time ends → reset  

Simple understanding:
- Easy but can be unfair  

⚠ Problem:
- Burst possible at window boundaries  

---

## 4. Sliding Window

Idea:
Check requests in the **last X seconds continuously**

Simple understanding:
- Always looking at recent requests  
- More fair  

✔ Good for: accurate control  

---

##  Quick Comparison

- Token Bucket → flexible, allows bursts  
- Leaky Bucket → smooth, no bursts  
- Fixed Window → simple but inaccurate  
- Sliding Window → fair and accurate  

---

##  When to Use

- Most systems → Token Bucket  
- Need smooth flow → Leaky Bucket  
- Simple use case → Fixed Window  
- Need accuracy → Sliding Window  

---
Rate limiting:
- Token Bucket = save and spend tokens  
- Leaky Bucket = constant flow  
- Fixed Window = count per time block  
- Sliding Window = count recent requests  
---
```
