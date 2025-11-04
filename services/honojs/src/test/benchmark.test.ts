import { test } from 'node:test';
import * as assert from 'node:assert';
import app from '../index.js';

// TestRESTUsersEndpoint benchmarks the REST endpoint performance
test('TestRESTUsersEndpoint', async () => {
  const iterations = 10000;
  const request = new Request('http://localhost/users');

  // Warm up
  for (let i = 0; i < 10; i++) {
    await app.fetch(request);
  }

  // Measure
  const start = performance.now();
  for (let i = 0; i < iterations; i++) {
    const response = await app.fetch(request);
    assert.strictEqual(response.status, 200);
  }
  const elapsed = performance.now() - start;

  const avgTime = elapsed / iterations;
  const reqPerSec = 1000 / avgTime;

  console.log('REST /users endpoint:');
  console.log(`  Iterations: ${iterations}`);
  console.log(`  Total time: ${elapsed.toFixed(2)}ms`);
  console.log(`  Time per request: ${avgTime.toFixed(4)}ms`);
  console.log(`  Requests per second: ${Math.round(reqPerSec)}`);
});

// TestRESTUsersEndpointWithBodyRead benchmarks including reading the response body
test('TestRESTUsersEndpointWithBodyRead', async () => {
  const iterations = 10000;
  const request = new Request('http://localhost/users');

  // Warm up
  for (let i = 0; i < 10; i++) {
    await app.fetch(request);
  }

  // Measure
  const start = performance.now();
  for (let i = 0; i < iterations; i++) {
    const response = await app.fetch(request);
    assert.strictEqual(response.status, 200);

    // Read the body to simulate full request/response cycle
    const body = await response.text();
    assert.notStrictEqual(body.length, 0, 'expected non-empty response body');
  }
  const elapsed = performance.now() - start;

  const avgTime = elapsed / iterations;
  const reqPerSec = 1000 / avgTime;

  console.log('REST /users endpoint (with body read):');
  console.log(`  Iterations: ${iterations}`);
  console.log(`  Total time: ${elapsed.toFixed(2)}ms`);
  console.log(`  Time per request: ${avgTime.toFixed(4)}ms`);
  console.log(`  Requests per second: ${Math.round(reqPerSec)}`);
});

