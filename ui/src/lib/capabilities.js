// SPDX-License-Identifier: AGPL-3.0-or-later

export const USER_CAPABILITIES = Object.freeze([
  Object.freeze({
    key: 'archive_chat',
    label: 'Ask your documents',
    description: 'Ask questions across documents they can read, with links back to each document.',
  }),
  Object.freeze({
    key: 'archive_intelligence',
    label: 'Manage document dates',
    description: 'Extract dates, review uncertain results, and use Calendar dates in answers and search.',
  }),
  Object.freeze({
    key: 'mailboxes',
    label: 'Manage mailboxes',
    description: 'Connect and manage their own mail intake.',
  }),
  Object.freeze({
    key: 'share_links',
    label: 'Create share links',
    description: 'Share documents using revocable links.',
  }),
  Object.freeze({
    key: 'share_views',
    label: 'Share saved views',
    description: 'Publish saved views to others in the same filing system.',
  }),
])

export function hasCapability(user, capability) {
  return user?.role === 'admin' || (user?.capabilities || []).includes(capability)
}
