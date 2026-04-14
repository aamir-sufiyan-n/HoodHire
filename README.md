# HoodHire — Backend System Documentation

## Overview

HoodHire is a location-based job marketplace that connects local businesses with part-time and full-time job seekers in their neighborhood. The platform is built around trust, privacy, and real-world hiring workflows — from discovery to a confirmed hire.

The backend is built with Go (Fiber) following Clean Architecture principles, ensuring modularity, testability, and scalability. The frontend is built with React, Tailwind CSS, and Antigravity.

---

## Tech Stack

### Backend
- **Go** with **Fiber** — high-performance HTTP framework
- **GORM** — ORM for PostgreSQL
- **PostgreSQL** — primary relational database
- **Redis** — OTP caching, session management, rate limiting

### Frontend
- **React** — component-based UI
- **Tailwind CSS** — utility-first styling
- **Antigravity** — UI component framework

### Authentication & Security
- **JWT** — access and refresh token pair
- **bcrypt** — password hashing
- **OTP verification** — email-based OTP stored in Redis with TTL

### Infrastructure & Integrations
- **Docker & Docker Compose** — containerized deployment
- **Razorpay** — subscription payment gateway
- **Cloudinary** — image and PDF (resume) storage
- **gomail** — transactional email delivery (SMTP)
- **robfig/cron** — cron job scheduling

---

## Architecture

The project follows Clean Architecture with strict layer separation:

```
HTTP Request
    ↓
Middleware (Auth, Role, Permission, Feature, Block)
    ↓
Controller (Request/Response handling)
    ↓
Service (Business logic)
    ↓
Repository (Database queries)
    ↓
PostgreSQL / Redis
```

This separation ensures:
- Controllers stay thin
- Business logic is centralized in services
- Database logic is isolated in repositories
- Easy to test each layer independently

---

## Folder Structure

```
hoodhire/
│
├── cmd/
│   └── main.go
│
├── config/
│   ├── dbConfig.go
│   └── redis.go
│
├── database/
│   ├── db.go
│   └── seed.go
│
├── internal/
│   ├── app/
│   │   └── app.go
│   │
│   ├── controllers/
│   │   ├── authController.go
│   │   ├── bondController.go
│   │   ├── categoryController.go
│   │   ├── follow.go
│   │   ├── hirerController.go
│   │   ├── jobControllers.go
│   │   ├── roleControllers.go
│   │   ├── seekerController.go
│   │   ├── subscription.go
│   │   ├── ticketController.go
│   │   └── webConfig.go
│   │
│   ├── middlewares/
│   │   ├── authMiddleware.go
│   │   ├── blockedMiddleware.go
│   │   ├── configMiddleware.go
│   │   ├── permissionMiddleware.go
│   │   ├── roleMiddleware.go
│   │   └── serviceMiddleware.go
│   │
│   ├── repositories/
│   │   ├── authRepos.go
│   │   ├── bondRepo.go
│   │   ├── categoryRepository.go
│   │   ├── followReviewRepo.go
│   │   ├── hirerRepos.go
│   │   ├── jobRepos.go
│   │   ├── roleRepo.go
│   │   ├── seekerRepos.go
│   │   ├── subscriptionRepo.go
│   │   ├── ticketRepo.go
│   │   └── webConfig.go
│   │
│   ├── routes/
│   │   ├── adminRoutes.go
│   │   └── userRoutes.go
│   │
│   ├── services/
│   │   ├── authServices.go
│   │   ├── bondServices.go
│   │   ├── categoryServices.go
│   │   ├── follow.go
│   │   ├── hirerAndBusinessServices.go
│   │   ├── jobServices.go
│   │   ├── roleServices.go
│   │   ├── seekerServ.go
│   │   ├── subscription.go
│   │   ├── ticketServices.go
│   │   └── webConfig.go
│   │
│   └── structures/
│       ├── dto/
│       │   ├── adminDto.go
│       │   ├── authDto.go
│       │   ├── othersDto.go
│       │   └── profileDto.go
│       │
│       └── models/
│           ├── admin.go
│           ├── bond.go
│           ├── followReview.go
│           ├── hirer.go
│           ├── jobs.go
│           ├── seeker.go
│           ├── subscription.go
│           ├── ticket.go
│           └── user.go
│
└── utils/
    ├── bcrypt.go
    ├── bind.go
    ├── cloudinary.go
    ├── cookie.go
    ├── cron.go
    ├── jwt.go
    ├── mail.go
    └── pdf.go
```

