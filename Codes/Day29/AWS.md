

# 1. Amazon EC2 (Elastic Compute Cloud)

**Service Model:** IaaS  

## What it actually is
A virtual machine running on AWS hardware using virtualization (hypervisor). You get full control over OS and software.

## Purpose
Run applications with complete control over infrastructure

## Real Usage
- Backend APIs
- Monolithic applications
- Custom software deployments

## Key Points
- Full OS access (Linux/Windows)
- Manual patching and updates
- Attach storage (EBS)
- Works with Auto Scaling

## Subtopics
- AMI (Amazon Machine Image)
- Instance types (compute vs memory optimized)
- Auto Scaling Groups
- Spot vs Reserved instances

---

# 2. AWS Lambda

**Service Model:** PaaS (Serverless / FaaS)

## What it actually is
Executes your code inside short-lived containers triggered by events.

## Purpose
Run code without managing servers

## Real Usage
- API backends (with API Gateway)
- File processing (S3 triggers)
- Automation tasks

## Key Points
- No server management
- Event-driven execution
- Stateless

## Subtopics
- Cold starts
- Event sources
- Memory & timeout tuning

---

# 3. Amazon ECS (Elastic Container Service)

**Service Model:** PaaS  

## What it actually is
A container orchestration system to run Docker containers on managed clusters.

## Purpose
Run containerized applications

## Real Usage
- Microservices
- Backend systems

## Key Points
- AWS-managed scheduler
- Works with EC2 or Fargate
- Easier than Kubernetes

## Subtopics
- Task definitions
- Services vs tasks
- Fargate launch type

---

# 4. Amazon EKS (Elastic Kubernetes Service)

**Service Model:** PaaS  

## What it actually is
Managed Kubernetes control plane where AWS manages master nodes.

## Purpose
Run Kubernetes workloads

## Real Usage
- Large-scale distributed systems
- Multi-cloud apps

## Key Points
- Full Kubernetes ecosystem
- Complex but powerful

## Subtopics
- Pods, deployments
- Helm charts
- Ingress controllers

---

# 5. Amazon S3 (Simple Storage Service)

**Service Model:** PaaS  

## What it actually is
Object storage where data is stored as objects (file + metadata + key).

## Purpose
Store files at massive scale

## Real Usage
- Images, videos
- Backups
- Static websites

## Key Points
- Extremely durable (11 9’s)
- Unlimited storage
- Accessible via HTTP

## Subtopics
- Storage classes
- Lifecycle rules
- Versioning

---

# 6. Amazon EBS (Elastic Block Store)

**Service Model:** IaaS  

## What it actually is
A virtual hard disk attached to EC2 instances.

## Purpose
Provide persistent storage

## Real Usage
- OS disk
- Database storage

## Key Points
- Low latency
- Persistent data

## Subtopics
- Volume types (gp3, io1)
- Snapshots
- IOPS tuning

---

# 7. Amazon RDS (Relational Database Service)

**Service Model:** PaaS  

## What it actually is
Managed relational database service.

## Purpose
Run SQL databases without managing infrastructure

## Real Usage
- Web applications
- Transaction systems

## Key Points
- Automated backups
- Multi-AZ availability

## Subtopics
- Read replicas
- Failover
- Parameter tuning

---

# 8. Amazon DynamoDB

**Service Model:** PaaS  

## What it actually is
Distributed NoSQL key-value database.

## Purpose
High-performance and scalable database

## Real Usage
- Real-time apps
- Gaming systems
- IoT

## Key Points
- Millisecond latency
- Auto scaling

## Subtopics
- Partition keys
- Secondary indexes
- Capacity modes

---

# 9. Amazon VPC (Virtual Private Cloud)

**Service Model:** IaaS  

## What it actually is
A logically isolated virtual network in AWS.

## Purpose
Secure network environment

## Real Usage
- Private backend systems
- Secure infrastructure

## Key Points
- Full control over IP ranges
- Subnet isolation

## Subtopics
- Route tables
- NAT Gateway
- Internet Gateway

---

# 10. Elastic Load Balancer (ELB)

**Service Model:** PaaS  

## What it actually is
Distributes incoming traffic across multiple resources.

## Purpose
Improve availability and scalability

## Real Usage
- High traffic systems

## Key Points
- Health checks
- Traffic routing

## Subtopics
- ALB vs NLB
- SSL termination
- Sticky sessions

---

# 11. Amazon Route 53

**Service Model:** PaaS  

## What it actually is
DNS service that maps domain names to IP addresses.

## Purpose
Route user requests

## Real Usage
- Website domain routing

## Key Points
- Highly available
- Global service

## Subtopics
- Routing policies
- Hosted zones
- Health checks

---

# 12. AWS IAM (Identity and Access Management)

**Service Model:** SaaS (Conceptual)

## What it actually is
Centralized permission system for AWS resources.

## Purpose
Control access and security

## Real Usage
- User access control
- Service permissions

## Key Points
- JSON policies
- Role-based access

## Subtopics
- Users vs roles
- Policy structure
- Temporary credentials

---

# 13. Amazon CloudWatch

**Service Model:** PaaS  

## What it actually is
Monitoring and logging system for AWS resources.

## Purpose
Track performance and health

## Real Usage
- System monitoring
- Log aggregation

## Key Points
- Metrics and logs
- Alerts (alarms)

## Subtopics
- Dashboards
- Log groups
- Metrics

---

# 14. Amazon SQS (Simple Queue Service)

**Service Model:** PaaS  

## What it actually is
Distributed message queue system.

## Purpose
Decouple services

## Real Usage
- Async processing
- Background jobs

## Key Points
- Reliable messaging
- Handles traffic spikes

## Subtopics
- FIFO vs Standard
- Visibility timeout
- Dead-letter queue

---

# 15. Amazon CloudFront

**Service Model:** PaaS  

## What it actually is
Content Delivery Network (CDN) with global edge locations.

## Purpose
Deliver content faster worldwide

## Real Usage
- Static content delivery
- Media streaming

## Key Points
- Low latency
- Global distribution

## Subtopics
- Caching strategies
- TTL
- Origin configuration

---

# Service Model

| Model | Description | Examples |
|------|------------|---------|
| IaaS | Full control over infra | EC2, EBS, VPC |
| PaaS | Managed platform | RDS, Lambda, ECS |
| SaaS | Fully managed software | IAM |

---


