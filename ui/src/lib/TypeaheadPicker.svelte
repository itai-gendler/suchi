<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { untrack } from 'svelte'
  import { localeState } from './i18n.svelte.js'
  import { compareLocalized } from './format.js'

  let {
    items = [],
    excludedValues = [],
    valueKey = 'id',
    disabled = false,
    label = '',
    menuId = 'typeahead-options',
    placeholder = '',
    emptyText = '',
    initialQuery = '',
    clearOnChoose = true,
    onQueryChange,
    onChoose,
  } = $props()
  let query = $state(untrack(() => initialQuery))
  let open = $state(false)
  let highlight = $state(-1)
  let working = $state(false)
  let box

  const matches = $derived.by(() => {
    const needle = query.trim().toLocaleLowerCase(localeState.locale)
    const excluded = new Set(excludedValues.map(value => String(value)))
    return items.filter(item => {
      const name = String(item?.name || '').toLocaleLowerCase(localeState.locale)
      const secondary = String(item?.slug || '').toLocaleLowerCase(localeState.locale)
      return !excluded.has(String(item?.[valueKey])) &&
        (!needle || name.includes(needle) || secondary.includes(needle))
    }).sort((a, b) => compareLocalized(a.name, b.name) || a.id - b.id).slice(0, 8)
  })

  function close() {
    open = false
    highlight = -1
  }

  function outside(event) {
    if (box && !box.contains(event.target)) close()
  }

  function queryChanged() {
    open = true
    highlight = -1
    onQueryChange?.(query)
  }

  async function choose(item) {
    if (working || disabled || !item) return
    working = true
    try {
      if (await onChoose(item)) {
        query = clearOnChoose ? '' : item.name
        close()
      }
    } finally {
      working = false
    }
  }

  function keydown(event) {
    if (event.key === 'Escape') {
      if (open) event.preventDefault()
      close()
      return
    }
    if (event.key === 'ArrowDown' || event.key === 'ArrowUp') {
      event.preventDefault()
      if (!open) {
        open = true
        highlight = -1
      }
      highlight = event.key === 'ArrowDown'
        ? Math.min(highlight + 1, matches.length - 1)
        : Math.max(highlight - 1, 0)
    } else if (event.key === 'Enter' && open) {
      event.preventDefault()
      if (highlight >= 0) choose(matches[highlight])
    }
  }
</script>

<svelte:window onmousedown={outside} />
<div class="typeahead-picker" bind:this={box}>
  <input class="input" type="search" bind:value={query} {disabled}
         autocomplete="off" role="combobox" aria-label={label} aria-autocomplete="list"
         aria-expanded={open} aria-controls={menuId}
         aria-activedescendant={open && highlight >= 0 && matches[highlight] ? `${menuId}-${highlight}` : undefined}
         {placeholder} onfocus={() => (open = true)}
         oninput={queryChanged} onkeydown={keydown} />
  {#if open && !disabled}
    <div class="omni-drop" id={menuId} role="listbox" aria-label={label}>
      {#each matches as item, i (item.id)}
        <button type="button" class="omni-row" class:hot={highlight === i}
                id={`${menuId}-${i}`} role="option" aria-selected={highlight === i}
                disabled={working} onmousedown={(event) => event.preventDefault()}
                onclick={() => choose(item)}>
          {item.name}
          {#if item.slug}<span class="secondary">{item.slug}</span>{/if}
        </button>
      {:else}
        <div class="empty">{emptyText}</div>
      {/each}
    </div>
  {/if}
</div>

<style>
  .typeahead-picker { position:relative; min-width:0; width:min(100%, 260px) }
  .typeahead-picker input { width:100%; min-width:0 }
  .omni-drop { left:0; right:auto; width:min(320px, 90vw); max-width:100% }
  .omni-row { min-width:0; justify-content:space-between; gap:12px }
  .secondary { color:var(--muted); overflow:hidden; text-overflow:ellipsis; white-space:nowrap; font-size:.75rem }
  .empty { padding:10px; color:var(--muted); font-size:.85rem }
</style>