---

## Core Features

### 1. Authentication System

- Email-based OTP registration with Redis TTL (5 minute expiry)
- JWT access + refresh token pair stored in HttpOnly cookies
- bcrypt password hashing
- Separate registration flows for Seekers and Hirers
- Role-based route protection (`seeker`, `hirer`, `admin`, custom roles)
- Blocked user middleware — blocked accounts are rejected at middleware level

### 2. Seeker Profile System

Seekers can build a complete profile including:
- Personal details (name, age, gender, location)
- Education history
- Work experience
- Work preferences (shift, days, availability)
- Job interest categories (up to 5)
- Profile picture (Cloudinary)
- Resume upload (PDF, Cloudinary)

### 3. Hirer & Business Profile System

Hirers register with a personal profile and a linked business profile:
- Business details (name, niche, location, employee count, website)
- Business profile picture
- Admin approval required before job posting
- Subscription-based verification badge

### 4. Job Management

- Full CRUD for job postings
- Rich job description (title, type, shift, salary range, age/gender preference, schedule, skills, responsibilities)
- Resume requirement flag per job
- Job status management (`open`, `closed`, `filled`)
- Deadline-based auto-close via cron job
- Locality and category-based filtering
- Free plan limited to 5 open jobs simultaneously

### 5. Application & Bond System (Core Innovation)

This is the core hiring workflow of HoodHire:

```
Seeker applies to job (with optional resume)
        ↓
Application stored as pending
        ↓
Hirer reviews applicants and their profiles
        ↓
Hirer accepts → Bond is automatically created
        ↓
Contact information revealed only after Bond
        ↓
Chat unlocked between seeker and hirer
```

Privacy is enforced at the data layer — phone numbers and emails are only exposed after a Bond exists between the two parties. This prevents spam and builds trust.

### 6. Real-Time Messaging (WebSocket Microservice)

Live chat is built as a separate microservice running on a different port:

- WebSocket-based real-time messaging
- Only available after a Bond is created between seeker and hirer
- Message persistence in PostgreSQL
- Unread count tracking
- File/image sharing via upload endpoint
- Auth via callback to main service's `/auth/verify` endpoint
- Read receipts and conversation list

### 7. Follow, Favorites & Reviews

- Seekers can follow businesses
- Seekers can favorite businesses and jobs for later
- Seekers can leave reviews with star ratings on businesses
- Review management (create, update, delete)
- Aggregate rating and review count on business profile

### 8. Subscription System (Razorpay)

Hirers can subscribe to HoodHire Pro:

| Plan | Price | Duration |
|------|-------|----------|
| Monthly | ₹199 | 30 days |
| Yearly | ₹999 | 365 days |

**Pro benefits:**
- Unlimited open job postings (free: max 5)
- Verified badge on business profile and job listings
- Priority listing in job search results

**Payment flow:**
1. Hirer selects plan → backend creates Razorpay order
2. Frontend opens Razorpay checkout popup
3. Hirer completes payment with test/real card
4. Frontend sends `order_id`, `payment_id`, `signature` to backend
5. Backend verifies HMAC SHA256 signature
6. Subscription activated, business `is_verified` set to `true`

**Cron job (runs daily at midnight):**
- 7 days before expiry → send reminder email
- On expiry → set `is_verified` to `false`, update subscription status
- All handled automatically without manual intervention

### 9. Ticket / Support System

- Seekers and hirers can raise complaints or reports
- Admin can review, resolve, or dismiss tickets
- Filter by status (`open`, `reviewed`, `resolved`, `dismissed`)
- Filter by business

### 10. Admin Panel with RBAC

The admin panel is fully dynamic with role-based access control:

**Role & Permission Management:**
- Admin can create custom roles (e.g. `manager`, `moderator`)
- Each role has toggleable permissions:
  - `user_management`
  - `business_management`
  - `ticket_management`
  - `rbac_control`
  - `web_config_control`
  - `jobs_management`
- Permission middleware checks DB on every request — no hardcoded role checks
- Admin panel sidebar hides/disables tabs based on user permissions

