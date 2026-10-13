// SPDX-License-Identifier: AGPL-3.0-or-later

import {
  applyLocaleState, COOKIE_NAME, DEFAULT_LOCALE, localeRuntimeState,
} from './i18n.svelte.js'

const loaders = typeof import.meta.glob === 'function'
  ? import.meta.glob('../i18n/locales/*.json')
  : {}

for (const path of Object.keys(loaders)) {
  const match = path.match(/\/([^/]+)\.json$/)
  if (match) localeRuntimeState.available.set(match[1], { locale: match[1], direction: 'ltr' })
}

function normalizeTag(value) {
  const raw = String(value || '').trim().replaceAll('_', '-')
  if (!raw) return ''
  try { return Intl.getCanonicalLocales(raw)[0] || '' } catch { return '' }
}

function resolveLocale(value) {
  const exact = normalizeTag(value)
  if (!exact) return ''
  if (localeRuntimeState.available.has(exact)) return exact
  const base = exact.split('-')[0]
  return localeRuntimeState.available.has(base) ? base : ''
}

function cookieLocale() {
  if (typeof document === 'undefined') return ''
  const prefix = `${COOKIE_NAME}=`
  for (const part of document.cookie.split(';')) {
    const item = part.trim()
    if (item.startsWith(prefix)) return decodeURIComponent(item.slice(prefix.length))
  }
  return ''
}

function browserLocale() {
  if (typeof navigator === 'undefined') return DEFAULT_LOCALE
  for (const candidate of navigator.languages || [navigator.language]) {
    const locale = resolveLocale(candidate)
    if (locale) return locale
  }
  return DEFAULT_LOCALE
}

async function loadCatalog(locale) {
  if (localeRuntimeState.catalogs.has(locale)) return localeRuntimeState.catalogs.get(locale)
  const entry = Object.entries(loaders).find(([path]) => path.endsWith(`/${locale}.json`))
  if (!entry) throw new Error(`locale catalog is not available: ${locale}`)
  const module = await entry[1]()
  const catalog = module.default || module
  validateCatalog(locale, catalog)
  localeRuntimeState.catalogs.set(locale, catalog)
  localeRuntimeState.available.set(locale, { locale, direction: catalogDirection(locale, catalog) })
  return catalog
}

function validateCatalog(locale, catalog) {
  if (!catalog || typeof catalog !== 'object' || Array.isArray(catalog)) {
    throw new Error(`invalid locale catalog: ${locale}`)
  }
  if (catalog['_meta.locale'] !== locale) throw new Error(`locale metadata mismatch: ${locale}`)
  catalogDirection(locale, catalog)
  const english = localeRuntimeState.fallback
  const englishKeys = Object.keys(english).filter((key) => !key.startsWith('_meta.')).sort()
  for (const key of englishKeys) {
    if (typeof catalog[key] !== 'string' || catalog[key] === '') throw new Error(`locale catalog is incomplete: ${locale}`)
    if (!samePlaceholders(english[key], catalog[key])) {
      throw new Error(`locale placeholders differ: ${locale}:${key}`)
    }
  }
  for (const key of Object.keys(catalog).filter((key) => !key.startsWith('_meta.') && !Object.hasOwn(english, key))) {
    const match = key.match(/^(.*)\.(zero|two|few|many)$/)
    const fallback = match && english[`${match[1]}.other`]
    if (!fallback || typeof catalog[key] !== 'string' || catalog[key] === '') {
      throw new Error(`locale catalog has unsupported message: ${locale}:${key}`)
    }
    if (!samePlaceholders(fallback, catalog[key])) {
      throw new Error(`locale placeholders differ: ${locale}:${key}`)
    }
  }
}

function catalogDirection(locale, catalog) {
  const direction = catalog['_meta.direction']
  if (direction !== 'ltr' && direction !== 'rtl') throw new Error(`invalid locale direction: ${locale}`)
  return direction
}

function samePlaceholders(expectedMessage, actualMessage) {
  const expected = placeholders(expectedMessage)
  const actual = placeholders(actualMessage)
  return expected.length === actual.length && expected.every((name, index) => name === actual[index])
}

function placeholders(message) {
  return [...new Set([...String(message).matchAll(/\{([A-Za-z][A-Za-z0-9_]*)\}/g)].map((match) => match[1]))].sort()
}

export async function initializeLocale(preference = '') {
  const requested = preference === 'auto' || preference === ''
    ? resolveLocale(cookieLocale()) || browserLocale()
    : resolveLocale(preference)
  const locale = requested || DEFAULT_LOCALE
  try {
    applyLocaleState(locale, await loadCatalog(locale))
  } catch {
    applyLocaleState(DEFAULT_LOCALE, localeRuntimeState.fallback)
  }
  return locale
}

export async function setLocale(preference, { persist = true, reload = false } = {}) {
  const requested = preference === 'auto' || preference === '' ? browserLocale() : resolveLocale(preference)
  const locale = requested || DEFAULT_LOCALE
  applyLocaleState(locale, await loadCatalog(locale))
  if (persist && typeof document !== 'undefined') {
    if (preference === 'auto' || preference === '') {
      document.cookie = `${COOKIE_NAME}=; Path=/; Max-Age=0; SameSite=Lax`
    } else {
      document.cookie = `${COOKIE_NAME}=${encodeURIComponent(locale)}; Path=/; Max-Age=31536000; SameSite=Lax`
    }
  }
  if (reload && typeof location !== 'undefined') location.reload()
  return locale
}

export function registerCatalogForTest(locale, catalog) {
  const normalized = normalizeTag(locale)
  validateCatalog(normalized, catalog)
  localeRuntimeState.catalogs.set(normalized, catalog)
  localeRuntimeState.available.set(normalized, { locale: normalized, direction: catalogDirection(normalized, catalog) })
}
