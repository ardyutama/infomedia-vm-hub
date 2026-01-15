# Infomedia VM Hub

**Comprehensive VM Rental Management & Billing System for PT Infomedia Nusantara**

A full-stack application designed to digitalize and streamline VM rental operations, client contract management, and automated pricing calculations for virtual machine services.

## Overview

Infomedia VM Hub is an internal tool that enables PT Infomedia Nusantara to:
- Manage client VM rental contracts and specifications
- Track VM resource allocation (CPU, memory, storage, network)
- Calculate and manage dynamic pricing based on resource usage and contract terms
- Monitor VM status and resource utilization
- Generate billing and revenue reports
- Handle service requests and customer contracts

## Project Structure

```
infomedia-vm-hub/
├── backend/           # Go/Fiber REST API backend
│   ├── controllers/   # API endpoint handlers
│   ├── models/        # Database models
│   ├── database/      # Database connections & migrations
│   ├── handlers/      # Request/response handlers
│   ├── route/         # API route definitions
│   └── utils/         # Utility functions
└── frontend/          # Next.js React web application
    └── src/
        ├── app/       # Next.js app directory
        └── components/# Reusable React components
```

## Technology Stack

### Backend
- **Language:** Go 1.21.6
- **Framework:** Fiber v2 (Express-inspired web framework)
- **Database:** MySQL with GORM ORM
- **Authentication:** Password hashing with bcrypt
- **Architecture:** MVC Pattern with Controllers

### Frontend
- **Framework:** Next.js 14.1.0
- **UI Library:** React 18
- **Styling:** Tailwind CSS
- **UI Components:** Radix UI
- **Form Handling:** React Hook Form
- **Data Tables:** TanStack React Table

## Key Features

- 🔐 User authentication and role-based access control
- 🖥️ VM type and specification management
- 💾 Disk type and storage configuration
- 🌐 Network configuration management
- 💰 Dynamic pricing and billing system
- 📊 Revenue and resource utilization reports
- 📋 Client contract management
- 🔧 Service request handling
- 📈 Price history tracking
- 🏢 Multi-site location support

## Getting Started

### Backend Setup

```bash
cd backend
go mod download
go run main.go
```

Backend runs on: `http://localhost:3000`

### Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

Frontend runs on: `http://localhost:3000`

## API Endpoints

The backend provides RESTful API endpoints for:
- User management
- VM type and specification management
- Contract management
- Pricing and billing
- Service requests
- Resource allocation
- Revenue reporting

## Database Models

- User, Role, Contact
- Contract, ContractHistory, ContractRequestType
- VM Type, VM Status, VM Specification, VM SpecificationHistory
- Service Type, Request Based Type
- Price, Price History, Resource Per Hour
- Disk Type, Network, Operating System
- Project, Site Location
- Additional Features, Object Storage

## Environment Configuration

Both backend and frontend support environment-based configuration via `.env` files. See respective README files for detailed setup instructions.

## Development

This is an internal tool for PT Infomedia Nusantara. For development questions or contributions, please contact the development team.

## License

© PT Infomedia Nusantara - All Rights Reserved

---

**Quick Links:**
- [Backend Documentation](./backend/README.md)
- [Frontend Documentation](./frontend/README.md)