**Admin Capabilities:**
- Full CRUD on users, businesses, jobs, categories, roles, permissions
- Block/unblock users and businesses
- Approve or reject business registrations
- Export data as PDF (users, jobs, businesses)
- Manage subscription plans

### 11. Web Configuration System

Admin can toggle platform features on/off in real time without code changes:

| Feature Flag | What it controls |
|---|---|
| `user_registration` | Allow new seeker signups |
| `business_registration` | Allow new hirer signups |
| `job_posting` | Allow hirers to post jobs |
| `job_applying` | Allow seekers to apply |
| `chat` | Enable/disable messaging |

Changes take effect immediately — no server restart needed.

### 12. Category Management

- Full CRUD on job categories
- Stats endpoint — job count per category with sort by most/least popular
- Pagination support
- Categories used for seeker interest matching and job filtering

### 13. Transactional Emails

All automated emails use HTML templates via SMTP (gomail):

- OTP verification email
- Application accepted notification (to seeker)
- Application rejected notification (to seeker)
- Job cancelled/deleted notification (to pending applicants)
- Subscription activated confirmation
- Subscription expiry reminder (7 days before)
- Subscription expired notification

---

## Database Design

### Key Entities

| Entity | Description |
|--------|-------------|
| `users` | Base identity for all account types |
| `seekers` | Extended seeker profile linked to user |
| `hirers` | Extended hirer profile linked to user |
| `businesses` | Business profile linked to hirer |
| `jobs` | Job postings with description |
| `job_applications` | Applications with status and optional resume |
| `bonds` | Confirmed hires — unlocks contact and chat |
| `subscriptions` | Payment and plan records per hirer |
| `plans` | Subscription plan definitions |
| `roles` | Dynamic roles |
| `permissions` | Platform permissions |
| `role_permissions` | Join table with `is_allowed` toggle |
| `web_configs` | Feature flags for platform features |
| `tickets` | Support tickets from users |
| `reviews` | Business reviews from seekers |

### Privacy Model

Contact information (phone, email) is only exposed after a `Bond` record exists between the seeker and hirer. This is enforced at the service layer, not just the frontend.

---

## API Design

### Route Groups

| Prefix | Auth | Description |
|--------|------|-------------|
| `/auth/*` | None | Registration, login, OTP |
| `/seeker/*` | JWT + seeker role | Seeker profile and actions |
| `/hirer/*` | JWT + hirer role | Hirer profile and actions |
| `/admin/*` | JWT + permission middleware | Admin panel APIs |
| `/me/permissions` | JWT | Current user's permissions |
| Public routes | None | Jobs, businesses, categories |

### Middleware Stack (per request)

```
AuthMiddleware → extracts user from JWT cookie
RoleMiddleware → checks user.role string
PermissionMiddleware → checks role_permissions table
BlockMiddleware → rejects blocked accounts
FeatureMiddleware → checks web_config feature flags
```

---

## Security

- Passwords hashed with bcrypt (default cost)
- JWT stored in HttpOnly cookies — not accessible to JavaScript
- OTP stored in Redis with 5 minute TTL — expires automatically
- SQL injection protection via GORM parameterized queries
- Input validation on all DTOs using Go validator
- Permission checks at middleware level — cannot be bypassed by frontend
- HMAC SHA256 signature verification for Razorpay payments
- Blocked users rejected at middleware before reaching any controller

---

## Deployment

- Fully Dockerized with Docker Compose
- Environment-based configuration via `.env`
- Stateless API design — horizontally scalable
- Chat microservice runs as a separate container on port 8081
- Main API runs on port 8080

---

## What This Project Demonstrates

HoodHire is not a standard CRUD project. It demonstrates:

- **Real-world system design** — privacy-first architecture, bond-based information disclosure
- **Clean Architecture in Go** — strict layer separation across controllers, services, repositories
- **Payment integration** — Razorpay order creation, HMAC signature verification, subscription lifecycle
- **RBAC at scale** — fully dynamic, admin-controlled, DB-driven permission system
- **Real-time systems** — WebSocket microservice with auth, persistence, and read tracking
- **Cron job automation** — subscription expiry handling without manual intervention
- **Microservice thinking** — chat separated as its own service with its own auth callback
- **Production-level email system** — transactional emails for every meaningful user event
- **Admin tooling** — full platform control panel with feature flags and role management

