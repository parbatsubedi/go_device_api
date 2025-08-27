# Go Device Tracking API

A comprehensive device tracking and management system built with Go (Gin framework) featuring real-time device monitoring, user management, and location tracking capabilities.

## 🎯 Project Objectives

- **Device Management**: Track and manage mobile devices with detailed information including IMEI numbers, manufacturer, model, and status
- **Real-time Location Tracking**: Monitor device locations with GPS coordinates, battery levels, and network information
- **User Authentication**: Secure user management system with role-based access control
- **Admin Dashboard**: Web-based interface for monitoring device status and system analytics
- **Audit Trail**: Comprehensive logging of all system activities and changes

## 🚀 Features

### Core Functionality
- **Device CRUD Operations**: Create, read, update, and delete device records
- **Location Tracking**: Record and monitor device GPS coordinates with timestamps
- **User Management**: Register, authenticate, and manage users with different roles
- **Real-time Dashboard**: Monitor device status, battery levels, and location data
- **Audit Logging**: Track all system changes and user activities

### Technical Features
- RESTful API architecture
- JWT-based authentication
- PostgreSQL database with GORM ORM
- HTML templating with responsive design
- Real-time data visualization
- Docker containerization support

## 🛠️ Technology Stack

- **Backend**: Go 1.24.3, Gin Web Framework
- **Database**: PostgreSQL with PostGIS extension
- **ORM**: GORM
- **Frontend**: HTML5, CSS3, JavaScript
- **Authentication**: JWT (JSON Web Tokens)
- **Containerization**: Docker, Docker Compose
- **Logging**: Structured logging with slog

<!-- ## 📦 Installation & Setup

### Prerequisites

- Go 1.24.3 or later
- PostgreSQL database
- Docker and Docker Compose (optional)

### Method 1: Using Docker (Recommended)

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd go_api
   ```

2. **Start the database services**:
   ```bash
   docker-compose up -d
   ```

3. **Set up environment variables**:
   ```bash
   cp .env.example .env
   # Edit .env file with your configuration
   ```

4. **Install Go dependencies**:
   ```bash
   go mod download
   ```

5. **Run database migrations**:
   ```bash
   go run main.go
   # The application will automatically run migrations and seeders
   ```

6. **Start the application**:
   ```bash
   go run main.go
   ```

### Method 2: Manual Setup

1. **Install PostgreSQL** and create a database
2. **Set up environment variables**:
   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=your_database_name
   DB_USER=your_username
   DB_PASSWORD=your_password
   SERVER_PORT=8080
   APP_SECRET=your-jwt-secret-key
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

## 🏃‍♂️ Running the Application

### Development Mode
```bash
# Start the server
go run main.go

# Server will be available at: http://localhost:8080
```

### Production Mode
```bash
# Build the binary
go build -o go-api

# Run the binary
./go-api
``` -->

## 📊 API Endpoints

### Authentication Routes
- `POST /api/login` - User login
- `POST /api/register` - User registration

### Protected Routes (Require Authentication)
- `GET /api/devices` - Get all devices
- `POST /api/devices` - Create new device
- `GET /api/devices/:id` - Get device by ID
- `POST /api/devices/:id` - Update device
- `DELETE /api/devices/:id` - Delete device

### Admin Routes
- `GET /admin/` - Admin dashboard
- `GET /admin/dashboard` - Dashboard view

## 🖥️ User Interface

### Admin Dashboard
The web interface provides a comprehensive dashboard with:

- **Real-time Device Monitoring**: View active devices with status indicators
- **Location Tracking**: Interactive map showing device locations
- **Battery Status**: Monitor device battery levels
- **Statistics**: System metrics and performance indicators
- **Device Management**: CRUD operations through intuitive UI

### Accessing the UI
1. Start the application
2. Open your browser and navigate to `http://localhost:8080/admin`
3. Use the login functionality to access protected routes

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | Application port | 8080 |
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 5432 |
| `DB_NAME` | Database name | postgres |
| `DB_USER` | Database user | postgres |
| `DB_PASSWORD` | Database password | password |
| `APP_SECRET` | JWT secret key | your-secret-key-here |
| `APP_ENV` | Application environment | development |

### Database Configuration

The application uses PostgreSQL with the following tables:
- `users` - User accounts and authentication
- `devices` - Device information and metadata
- `device_locations` - GPS coordinates and tracking data
- `audit_trails` - System activity logs

## 📁 Project Structure

```
go_api/
├── apiRequests/          # Request models and validation
├── apiResponses/         # Response models and formatting
├── config/               # Configuration management
├── controllers/          # HTTP controllers
├── database/             # Database connection and handlers
├── helpers/              # Utility functions
├── interfaces/           # Service interfaces
├── middlewares/          # HTTP middleware
├── models/               # Database models
├── repository/           # Data access layer
├── routes/               # Route definitions
├── seeders/              # Database seed data
├── services/             # Business logic
├── static/               # Static assets (CSS, JS, images)
├── templates/            # HTML templates
├── main.go              # Application entry point
└── docker-compose.yml   # Docker configuration
```

## 🔧 Development

### Adding New Features
1. Create models in `models/` directory
2. Implement repositories in `repository/`
3. Add controllers in `controllers/`
4. Define routes in `routes/`
5. Update templates if needed

### Database Migrations
The application uses GORM auto-migration. Models are automatically migrated when the application starts.

### Testing
```bash
# Run tests
go test ./...
```

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Issues**
   - Ensure PostgreSQL is running
   - Check environment variables
   - Verify database credentials

2. **Port Conflicts**
   - Change `SERVER_PORT` in environment variables
   - Check if port 8080 is available

3. **Docker Issues**
   - Ensure Docker is running
   - Check container logs: `docker-compose logs`

### Logs
Application logs are stored in the `logs/` directory with daily rotation.

<!-- ## 📝 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📞 Support

For support and questions, please open an issue in the GitHub repository.

---

**Note**: This is a development version. Ensure proper security measures are implemented before deploying to production. -->
