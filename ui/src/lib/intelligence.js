// SPDX-License-Identifier: AGPL-3.0-or-later
import { formatDate } from './format.js'
import { t } from './i18n.svelte.js'


export const DATE_ROLES = Object.freeze([
  'issued', 'due', 'start', 'end', 'expiry', 'renewal', 'service', 'other',
])

export function intelligenceRoleLabel(value) {
  const role = String(value || '').trim()
  return t(`intelligence.role.${DATE_ROLES.includes(role) ? role : 'date'}`)
}

export function intelligenceDateValue(candidate) {
  return candidate?.value?.date || candidate?.sort_value || ''
}

// The compact month grid shows documents, not every semantic role attached to
// the same date. Keep the first API-ordered event for each document and day.
// Month/year values contain sorting placeholders, not exact dates.
export function groupCalendarEvents(items) {
  const grouped = new Map()
  const seen = new Set()
  for (const event of items) {
    const precision = event.value?.precision
    if (precision === 'month' || precision === 'year') continue
    const date = intelligenceDateValue(event)
    const identity = `${date}:${event.document_id ?? `event-${event.id}`}`
    if (seen.has(identity)) continue
    seen.add(identity)
    if (!grouped.has(date)) grouped.set(date, [])
    grouped.get(date).push(event)
  }
  return grouped
}

export function formatArchiveDate(value, options = { year: 'numeric', month: 'short', day: 'numeric' }) {
  if (!value) return t('intelligence.unknownDate')
  const date = new Date(`${value}T00:00:00`)
  return formatDate(date, options)
}

export function formatIntelligenceDate(candidate, options) {
  const precision = candidate?.value?.precision
  if (precision === 'month') options = { year: 'numeric', month: 'short' }
  else if (precision === 'year') options = { year: 'numeric' }
  return formatArchiveDate(intelligenceDateValue(candidate), options)
}

export function formatIntelligenceValue(candidate) {
  if (candidate?.type === 'date') {
    return t('intelligence.dateValue', {
      date: formatIntelligenceDate(candidate),
      role: intelligenceRoleLabel(candidate.role),
    })
  }
  return candidate?.raw_text || candidate?.sort_value || candidate?.type || t('intelligence.extractedFact')
}

// A score never selects or authorizes a candidate. Only the live review
// projection can make a known date available for explicit selection.
export function canReviewDate(candidate) {
  return candidate?.type === 'date' && DATE_ROLES.includes(candidate.role) &&
    candidate.status === 'pending' && candidate.source_current === true &&
    ['important_fact', 'low_confidence', 'review_first'].includes(candidate.reason)
}

export function reviewFailure(error) {
  const code = typeof error === 'string' ? error : error?.code
  if (['stale_source', 'stale_proposal', 'source_changed'].includes(code)) {
    return t('intelligence.reviewFailure.staleSource')
  }
  if (['conflict', 'review_conflict', 'human_changed', 'already_resolved'].includes(code)) {
    return t('intelligence.reviewFailure.conflict')
  }
  if (['invalid_evidence', 'evidence_mismatch', 'evidence_missing'].includes(code)) {
    return t('intelligence.reviewFailure.evidence')
  }
  if (['hidden', 'reveal_required', 'sensitive_reveal_required'].includes(code)) {
    return t('intelligence.reviewFailure.hidden')
  }
  if (['forbidden', 'unauthorized', 'interactive_required', 'interactive_session_required'].includes(code) || [401, 403].includes(error?.status)) {
    return t('intelligence.reviewFailure.forbidden')
  }
  if (['not_found', 'no_task', 'system_unavailable', 'source_unavailable'].includes(code) || error?.status === 404) {
    return t('intelligence.reviewFailure.notFound')
  }
  if (['unsupported', 'unknown_action', 'unknown_handler'].includes(code)) {
    return t('intelligence.reviewFailure.unsupported')
  }
  if (error?.name === 'TypeError' || error?.status === 503) {
    return t('intelligence.reviewFailure.offline')
  }
  return t('intelligence.reviewFailure.unknown')
}
