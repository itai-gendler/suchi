<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "./i18n.svelte.js";
  import Icon from './Icon.svelte'

  let { state, expanded = false } = $props()
</script>

{#if state === 'decrypted'}
  <span class="unlock-status" role="img" aria-label={translate("documentUnlockStatus.passwordUnlocked")}
        title={translate("documentUnlockStatus.passwordUnlockedSuchiUnlockedAWorkingCopy")}>
    <Icon name="unlock" size={15} />
    {#if expanded}<span aria-hidden="true">{translate("documentUnlockStatus.passwordUnlocked")}</span>{/if}
  </span>
{/if}

<style>
  .unlock-status { display:inline-flex; align-items:center; gap:6px; flex-shrink:0; color:var(--muted); font-size:.8rem; vertical-align:middle }
</style>
