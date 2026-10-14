// SPDX-License-Identifier: AGPL-3.0-or-later

import english from '../../i18n/locales/en.json'

const german = Object.fromEntries(Object.entries(english).map(([key, value]) => [key, `[de] ${value}`]))

Object.assign(german, {
  '_meta.locale': 'de',
  '_meta.direction': 'ltr',
  'language.label': 'Sprache',
  'language.automatic': 'Automatisch',
  'common.cancel': 'Abbrechen',
  'common.save': 'Speichern',
  'apiError.requestFailed': 'Die Anfrage ist fehlgeschlagen (HTTP {status}).',
  'archiveChat.documentCount.one': '{count} Dokument',
  'archiveChat.documentCount.other': '{count} Dokumente',
  'mailAccountForm.ruleCountSummary.one': '{count} Regel · archivieren, wenn sie zutrifft',
  'mailAccountForm.ruleCountSummary.other': '{count} Regeln · archivieren, wenn eine Regel zutrifft',
  'apiError.server': 'Der Server konnte die Anfrage nicht abschließen.',
})

export default german
