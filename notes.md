1. i'm gonna use helper folder to create helper functions that should be reused everywhere, like generating jwt everytime user, admin, refesh happens
--> "I have reusable functionality that doesn't naturally belong to a specific domain component, so I need somewhere sensible to put it."

My Authentication flow

POST /auth/register
        │
        ▼
registerHandler()
        │
        │ {name,email,password}
        ▼
authService.Register()
        │
        ├── validate data
        │
        ├── check whether user exists
        │
        ├── hash password
        │
        ▼
userRepository.Create()
        │
        ▼
   PostgreSQL

And then the response travels back

PostgreSQL
    │
    ▼
Repository
    │
    ▼
Auth Service
    │
    ▼
Handler
    │
    ▼
HTTP 201


--> Project is divided into 3 major sections;
1. authntication
2. wallet
3. Transactions

Authentication:
1. Register
2. Login

This is the high level overview of the services implemented

First of all, we can build the api gateway, because each internal networks and services need to be accessed by verified actors only.
This is like the the front door. Every request hits here first. JWT
validated, rate limited, role checked, then routed. 
Internal services never see an unauthenticated request. 
They trust the X-User-Id header
Gateway injects because only Gateway can set it.

2. Auth Service — issues identity. Register, login, JWT. That’s it.
When a user registers, it publishes a `user-registered` event to Kafka.
Payment Service picks that up and creates an account automatically.
No HTTP call between services. No coupling. Auth Service doesn’t know
Payment Service exists.


3. Payment Service — the brain. Every interesting problem lives here.
Double-entry ledger, idempotency, pessimistic locking, outbox pattern,
saga compensation. If this service gets something wrong, money is lost
or duplicated. There’s no margin for error here.

4. Fraud Service — the watchdog. Consumes every payment event and runs
it through a rules engine. High value transaction? Flag it. Round number?
Flag it. Self-transfer? Flag it. Flagged transactions go into a review
queue. Admins approve or reject. Rejection triggers a compensation event
back to Payment Service reverse the transaction, freeze the account.

5. Notification Service — the simplest service and the most
underappreciated. One job: consume payment events and notify the user.
It’s separate from Payment Service for one reason. If your email
provider is having a bad day, payments should still work. Decoupling
failure domains is the point.



PROJECT LOGS
Initially, i started building a simpple api gateway, in order to understand how to route my traffic to difference microservices efficiently
Then, i saw, that the main functionality of a gateway is the reverse proxy, where it magages, reads and intercepts user requests and all that, and masks server's  

