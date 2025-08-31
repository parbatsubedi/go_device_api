# Go API - Device Tracking System

A comprehensive device tracking and management system built with Go (Golang) and Gin framework. This system provides real-time device monitoring, location tracking, and administrative dashboard capabilities.

## 🚀 Features

- **Device Management**: Create, read, update, and delete devices
- **Device Commands**: Send commands to devices (lock, unlock, wipe, ring, locate, status)
- **Real-time Location Tracking**: Track device locations with GPS coordinates
- **User Authentication**: Secure user management with role-based access
- **Admin Dashboard**: Comprehensive dashboard with real-time statistics
- **Audit Trail**: Complete logging of all system activities
- **Dual Input Support**: Accept both JSON and form data for all endpoints
- **RESTful API**: Clean, well-documented API endpoints

## 📊 Dashboard Features

The admin dashboard provides:
- Real-time device statistics (active/offline devices)
- Battery level monitoring with alerts
- Device status overview
- Interactive device list with live data
- Location tracking visualization

## 🛠️ Tech Stack

- **Backend**: Go (Golang) 1.24.3
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL with PostGIS extension
- **ORM**: GORM
- **Templating**: Go HTML Templates
- **Frontend**: HTML5, CSS3, JavaScript
- **Authentication**: JWT-based authentication

## 📦 Installation

### Prerequisites

- Go 1.24.3 or later
- PostgreSQL with PostGIS extension
- Docker and Docker Compose (optional)

### Using Docker (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd go_api
```

2. Start the database services:
```bash
docker-compose up -d
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. Install Go dependencies:
```bash
go mod download
```

5. Run database migrations:
```bash
go run main.go
# The application will automatically run migrations on startup
```

6. Start the application:
```bash
go run main.go
```

### Manual Setup

1. Install PostgreSQL and create a database
2. Install PostGIS extension:
```sql
CREATE EXTENSION postgis;
```

3. Set environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=your_database
export DB_USER=your_username
export DB_PASSWORD=your_password
export APP_PORT=8080
```

4. Run the application:
```bash
go run main.go
```

## 🔧 Configuration

### Environment Variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=go_api
DB_USER=postgres
DB_PASSWORD=password
APP_PORT=8080
APP_URL=http://localhost:8080
JWT_SECRET=your-secret-key
```

### Database Schema

The application automatically creates the following tables:
- `users` - User management and authentication
- `devices` - Device information and status
- `device_locations` - GPS coordinates and device metrics
- `audit_trails` - System activity logging

## 📡 API Endpoints

### Authentication
- `POST /api/login` - User login
- `POST /api/register` - User registration

### Devices (Protected)
- `GET /api/devices` - Get all devices
- `POST /api/devices` - Create new device
- `GET /api/devices/:id` - Get device by ID
- `POST /api/devices/:id` - Update device
- `DELETE /api/devices/:id` - Delete device

### Device Commands (Protected)
- `POST /api/device-commands/send/:device_id` - Send command to device
- `GET /api/device-commands/device/:device_id` - Get all commands for device
- `GET /api/device-commands/pending/:device_id` - Get pending commands for device
- `POST /api/device-commands/acknowledge/:commandId` - Acknowledge command

### Device Locations
- `POST /api/device-locations` - Create location entry

### Admin Dashboard
- `GET /admin/` - Admin dashboard
- `GET /admin/dashboard` - Dashboard alternative route
- `GET /admin/api/dashboard` - Dashboard data API

## 🎯 Usage

### Accessing the Dashboard

1. Start the application:
```bash
go run main.go
```

2. Open your browser and navigate to:
```
http://localhost:8080/admin/
```

3. The dashboard will display:
- Real-time device statistics
- Battery levels and alerts
- Device status overview
- Interactive device list

### API Usage Examples

#### Authentication
```bash
# Login (JSON)
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"mobile_no":"9843723270","password":"password"}'

# Login (Form Data)
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "mobile_no=9843723270&password=password"
```

