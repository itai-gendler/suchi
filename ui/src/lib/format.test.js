// SPDX-License-Identifier: AGPL-3.0-or-later

import assert from 'node:assert/strict'
import test from 'node:test'
import {
  SENSITIVITY_OPTIONS, customFieldLabel, isHighSensitivity, sensDot, sensitivityLabel,
} from './format.js'

test('keeps the sensitivity vocabulary and privacy posture aligned', () => {
  assert.deepEqual(
    SENSITIVITY_OPTIONS.map((option) => option.value),
    ['public', 'internal', 'confidential', 'restricted'],
  )
  assert.equal(isHighSensitivity('confidential'), true)
  assert.equal(isHighSensitivity('restricted'), true)
  assert.equal(isHighSensitivity('internal'), false)
  assert.equal(sensDot('restricted'), 'danger')
  assert.equal(sensitivityLabel('restricted'), 'Restricted')
  assert.equal(sensitivityLabel(''), 'Unset')
})

test('labels extraction-managed invoice fields without changing ordinary field names', () => {
  const labels = {
    invoice_number: 'Invoice number',
    invoice_date: 'Invoice date',
    invoice_currency: 'Invoice currency',
    invoice_total_gross: 'Invoice total (gross)',
    invoice_total_net: 'Invoice total (net)',
    invoice_seller: 'Invoice seller',
    invoice_seller_vat: 'Seller VAT number',
    invoice_buyer: 'Invoice buyer',
  }

  for (const [name, label] of Object.entries(labels)) {
    assert.equal(customFieldLabel({ name, managed_by: 'zugferd' }), label)
  }
  assert.equal(customFieldLabel({ name: 'Warranty end' }), 'Warranty end')
  assert.equal(customFieldLabel({ name: 'plugin_field', managed_by: 'plugin' }), 'plugin_field')
})
