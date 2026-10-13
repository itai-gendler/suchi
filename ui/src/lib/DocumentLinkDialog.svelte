<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "./i18n.svelte.js";
  import Icon from './Icon.svelte'
  import LinkQR from './LinkQR.svelte'
  import { copyText } from './clipboard.js'
  import { captureScope, scopedHash } from './systems.svelte.js'

  let { id, onClose } = $props()
  let dialog
  let copied = $state('')
  const scope = captureScope()
  const url = $derived(new URL(`/app/${scopedHash(`#/doc/${id}`, scope.code)}`, location.origin).href)
  const loopback = ['localhost', '127.0.0.1', '[::1]'].includes(location.hostname)

  $effect(() => {
    dialog.showModal()
    return () => dialog.close()
  })

  async function copy() {
    copied = await copyText(url) ? translate("documentLinkDialog.linkCopied") : translate("documentLinkDialog.selectTheLinkAndCopyItManually")
  }
</script>

<dialog bind:this={dialog} class="modal" aria-labelledby="document-link-title" onclose={onClose}>
  <div class="modal-head">
    <h3 id="document-link-title">{translate("documentLinkDialog.openOnMyPhone")}</h3>
    <button class="btn sm" onclick={() => dialog.close()} aria-label={translate("documentLinkDialog.closeDocumentLink")}><Icon name="x" size={13} /></button>
  </div>
  <p class="sub">{translate("documentLinkDialog.scanWithYourPhoneSCameraSign")}</p>
  {#if loopback}
    <p class="sub">{translate("documentLinkDialog.thisAddressOnlyWorksOnThisComputer")}</p>
  {:else}
    <p class="sub">{translate("documentLinkDialog.yourPhoneNeedsAccessToThisServer")}</p>
  {/if}
  <div class="toolbar">
    <input class="input mono" style="flex:1;min-width:0" aria-label={translate("documentLinkDialog.documentLink")} readonly value={url} onclick={(event) => event.currentTarget.select()} />
    <button class="btn sm" onclick={copy}>{translate("documentLinkDialog.copyLink")}</button>
  </div>
  <p class="sub" role="status">{copied}</p>
  <LinkQR {url} open />
</dialog>

<style>
  dialog { margin: auto; width: min(460px, 94vw); color: var(--ink); }
  dialog::backdrop { background: rgba(0, 0, 0, .45); }
  p { color: var(--muted); }
</style>
