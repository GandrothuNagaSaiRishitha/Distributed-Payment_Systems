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
