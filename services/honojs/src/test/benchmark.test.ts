import { test } from 'node:test';
import * as assert from 'node:assert';
import app from '../index.js';

// Start a real HTTP server (Bun) and run benchmark requests against it.
// This intentionally uses Bun.serve so the benchmark hits the network stack
// instead of calling `app.fetch()` directly.

let server: any = null;
let baseUrl = 'http://127.0.0.1:3000';

// test.before(() => {
//   // Require Bun to run the actual server. If Bun isn't present, instruct the user.
//   if (typeof Bun !== 'undefined' && typeof Bun.serve === 'function') {
//     server = Bun.serve({ fetch: app.fetch, port: 0 });
//     baseUrl = `http://127.0.0.1:${server.port}`;
//     console.log(`Started test server on ${baseUrl}`);
//   } else {
//     throw new Error('Bun runtime is required to run this benchmark against a real server. Run tests with `bun test`.');
//   }
// });

// test.after(() => {
//   if (server && typeof server.shutdown === 'function') {
//     server.shutdown();
//     console.log('Stopped test server');
//   }
// });

// TestRESTUsersEndpoint benchmarks the REST endpoint performance including reading the response body
test('TestHonojsRESTUsersEndpoint', async () => {
  const iterations = 20000;
  const url = `${baseUrl}/users`;

  // Warm up
  for (let i = 0; i < 10; i++) {
    const r = await fetch(url);
    assert.strictEqual(r.status, 200);
    await r.text();
  }

  // Measure
  const start = performance.now();
  for (let i = 0; i < iterations; i++) {
    const response = await fetch(url);
    assert.strictEqual(response.status, 200);

    // Read the body to simulate full request/response cycle
    const body = await response.text();
    assert.notStrictEqual(body.length, 0, 'expected non-empty response body');
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

