<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { scopedHash as filingHref } from '../lib/systems.svelte.js'
  import { onDestroy } from 'svelte'
  import { listSavedViews } from '../lib/api.js'
  import { documentListHash, parseSavedViewFilters } from '../lib/documentFilters.js'
  import { fmtDate, sensDot } from '../lib/format.js'
  import Icon from '../lib/Icon.svelte'

  let {
    st, statsError = '', inboxCategory, taxonomyLoaded = false, taxonomyError = '', recent,
    recentError = '', onRetryRecent, setupNeeded = false, setupError = false, onRetrySetup,
  } = $props()
  const inboxCount = $derived(st?.inbox_count ?? 0)
  const pending = $derived((st?.pending_approvals ?? 0) + (st?.pending_intelligence ?? 0))
  const dead = $derived(st?.dead_jobs ?? 0)

  const total = $derived(st?.documents_total ?? null)
  const dashboardViewLimit = 4
  let views = $state([])
  let viewsLoading = $state(true)
  let viewsError = $state('')
  let viewTotal = $state(0)
  let loadVersion = 0
  let activeController

  const onboarding = $derived.by(() => {
    if (setupNeeded) {
      return {
        step: 'Step 1 of 3',
        icon: 'settings',
        title: setupError ? 'Check archive setup' : 'Choose your filing tree',
        body: setupError
          ? 'Suchi could not confirm whether setup is complete. Retry, or open the filing tree settings.'
          : 'Choose how documents will be organized before adding your first one.',
        action: 'Continue setup',
        href: '#/settings?tab=archive&section=filing-tree',
      }
    }
    if (!st) return null
    if (total === 0) {
      return {
        step: 'Step 2 of 3',
        icon: 'upload',
        title: 'Add your first document',
        body: 'Start with one representative document. Suchi keeps the original and prepares it for filing.',
        action: 'Upload a document',
        href: '#/upload',
      }
    }
    if (total === 1 && (inboxCount > 0 || pending > 0 || dead > 0)) {
      if (pending > 0) {
        return {
          step: 'Step 3 of 3',
          icon: 'tasks',
          title: 'Finish your first document',
          body: inboxCount > 0
            ? 'Review the waiting suggestion or decision, then file the document from Inbox.'
            : 'A suggestion or decision is waiting for your review.',
          action: 'Review it',
          href: '#/tasks',
          secondaryAction: inboxCount > 0 ? 'Open Inbox' : '',
          secondaryHref: '#/inbox',
        }
      }
      if (inboxCount > 0) {
        return {
          step: 'Step 3 of 3',
          icon: 'inbox',
          title: 'Finish filing your first document',
          body: 'The document is ready for a filing decision in Inbox.',
          action: 'Open Inbox',
          href: '#/inbox',
        }
      }
      return {
        step: 'Step 3 of 3',
        icon: 'zap',
        title: 'Finish processing your first document',
        body: 'Processing stopped before the document was ready. Open the issue to retry or dismiss it.',
        action: 'Open processing issues',
        href: '#/tasks',
      }
    }
    return null
  })

  async function load() {
    const version = ++loadVersion
    activeController?.abort()
    const controller = new AbortController()
    activeController = controller
    viewsLoading = true
    viewsError = ''
    try {
      const res = await listSavedViews({ include: 'shared' }, controller.signal)
      if (version !== loadVersion) return
      const raw = res?.results || res || []
      viewTotal = res?.count ?? raw.length
      const sorted = raw.map(v => ({
        ...v,
        filters: parseSavedViewFilters(v.filter_json),
      })).sort((a, b) => a.position - b.position)

      const selected = sorted.slice(0, dashboardViewLimit)
      const shared = sorted.find((view) => view.owner_id)
      if (shared && !selected.includes(shared)) {
        if (selected.length === dashboardViewLimit) selected[selected.length - 1] = shared
        else selected.push(shared)
      }

      views = selected
    } catch (ex) {
      if (version !== loadVersion) return
      views = []
      viewsError = ex.message || 'Could not load views.'
    } finally {
      if (activeController === controller) activeController = undefined
      if (version === loadVersion) viewsLoading = false
    }
  }

  function filterSummary(f) {
    const parts = []
    if (f.q) parts.push(`“${f.q}”`)
    if (f.tags__id__in) parts.push('tag')
    if (f.correspondents__id__in) parts.push('correspondent')
    if (f.jd_category_id) parts.push('filing category')
    if (f.sensitivity) parts.push(f.sensitivity)
    return parts.join(' · ') || 'all documents'
  }

  load()
  onDestroy(() => {
    loadVersion++
    activeController?.abort()
  })
