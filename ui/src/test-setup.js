// SPDX-License-Identifier: AGPL-3.0-or-later

// Svelte compiles runes before browser modules execute. Bun loads plain .svelte.js
// modules directly in unit tests, so provide the state identity used by those tests.
globalThis.$state = value => value
