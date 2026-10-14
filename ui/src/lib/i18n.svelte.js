// SPDX-License-Identifier: AGPL-3.0-or-later

const englishModule = await import('../i18n/locales/en.json')
const english = englishModule.default || englishModule

export const DEFAULT_LOCALE = 'en'
export const COOKIE_NAME = 'suchi_locale'

const defaultDirection = english['_meta.direction']
export const localeRuntimeState = {
  fallback: english,
  catalogs: new Map([[DEFAULT_LOCALE, english]]),
  available: new Map([[DEFAULT_LOCALE, { locale: DEFAULT_LOCALE, direction: defaultDirection }]]),
  numberFormat: new Intl.NumberFormat(DEFAULT_LOCALE),
}

export const localeState = $state({
  locale: DEFAULT_LOCALE,
  direction: defaultDirection,
})

export function applyLocaleState(locale, catalog) {
  localeState.locale = locale
  localeState.direction = catalog['_meta.direction']
  localeRuntimeState.numberFormat = new Intl.NumberFormat(locale)
  if (typeof document !== 'undefined') {
    document.documentElement.lang = locale
    document.documentElement.dir = localeState.direction
  }
}

export async function initializeLocale(preference = '') {
  const runtime = await import('./locale-runtime.js')
  return runtime.initializeLocale(preference)
}

export async function setLocale(preference, options) {
  const runtime = await import('./locale-runtime.js')
  return runtime.setLocale(preference, options)
}


export function t(key, values = {}) {
  const catalog = localeRuntimeState.catalogs.get(localeState.locale) || english
  const message = catalog[key] ?? english[key]
  if (message === undefined) return key
  return message.replace(/\{([A-Za-z][A-Za-z0-9_]*)\}/g, (_, name) => {
    const value = values[name]
    if (value == null) return `{${name}}`
    return typeof value === 'number' ? localeRuntimeState.numberFormat.format(value) : value
  })
}


export async function registerCatalogForTest(locale, catalog) {
  const runtime = await import('./locale-runtime.js')
  return runtime.registerCatalogForTest(locale, catalog)
}
