<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "./i18n.svelte.js";
  let {
    value = $bindable({ visibility: 'network', expiry: '0', password: '' }),
    publicAvailable = false,
    highSensitivity = false,
    busy = false,
    submitLabel = translate("shareLinkForm.createCopy"),
    onSubmit,
  } = $props()

  const NETWORK_EXPIRIES = [
    ['0', translate("shareLinkForm.neverExpires")],
    ['86400', translate("shareLinkForm.1Day")],
    ['604800', translate("shareLinkForm.7Days")],
    ['2592000', translate("shareLinkForm.30Days")],
    ['7776000', translate("shareLinkForm.90Days")],
    ['31536000', translate("shareLinkForm.365Days")],
  ]
  const PUBLIC_EXPIRIES = NETWORK_EXPIRIES.slice(1)
  const expiries = $derived(value.visibility === 'public' ? PUBLIC_EXPIRIES : NETWORK_EXPIRIES)

  function chooseVisibility(visibility) {
    value = {
      ...value,
      visibility,
      expiry: visibility === 'public' && value.expiry === '0' ? '604800' : value.expiry,
    }
  }

  function submit(event) {
    event.preventDefault()
    if (!busy) onSubmit?.()
  }
</script>

<form class="share-form" onsubmit={submit}>
  <fieldset class="share-surfaces" disabled={busy}>
    <legend>{translate("shareLinkForm.whoCanUseThisLink")}</legend>
    <label class:chosen={value.visibility === 'network'}>
      <input type="radio" name="share-visibility" value="network"
             checked={value.visibility === 'network'} onchange={() => chooseVisibility('network')} />
      <span><strong>{translate("shareLinkForm.privateNetwork")}</strong><small>{translate("shareLinkForm.usesYourNormalSuchiAddressItIs")}</small></span>
    </label>
    {#if publicAvailable}
      <label class:chosen={value.visibility === 'public'}>
        <input type="radio" name="share-visibility" value="public"
               checked={value.visibility === 'public'} onchange={() => chooseVisibility('public')} />
        <span><strong>{translate("shareLinkForm.publicInternet")}</strong><small>{translate("shareLinkForm.anyoneWithTheLinkCanDownloadFrom")}</small></span>
      </label>
    {/if}
  </fieldset>

  <div class="share-fields">
    <label>
      <span>{translate("shareLinkForm.expiry")}</span>
      <select class="input" bind:value={value.expiry} aria-label={translate("shareLinkForm.linkExpiry")} disabled={busy}>
        {#each expiries as [seconds, label]}<option value={seconds}>{label}</option>{/each}
      </select>
    </label>
    <label class="password-field">
      <span>{translate("shareLinkForm.password")} <small>{translate("shareLinkForm.optional")}</small></span>
      <input class="input" type="password" bind:value={value.password}
             aria-label={translate("shareLinkForm.sharePassword")} autocomplete="new-password" disabled={busy} />
    </label>
  </div>

  {#if highSensitivity}
    <p class="share-advice">{translate("shareLinkForm.addAPasswordForConfidentialOrRestricted")}</p>
  {/if}

  <button class="btn primary sm" type="submit" disabled={busy}>
    {busy ? translate("shareLinkForm.creating") : submitLabel}
  </button>
</form>

<style>
  .share-form { display:grid; gap:12px; }
  .share-surfaces { display:grid; gap:7px; margin:0; padding:0; border:0; }
  .share-surfaces legend { margin-bottom:6px; color:var(--muted); font-size:.76rem; font-weight:650; }
  .share-surfaces label {
    display:grid; grid-template-columns:auto 1fr; gap:9px; align-items:start;
    padding:10px; border:1px solid var(--line); border-radius:9px; cursor:pointer;
  }
  .share-surfaces label.chosen { border-color:var(--accent); background:var(--tint); }
  .share-surfaces input { margin-top:3px; }
  .share-surfaces span { display:grid; gap:2px; }
  .share-surfaces strong { font-size:.82rem; }
  .share-surfaces small { color:var(--muted); font-size:.74rem; line-height:1.4; }
  .share-fields { display:grid; grid-template-columns:minmax(120px,.7fr) minmax(170px,1fr); gap:9px; }
  .share-fields label { display:grid; gap:5px; color:var(--muted); font-size:.72rem; font-weight:650; }
  .share-fields label > span small { color:var(--faint); font-weight:500; }
  .share-advice { margin:0; color:var(--warn); font-size:.76rem; }
  .share-form > button { justify-self:start; }
  @media (max-width: 540px) {
    .share-fields { grid-template-columns:1fr; }
    .share-form > button { width:100%; justify-content:center; }
  }
</style>
