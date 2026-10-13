// SPDX-License-Identifier: AGPL-3.0-or-later

import { mount } from 'svelte'
import './app.css'
import App from './App.svelte'
import { initializeLocale } from './lib/i18n.svelte.js'

await initializeLocale()

export default mount(App, { target: document.body })
