# Welcome to this humble repo of mine

This repo is consisted of both backend and frontend app for Voucher Management System.

Here are simple introduction to both systems and how to run them.

## Frontend
### What is this built on?
The frontend side of the app is built with NextJs 13 written in Typescript using the pages directory. Here are some notable tools that were used :
- TailwindCSS for styling
- React-Query & axios for data fetching and server state management
- Jotai for local global state management

### How to run this program?
1. Make sure to go inside the `./frontend` directory
2. Run the following commands
```sh
pnpm i

# to run in development mode
pnpm dev

# to run in production mode
pnpm start

```

## Backend
### What is this built on?
The backend side of the app is built with Fiber Framework written in Go. Here are some notable tools that were used :
- Gorm for ORM
- PostgreSQL as the main database
- JWT for authentication and authorization

### How to run this program?
1. Make sure to go inside the `./backend` directory
2. Create a `.env` file in the root directory that stores the following enviroment variables
```
DB_URL
JWT_SECRET
```
3. Run the following commands
```sh
# to create temporary build
go run .

# to create executable build
go build

```