</script>

  {#if onboarding}
    <section class="onboarding-card card" aria-labelledby="onboarding-heading">
      <span class="onboarding-icon"><Icon name={onboarding.icon} size={20} /></span>
      <div class="onboarding-copy">
        <span class="onboarding-step">Getting started · {onboarding.step}</span>
        <h2 id="onboarding-heading">{onboarding.title}</h2>
        <p>{onboarding.body}</p>
        <div class="onboarding-actions">
          <a class="btn primary" href={filingHref(onboarding.href)}>{onboarding.action}</a>
          {#if onboarding.secondaryAction}
            <a class="btn" href={filingHref(onboarding.secondaryHref)}>{onboarding.secondaryAction}</a>
          {/if}
          {#if setupError}
            <button class="btn" type="button" onclick={() => onRetrySetup?.()}>Retry setup status</button>
          {/if}
        </div>
      </div>
    </section>
  {/if}

<div class="metrics">
  <a class="metric card" href={filingHref("#/documents")}>
    <span class="m-label"><Icon name="docs" size={14} /> Total documents</span>
    <span class="m-value">{total ?? '—'}</span>
    <span class="m-sub">{st?.ingested_7d ? `${st.ingested_7d} added in the last 7 days` : 'across the whole archive'}</span>
  </a>
  <a class="metric card" href={filingHref("#/inbox")} class:attn={inboxCount > 0}>
    <span class="m-label"><Icon name="inbox" size={14} /> Inbox</span>
    <span class="m-value">{inboxCategory ? inboxCount : '—'}</span>
    <span class="m-sub">{inboxCategory
      ? (inboxCount > 0 ? 'waiting to be filed' : 'everything is filed')
      : (taxonomyLoaded ? (taxonomyError || 'inbox unavailable') : 'loading archive structure')}</span>
  </a>
  <a class="metric card" href={filingHref("#/tasks")} class:attn={pending > 0}>
    <span class="m-label"><Icon name="tasks" size={14} /> Approvals</span>
    <span class="m-value">{st ? pending : '—'}</span>
    <span class="m-sub">{st ? (pending > 0 ? 'waiting on you' : 'none pending') : (statsError || 'loading archive status')}</span>
  </a>
  <a class="metric card" href={filingHref("#/tasks")} class:bad={dead > 0}>
    <span class="m-label"><Icon name="zap" size={14} /> Failed jobs</span>
    <span class="m-value">{st ? dead : '—'}</span>
    <span class="m-sub">{st ? (dead > 0 ? 'dead-lettered — needs attention' : 'pipeline healthy') : (statsError || 'loading archive status')}</span>
  </a>
</div>

<div class="dash-grid">
  <div>
    <div class="dash-head">
      <h3>Recently added</h3>
      <a class="btn sm" href={filingHref("#/documents")}>All documents</a>
    </div>
    {#if recent === undefined}
      <div class="index" aria-label="Loading recent documents">
        {#each Array(3) as _}<div class="irow"><div class="skel" style="width:68%"></div></div>{/each}
      </div>
    {:else if recentError}
      <div class="card" style="display:flex;align-items:center;justify-content:space-between;gap:12px;color:var(--muted);font-size:.86rem">
        <span>{recentError}</span><button class="btn sm" onclick={() => onRetryRecent?.({ background: false })}>Retry</button>
      </div>
    {:else if recent.length}
      <div class="index">
        {#each recent as d (d.id)}
          <a class="irow" href={filingHref(`#/doc/${d.id}`)}>
            <span class="dot {sensDot(d.sensitivity)}" class:accent={!d.sensitivity}></span>
            {#if d.jd_address || d.jd_category_code}<span class="chip" title={d.jd_category_name}>{d.jd_address || d.jd_category_code}</span>{/if}
            <span class="title grow">{d.title || `Document #${d.id}`}</span>
            <span class="sub">{fmtDate(d.created_at)}</span>
          </a>
        {/each}
      </div>
    {:else}
      <div class="empty" style="padding:36px 20px">
        <Icon name="docs" size={44} />
        <b>Nothing here yet.</b>
        <span><a href={filingHref("#/upload")}>Upload the first document</a>, point a watched folder at it, or forward an email.</span>
      </div>
    {/if}
  </div>

  <div>
    <div class="dash-head">
      <h3>Views</h3>
      <a class="btn sm" href={filingHref("#/views?new=1")}><Icon name="plus" size={13} /> New view</a>
    </div>
    {#if viewsLoading}
      <div class="card" style="color:var(--muted);font-size:.86rem">Loading views…</div>
    {:else if viewsError}
      <div class="card" style="display:flex;align-items:center;justify-content:space-between;gap:12px;color:var(--muted);font-size:.86rem">
        <span>{viewsError}</span><button class="btn sm" onclick={load}>Retry</button>
      </div>
    {:else if views.length}
      <div class="views">
        {#each views as v (v.id)}
          <a class="card view" href={filingHref(documentListHash(v.filters))}>
            <span class="m-label">{v.name}</span>
            <span class="m-sub mono" style="font-size:.68rem">{filterSummary(v.filters)}</span>
          </a>
        {/each}
      </div>
      {#if viewTotal > views.length}<a class="sub" href={filingHref("#/views")}>View all {viewTotal} saved views</a>{/if}
    {:else}
      <div class="card" style="color:var(--muted);font-size:.86rem">
        No saved views yet.
      </div>
    {/if}
  </div>
</div>
