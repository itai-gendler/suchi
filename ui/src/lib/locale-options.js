// SPDX-License-Identifier: AGPL-3.0-or-later

import { DEFAULT_LOCALE, localeRuntimeState, localeState } from './i18n.svelte.js'

export function availableLocales() {
  return [...localeRuntimeState.available.keys()].sort((a, b) => a === DEFAULT_LOCALE ? -1 : b === DEFAULT_LOCALE ? 1 : a.localeCompare(b))
}

export function displayLanguage(locale, displayLocale = localeState.locale) {
  try { return new Intl.DisplayNames([displayLocale], { type: 'language' }).of(locale) || locale } catch { return locale }
}
