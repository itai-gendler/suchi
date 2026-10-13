// SPDX-License-Identifier: AGPL-3.0-or-later

import assert from 'node:assert/strict'
import test from 'node:test'
import german from '../testdata/locales/de.js'

globalThis.$state = value => value
const { localeState, registerCatalogForTest, setLocale, t } = await import('./i18n.svelte.js')
const { formatDate, formatNumber, plural } = await import('./format.js')
delete globalThis.$state

test('switches complete locale catalogs and locale-sensitive formatting', async () => {
  await registerCatalogForTest('de', german)
  await setLocale('de', { persist: false })
  try {
    assert.equal(localeState.locale, 'de')
    assert.equal(localeState.direction, 'ltr')
    assert.equal(t('language.label'), 'Sprache')
    assert.equal(t('apiError.requestFailed', { status: 418 }), 'Die Anfrage ist fehlgeschlagen (HTTP 418).')
    assert.equal(formatNumber(1234.5), '1.234,5')
    assert.equal(plural('archiveChat.documentCount', 1), '1 Dokument')
    assert.equal(plural('archiveChat.documentCount', 1234), '1.234 Dokumente')
    assert.equal(formatDate(new Date(2026, 8, 14), { day: 'numeric', month: 'long', year: 'numeric' }), '14. September 2026')
  } finally {
    await setLocale('en', { persist: false })
  }
})

test('rejects incomplete catalogs and placeholder drift', async () => {
  const incomplete = { ...german, '_meta.locale': 'de-x-missing' }
  delete incomplete['common.cancel']
  await assert.rejects(registerCatalogForTest('de-X-missing', incomplete), /incomplete/)

  const drifted = {
    ...german,
    '_meta.locale': 'de-x-drifted',
    'apiError.requestFailed': 'Fehler {code}',
  }
  await assert.rejects(registerCatalogForTest('de-X-drifted', drifted), /placeholders differ/)
})
