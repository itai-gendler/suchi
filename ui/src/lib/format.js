// SPDX-License-Identifier: AGPL-3.0-or-later

import { localeState, t } from './i18n.svelte.js'

const formatterCache = new Map()

export function formatDate(value, options = { year: 'numeric', month: 'short', day: 'numeric' }) {
  const date = value instanceof Date ? value : new Date(value)
  if (Number.isNaN(date.getTime())) return t('common.invalidDate')
  return formatter(`date:${JSON.stringify(options)}`, () => new Intl.DateTimeFormat(localeState.locale, options)).format(date)
}

export function formatNumber(value, options = {}) {
  return formatter(`number:${JSON.stringify(options)}`, () => new Intl.NumberFormat(localeState.locale, options)).format(value)
}

export function compareLocalized(left, right) {
  return formatter('collator', () => new Intl.Collator(localeState.locale, { numeric: true, sensitivity: 'base' })).compare(left, right)
}

export function plural(key, count, values = {}) {
  const category = formatter('plural', () => new Intl.PluralRules(localeState.locale)).select(count)
  const messageKey = `${key}.${category}`
  const parameters = { ...values, count }
  const message = t(messageKey, parameters)
  return message === messageKey ? t(`${key}.other`, parameters) : message
}

export function formatList(values, options = {}) {
  return formatter(`list:${JSON.stringify(options)}`, () => new Intl.ListFormat(localeState.locale, options)).format(values)
}

function formatter(key, create) {
  const fullKey = `${localeState.locale}:${key}`
  if (!formatterCache.has(fullKey)) formatterCache.set(fullKey, create())
  return formatterCache.get(fullKey)
}

export function fmtDate(unix) {
  if (!unix) return ''
  return formatDate(new Date(unix * 1000))
}
export function fmtBytes(n) {
  if (!n && n !== 0) return ''
  const units = ['B', 'KB', 'MB', 'GB']
  let i = 0
  while (n >= 1024 && i < units.length - 1) { n /= 1024; i++ }
  return `${formatNumber(n, { maximumFractionDigits: n < 10 && i > 0 ? 1 : 0 })} ${units[i]}`
}

export const SENSITIVITY_OPTIONS = Object.freeze([
  Object.freeze({ value: 'public', get label() { return t('sensitivity.public') } }),
  Object.freeze({ value: 'internal', get label() { return t('sensitivity.internal') } }),
  Object.freeze({ value: 'confidential', get label() { return t('sensitivity.confidential') } }),
  Object.freeze({ value: 'restricted', get label() { return t('sensitivity.restricted') } }),
])

export function isHighSensitivity(value) {
  return value === 'confidential' || value === 'restricted'
}

export function sensitivityLabel(value) {
  return SENSITIVITY_OPTIONS.find((option) => option.value === value)?.label || t('common.unset')
}

export function sensDot(s) {
  return { confidential: 'danger', restricted: 'danger', internal: 'warn', public: 'ok' }[s] || ''
}