#### Device Management
```bash
# Create device (JSON)
curl -X POST http://localhost:8080/api/devices \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{"name":"iPhone 15","device_imei1":"123456789012345","device_imei2":"123456789012346","manufacturer":"Apple","device_model":"iPhone 15"}'

# Create device (Form Data)
curl -X POST http://localhost:8080/api/devices \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d "name=iPhone 15&device_imei1=123456789012345&device_imei2=123456789012346&manufacturer=Apple&device_model=iPhone 15"

# Get all devices
curl -X GET http://localhost:8080/api/devices \
  -H "Authorization: Bearer <your-jwt-token>"
```

#### Device Commands
```bash
# Send command (JSON)
curl -X POST http://localhost:8080/api/device-commands/send/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{"command_type":"lock","command_data":"immediate"}'

# Send command (Form Data)
curl -X POST http://localhost:8080/api/device-commands/send/1 \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d "command_type=lock&command_data=immediate"

# Get device commands
curl -X GET http://localhost:8080/api/device-commands/device/1 \
  -H "Authorization: Bearer <your-jwt-token>"

# Get pending commands
curl -X GET http://localhost:8080/api/device-commands/pending/1 \
  -H "Authorization: Bearer <your-jwt-token>"

# Acknowledge command
curl -X POST http://localhost:8080/api/device-commands/acknowledge/1 \
  -H "Authorization: Bearer <your-jwt-token>"
```

### Supported Command Types

The device command system supports the following command types:

- **lock**: Lock the device immediately or schedule for later
- **unlock**: Unlock the device immediately or schedule for later
- **wipe**: Factory reset the device (immediate only)
- **ring**: Make the device ring loudly to locate it
- **locate**: Get current GPS location of the device
- **status**: Request device status and battery information

#### Command Data Options

- **immediate**: Execute command immediately
- **scheduled**: Schedule command for later execution
- **quiet**: Execute command silently (for lock/unlock)

## 📁 Project Structure

```
go_api/
├── apiRequests/          # Request models and validation
│   ├── deviceRequest.go      # Device CRUD request models
│   └── deviceCommandRequest.go # Device command request models
├── apiResponses/         # Response models and formatting
│   └── deviceCommandResponse.go # Device command response models
├── config/              # Configuration files
├── controllers/         # HTTP controllers
│   ├── deviceController.go   # Device CRUD operations
│   └── deviceCommandController.go # Device command operations
├── database/           # Database connection and setup
├── helpers/            # Utility functions
├── interfaces/         # Service interfaces
│   ├── device.go           # Device repository interface
│   └── device_command.go   # Device command repository interface
├── logs/               # Application logs
├── middlewares/        # HTTP middleware
├── models/             # Database models
│   └── device_command.go    # Device command model
├── repository/         # Data access layer
│   ├── deviceRepository.go      # Device repository implementation
│   └── deviceCommandRepository.go # Device command repository implementation
├── routes/             # Route definitions
├── seeders/           # Database seeders
├── services/          # Business logic
├── static/            # Static assets (CSS, JS, images)
├── templates/         # HTML templates
└── main.go           # Application entry point
```

## 🧪 Testing

Run tests with:
```bash
go test ./...
```

## 📈 Monitoring

The application includes:
- Automatic logging to `logs/app-YYYY-MM-DD.log`
- Daily log rotation
- Audit trail for all CRUD operations
- Performance monitoring

## 🔒 Security Features

- JWT-based authentication
- Password hashing with bcrypt
- CORS middleware
- Input validation
- SQL injection prevention
- XSS protection

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support and questions:
- Check the documentation
- Open an issue on GitHub
- Contact the development team

## 🚀 Deployment

### Production Deployment

1. Build the application:
```bash
go build -o go-api
```

2. Set production environment variables
3. Use a process manager (systemd, PM2, etc.)
4. Configure reverse proxy (nginx, Apache)
5. Set up SSL certificates

### Docker Production

```bash
docker build -t go-api .
docker run -p 8080:8080 --env-file .env go-api
```

---

**Note**: Make sure to replace placeholder values and update the documentation as needed for your specific deployment environment.
