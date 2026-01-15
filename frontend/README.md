# Infomedia VM Hub - Frontend

**Modern web dashboard for VM rental management, client contracts, and billing**

A Next.js React application providing an intuitive user interface for PT Infomedia Nusantara's VM rental and pricing management operations.

## Overview

The frontend application provides:
- User-friendly dashboard for VM management
- Client contract and billing administration
- Resource allocation and specification management
- Pricing and revenue tracking
- Service request handling
- Role-based access control
- Responsive design for desktop and tablet use

## Technology Stack

- **Framework:** Next.js 14.1.0
- **UI Library:** React 18
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **UI Components:** Radix UI
- **Form Handling:** React Hook Form
- **Data Tables:** TanStack React Table
- **Icons:** Lucide React
- **State Management:** React Context/Hooks

## Project Structure

```
frontend/
├── src/
│   ├── app/              # Next.js App Router
│   │   ├── (protected)/  # Protected routes requiring auth
│   │   │   ├── (forms)/  # Form pages for creating/editing
│   │   │   ├── admin/    # Admin dashboard
│   │   │   └── management/ # Management pages
│   │   ├── features/     # Feature-specific logic
│   │   └── globals.css   # Global styles
│   ├── components/       # Reusable React components
│   │   └── ui/           # UI component library
│   ├── hooks/            # Custom React hooks
│   ├── lib/              # Utility functions & helpers
│   └── public/           # Static assets
├── package.json
├── tsconfig.json
├── tailwind.config.ts
└── next.config.js
```

## Getting Started

### Prerequisites
- Node.js 16.8+
- npm, yarn, pnpm, or bun
- Backend API running at `http://localhost:3000`

### Installation

1. Clone the repository:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
# or
yarn install
# or
pnpm install
# or
bun install
```

3. Create `.env.local` for environment variables:
```env
NEXT_PUBLIC_API_URL=http://localhost:3000
```

4. Run the development server:
```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) to view the application.

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint

## Key Features

### Dashboard
- Overview of VM inventory
- Contract status monitoring
- Revenue and billing summary
- Resource utilization metrics

### VM Management
- Create and edit VM specifications
- Assign VMs to projects/contracts
- Monitor VM status and health
- Track resource allocation

### Contract Management
- Create and manage client contracts
- Define contract terms and pricing
- Track contract history
- Generate contract documents

### Billing & Pricing
- Configure pricing models
- Track price history
- Generate invoices
- Revenue reporting

### Administration
- User management
- Role and permission configuration
- System settings
- Resource type management

## UI Components

Built with Radix UI and Tailwind CSS:

- **Forms:** Input, Select, Checkbox, Radio, Calendar
- **Displays:** Table, Badge, Avatar, Card
- **Dialogs:** Dialog, Popover, Toast
- **Navigation:** Command/Search, Tabs, Separator
- **Data:** Tables with sorting and filtering

## Styling

- **Framework:** Tailwind CSS
- **Customization:** `tailwind.config.ts`
- **CSS:** `src/app/globals.css`
- **Component Variants:** Class variance authority (CVA)

## Environment Configuration

### Development
```env
NEXT_PUBLIC_API_URL=http://localhost:3000
NODE_ENV=development
```

### Production
```env
NEXT_PUBLIC_API_URL=https://api.infomedia.example.com
NODE_ENV=production
```

## Development Guidelines

### Code Style
- TypeScript for type safety
- React functional components
- Custom hooks for logic reuse
- Tailwind CSS for styling

### Component Structure
```typescript
// Example component structure
export interface ComponentProps {
  // Props definition
}

export function ComponentName(props: ComponentProps) {
  // Component logic
  return (
    // JSX
  );
}
```

## Building for Production

```bash
npm run build
npm start
```

The application will be optimized and ready for deployment.

## Testing

The project structure supports component testing. Add tests in the same directory as components:
```
components/
├── ComponentName.tsx
├── ComponentName.test.tsx
└── ...
```

## Deployment

### Vercel (Recommended)
```bash
# Deploy directly from GitHub
# Visit https://vercel.com/new
```

### Manual Deployment
```bash
npm run build
# Deploy the 'out' directory to your hosting provider
```

## Performance Optimization

- Image optimization with Next.js Image component
- Code splitting and lazy loading
- Font optimization with next/font
- CSS-in-JS optimization with Tailwind

## Security

- Environment variables for sensitive data
- Protected routes with authentication checks
- Input validation and sanitization
- CORS and API security configurations

## Contributing

This is an internal tool for PT Infomedia Nusantara. For development guidelines and contributions, contact the development team.

## License

© PT Infomedia Nusantara - All Rights Reserved

---

**Related Documentation:**
- [Backend README](../backend/README.md)
- [Project README](../README.md)

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/basic-features/font-optimization) to automatically optimize and load Inter, a custom Google Font.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
