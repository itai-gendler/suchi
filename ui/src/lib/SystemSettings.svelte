<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "./i18n.svelte.js";
  import { onDestroy } from 'svelte'
  import { systems, captureScope, scopeCurrent } from './systems.svelte.js'
  import { refreshSystems } from './system-routing.js'
  import { systemMembers, putSystemMembers, renameSystem, adminListUsers } from './api.js'

  let { notify } = $props()
  const scope = captureScope()
  const code = scope.code
  let name = $state(systems.results.find(item => item.code === code)?.name || '')
  let users = $state([])
  let members = $state(new Set())
  let initial = $state(new Set())
  let page = $state(1)
  let busy = $state(false)
  let loaded = $state(false)
  let error = $state('')
  let disposed = false
  const pageSize = 20
  const pages = $derived(Math.max(1, Math.ceil(users.length / pageSize)))
  const visible = $derived(users.slice((page - 1) * pageSize, page * pageSize))
  const current = () => !disposed && scopeCurrent(scope)
  onDestroy(() => { disposed = true })

  async function load() {
    try {
      const [grants, directory] = await Promise.all([systemMembers(code), adminListUsers()])
      if (!current()) return
      // Memberships, not the visible directory page, own the complete PUT set.
      initial = new Set(grants.user_ids || [])
      members = new Set(initial)
      users = directory.results || []
      loaded = true
    } catch (ex) { if (current()) error = ex.message || translate("systemSettings.couldNotLoadSystemMemberships") }
  }
  void load()

  function toggle(id, checked) {
    const next = new Set(members)
    if (checked) next.add(id)
    else next.delete(id)
    members = next
  }
  async function saveMembers() {
    if (!current() || busy || !loaded) return
    busy = true; error = ''
    try {
      const result = await putSystemMembers(code, [...members].sort((a, b) => a - b))
      if (!current()) return
      members = new Set(result.user_ids)
      initial = new Set(result.user_ids)
      notify?.(translate("systemSettings.systemMembershipsSaved"))
    } catch (ex) { if (current()) error = ex.message || translate("systemSettings.couldNotSaveMemberships") }
    finally { if (current()) busy = false }
  }
  async function saveName(event) {
    event.preventDefault()
    if (!current() || busy || !name.trim()) return
    busy = true; error = ''
    try {
      await renameSystem(code, name.trim())
      if (!current()) return
      await refreshSystems()
      if (current()) notify?.(translate("systemSettings.systemNameSaved"))
    } catch (ex) { if (current()) error = ex.message || translate("systemSettings.couldNotSaveSystemName") }
    finally { if (current()) busy = false }
  }
</script>

<section class="card" aria-label={translate("systemSettings.filingSystemAccess")} style="padding:18px;margin-bottom:18px">
  <h3>{translate("systemSettings.filingSystemValue", { "code": code })}</h3>
  <p class="sub">{translate("systemSettings.theCodeIsPermanentMembersMayEnter")}</p>
  <form class="toolbar" onsubmit={saveName}>
    <label>{translate("systemSettings.systemName")} <input class="input" bind:value={name} maxlength="80" required disabled={busy} /></label>
    <button class="btn sm" disabled={busy || !name.trim()}>{translate("systemSettings.saveSystemName")}</button>
  </form>
  <h4>{translate("systemSettings.members")}</h4>
  <p class="sub">{translate("systemSettings.removingAccessRevokesThisSystemSShares")}</p>
  {#if error}<p class="err" role="alert">{error}</p>{/if}
  {#if loaded}
    {#each visible as user (user.id)}
      <label style="display:flex;gap:10px;align-items:center;padding:8px 0;overflow-wrap:anywhere">
        <input type="checkbox" aria-label={translate("systemSettings.systemAccessForValue", { "email": user.email })}
          checked={user.role === 'admin' || members.has(user.id)}
          disabled={busy || user.role === 'admin' || (user.disabled && !initial.has(user.id))}
          onchange={event => toggle(user.id, event.currentTarget.checked)} />
        <span>{user.display_name || user.email}{#if user.display_name} · {user.email}{/if}
          {#if user.role === 'admin'} {translate("systemSettings.administratorImplicit")}{:else if user.disabled} {translate("systemSettings.inactive")}{/if}</span>
      </label>
    {/each}
    <div class="toolbar">
      <button class="btn sm" disabled={page === 1} onclick={() => page--}>{translate("systemSettings.previousMembers")}</button>
      <span>{translate("systemSettings.pageValueOfValue", { "page": page, "pages": pages })}</span>
      <button class="btn sm" disabled={page === pages} onclick={() => page++}>{translate("systemSettings.nextMembers")}</button>
      <button class="btn primary sm" disabled={busy} onclick={saveMembers}>{translate("systemSettings.saveMemberships")}</button>
    </div>
  {:else if !error}<p role="status">{translate("systemSettings.loadingMemberships")}</p>{/if}
</section>
