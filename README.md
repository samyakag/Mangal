# Mangal Chai - E-Commerce Platform

A modern e-commerce platform for Mangal Chai, built with React and Go.

## Tech Stack

- **Frontend**: React 19 + TypeScript + Vite + TailwindCSS
- **Backend**: Go + Gin Framework
- **Database**: MongoDB
- **Payment**: Razorpay Integration
- **Deployment**: GitHub Pages (Frontend) + Render (Backend) + MongoDB Atlas (Database)

## Quick Start

### Prerequisites

- Node.js 20+
- Go 1.23+
- MongoDB Atlas account (free tier)
- Razorpay account (for payments)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/samyakag/Mangal.git
   cd Mangal
   ```

2. **Set up environment variables**

   Backend (`backend/.env`):
   ```bash
   cp backend/.env.example backend/.env
   # Edit backend/.env with your values
   ```

   Frontend (`frontend/.env`):
   ```bash
   cp frontend/.env.example frontend/.env
   # Edit frontend/.env with your values
   ```

3. **Start the backend**
   ```bash
   cd backend
   go run main.go
   ```

4. **Start the frontend** (in a new terminal)
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

5. **Visit** http://localhost:5173

## Deployment

This project is configured for free deployment using:
- **Frontend**: GitHub Pages
- **Backend**: Render (Free Tier)
- **Database**: MongoDB Atlas (M0 Free Tier)

### Quick Deploy (30 minutes)

Follow the [QUICKSTART.md](QUICKSTART.md) guide to deploy in under 30 minutes.

### Deployment Guides

- **[QUICKSTART.md](QUICKSTART.md)** - Get deployed in 30 minutes
- **[SETUP_GUIDE.md](SETUP_GUIDE.md)** - Detailed step-by-step guide
- **[DEPLOYMENT.md](DEPLOYMENT.md)** - Complete deployment strategies
- **[DEPLOYMENT_SUMMARY.md](DEPLOYMENT_SUMMARY.md)** - Configuration overview

### Deployment Checklist

Track your deployment progress with [.deployment-checklist.md](.deployment-checklist.md)

## Project Structure

```
Mangal/
├── backend/                 # Go backend
│   ├── controllers/        # API controllers
│   ├── services/           # Business logic
│   ├── repositories/       # Database layer
│   ├── models/             # Data models
│   ├── database/           # Database connection
│   ├── Dockerfile          # Docker configuration
│   └── main.go             # Entry point
├── frontend/               # React frontend
│   ├── src/
│   │   ├── components/    # React components
│   │   ├── hooks/         # Custom hooks
│   │   ├── services/      # API services
│   │   └── types.ts       # TypeScript types
│   ├── public/            # Static assets
│   ├── Dockerfile         # Docker configuration
│   └── vite.config.ts     # Vite configuration
├── .github/
│   └── workflows/         # CI/CD pipelines
├── render.yaml            # Render deployment config
└── README.md
```

## Features

- Browse products by category
- Product search and filtering
- Shopping cart functionality
- Razorpay payment integration
- Order management
- Responsive design
- Docker support
- CI/CD with GitHub Actions

## API Endpoints

### Products
- `GET /api/products` - Get all products
- `GET /api/products/:id` - Get product by ID
- `GET /api/products/category/:category` - Get products by category
- `GET /api/categories` - Get all categories

### Orders
- `POST /api/orders` - Create new order
- `GET /api/orders/:id` - Get order by ID

### Payments
- `POST /api/payments/create-order` - Create Razorpay order

### Health
- `GET /api/health` - Health check endpoint

## Environment Variables

### Backend
| Variable | Description | Required |
|----------|-------------|----------|
| MONGO_URL | MongoDB connection string | Yes |
| RAZORPAY_KEY_ID | Razorpay API key | Yes |
| RAZORPAY_KEY_SECRET | Razorpay secret key | Yes |
| PORT | Server port | Yes |
| GIN_MODE | Gin mode (debug/release) | Yes |
| ALLOWED_ORIGINS | CORS allowed origins | No |

### Frontend
| Variable | Description | Required |
|----------|-------------|----------|
| VITE_API_BASE_URL | Backend API URL | Yes |
| VITE_RAZORPAY_KEY_ID | Razorpay API key | Yes |

## Testing

### Frontend Tests
```bash
cd frontend
npm test
```

### Backend Tests
```bash
cd backend
go test -v ./...
```

## Docker Support

### Build and run backend
```bash
cd backend
docker build -t mangal-backend .
docker run -p 8001:8001 --env-file .env mangal-backend
```

### Build and run frontend
```bash
cd frontend
docker build -t mangal-frontend .
docker run -p 80:80 mangal-frontend
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Write/update tests
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

For issues and questions:
- Check the [troubleshooting guide](SETUP_GUIDE.md#troubleshooting)
- Review deployment documentation
- Open an issue on GitHub

## Deployment Status

### Development
- Frontend: https://samyakag.github.io/Mangal/
- Backend: _[Configure on Render]_
- Database: MongoDB Atlas M0

### Production
- _Not yet deployed_

---

Made with ❤️ by Samyak
