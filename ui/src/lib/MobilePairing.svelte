<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "./i18n.svelte.js";
  import { onMount, onDestroy } from 'svelte'
  import { createMobilePairing, cancelMobilePairing } from './api.js'
  import { copyText } from './clipboard.js'
  import { session } from './session.svelte.js'
  import { captureScope, scopeCurrent } from './systems.svelte.js'
  import Icon from './Icon.svelte'

  let { onClose, notify } = $props()
  let dialog
  let deviceName = $state(translate("mobilePairing.suchiMobile"))
  let pairing = $state(null)
  let busy = $state(false)
  let error = $state('')
  let copyMessage = $state('')
  let now = $state(Date.now())
  const user = session.user
  const scope = captureScope()
  let disposed = false
  const remaining = $derived(Math.max(0, Math.ceil((pairing?.expires_at || 0) - now / 1000)))

  async function cancelCode(code) {
    if (!code || session.user !== user) return
    try { await cancelMobilePairing(code, scope.code) }
    catch {
      if (session.user === user && scopeCurrent(scope)) notify?.(translate("mobilePairing.couldNotCancelThePairingCodeItWill"))
    }
  }

  async function generate(event) {
    event.preventDefault()
    if (disposed || session.user !== user || !scopeCurrent(scope) || busy || !deviceName.trim()) return
    busy = true
    error = ''
    copyMessage = ''
    try {
      const result = await createMobilePairing(deviceName.trim(), scope.code)
      if (disposed || session.user !== user || !scopeCurrent(scope)) {
        await cancelCode(result?.code)
        return
      }
      pairing = result
      now = Date.now()
    } catch (ex) {
      if (!disposed && session.user === user && scopeCurrent(scope)) error = ex.message || translate("mobilePairing.couldNotCreateAPairingCodeTryAgain")
    } finally { busy = false }
  }

  async function copyLink() {
    if (disposed || session.user !== user || !scopeCurrent(scope)) return
    const link = pairing?.pairing_url
    if (!link || !remaining) return
    const copied = await copyText(link)
    if (disposed || session.user !== user || !scopeCurrent(scope) || pairing?.pairing_url !== link) return
    copyMessage = copied ? translate("mobilePairing.pairingLinkCopied") : translate("mobilePairing.selectThePairingLinkBelowAndCopyIt")
  }

  $effect(() => {
    if (session.user !== user || !scopeCurrent(scope)) {
      void cancelCode(pairing?.code)
      pairing = null
      copyMessage = ''
      error = ''
      onClose?.()
    }
  })

  onMount(() => {
    dialog.showModal()
    const timer = setInterval(() => now = Date.now(), 1000)
    return () => clearInterval(timer)
  })
  onDestroy(() => {
    disposed = true
    const code = pairing?.code
    pairing = null
    copyMessage = ''
    void cancelCode(code)
    dialog?.close()
  })
</script>

<dialog bind:this={dialog} class="modal pairing-dialog" aria-labelledby="pairing-title"
        oncancel={(event) => { event.preventDefault(); onClose?.() }}>
  <div class="modal-head">
    <h3 id="pairing-title">{translate("mobilePairing.pairMobileApp")}</h3>
    <button class="btn sm" onclick={() => onClose?.()} aria-label={translate("mobilePairing.closeMobilePairing")}><Icon name="x" size={13} /></button>
  </div>
  <p>{translate("mobilePairing.inTheSuchiAppChoose")} <strong>{translate("mobilePairing.scanQrCode")}</strong>{translate("mobilePairing.thenConfirmThisArchiveSAddressOn")}</p>
  <form onsubmit={generate}>
    <div class="field">
      <label for="pairing-device">{translate("mobilePairing.deviceName")}</label>
      <input id="pairing-device" class="input" bind:value={deviceName} maxlength="64" required disabled={busy} />
    </div>
    <button class="btn primary" disabled={busy || !deviceName.trim()}>
      {busy ? translate("mobilePairing.creating") : pairing ? translate("mobilePairing.generateNewCode") : translate("mobilePairing.generateQrCode")}
    </button>
  </form>
  {#if error}<p class="err" role="alert">{error}</p>{/if}
  {#if pairing && remaining > 0}
    <div class="pairing-code" aria-busy={busy}>
      <img src={pairing.qr_data_url} alt={translate("mobilePairing.scanThisQrCodeInTheSuchi")} width="320" height="320" />
      <p class="expiry">{translate("mobilePairing.expiresInValueValueUsableOnce", { "remaining": Math.floor(remaining / 60), "padStart": String(remaining % 60).padStart(2, '0') })}</p>
    </div>
    <div class="field">
      <label for="pairing-link">{translate("mobilePairing.canTScanPasteThisPairingLink")}</label>
      <input id="pairing-link" class="input mono" readonly value={pairing.pairing_url} onclick={(event) => event.currentTarget.select()} />
      <button class="btn sm" type="button" disabled={busy} onclick={copyLink}>{translate("mobilePairing.copyPairingLink")}</button>
      {#if copyMessage}<p class="copy-message" role="status">{copyMessage}</p>{/if}
    </div>
  {:else if pairing}
    <p role="status">{translate("mobilePairing.thisCodeHasExpiredGenerateANew")}</p>
  {/if}
  <p class="privacy">{translate("mobilePairing.anyoneWithThisCodeCanConnectAs")}</p>
  <p class="fallback">{translate("mobilePairing.youCanAlsoEnterTheServerAddress")}</p>
  <button class="btn" onclick={() => onClose?.()}>{translate("mobilePairing.done")}</button>
</dialog>

<style>
  .pairing-dialog { width:min(480px,94vw);margin:auto;color:var(--ink);max-height:90dvh;overflow:auto }
  dialog::backdrop { background:rgba(0,0,0,.45) }
  p { font-size:.86rem;line-height:1.5 }
  form { display:flex;gap:12px;align-items:end;flex-wrap:wrap }
  form .field { flex:1;min-width:160px;margin:0 }
  form button { min-height:37px }
  .pairing-code { text-align:center;margin:18px 0 }
  .pairing-code img { display:block;width:min(100%,320px);height:auto;margin:auto;background:white;border-radius:8px }
  .expiry { color:var(--muted);margin:8px 0 }
  .field button { align-self:flex-start }
  .mono { font-size:.76rem;min-width:0 }
  .privacy, .fallback { color:var(--muted);font-size:.8rem }
  .copy-message { margin:0 }
</style>
