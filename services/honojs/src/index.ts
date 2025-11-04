import { Hono } from 'hono'
import { listUsers } from './controllers/users';

const app = new Hono()

// Health check endpoint
app.get('/health', (c) => {
  return c.json({ status: 'ok' });
});

// Users endpoint
app.get('/users', (c) => {
  const users = listUsers();
  return c.json(users);
});

export default app
