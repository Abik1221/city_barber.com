# City Barber

City Barber is a web application for managing barber shops, bookings, and user accounts. It includes features like user authentication, Google OAuth login, password reset via SMS/email, and more.

  
## Features

- User registration and login
- Google OAuth login
- Forget password with SMS/email options
- Barber and shop management
- Booking and payment management
- Admin panel for managing promocodes and users

## Technologies Used

- **Backend**: Go (Gin framework)
- **Database**: MySQL
- **Authentication**: JWT
- **Email/SMS**: Third-party APIs
- **Google OAuth**: Google OAuth2

## Prerequisites

- Go 1.20 or higher
- MySQL 8.0 or higher
- Google OAuth credentials
- Email and SMS API keys

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/NahomKeneni/city-barber.git
   cd city-barber

2. Install dependencies for gin frame work

```bash
go get -u github.com/gin-gonic/gin
```

3. Copy the the following

```bash
cp .env.example .env
```

4. Run database migrations:

```bash
go run migrations/migrations.go
```

5. starting the server
```bash
go run cmd/city_barber/main.go
```

# Author

nahomkeneni4@gmail.com
