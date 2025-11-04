import type { User } from '../models/user.js';

// ListUsers returns a list of hardcoded users.
// This matches the Go service's ListUsers function.
export function listUsers(): User[] {
  return [
    { id: 1, name: 'Alice', email: 'alice@example.com' },
    { id: 2, name: 'Bob', email: 'bob@example.com' },
    { id: 3, name: 'Charlie', email: 'charlie@example.com' },
  ];
}