<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "./i18n.svelte.js";
  import { scopedHash as filingHref } from './systems.svelte.js'
  import Icon from './Icon.svelte'

  let { placement = 'side', error = false, onRetry, onContinue } = $props()
</script>

<aside class="setup-reminder" class:setup-reminder-side={placement === 'side'}
       class:setup-reminder-mobile={placement === 'mobile'}
       class:setup-reminder-settings={placement === 'settings'} aria-label={translate("setupReminder.archiveSetup")}>
  <div class="setup-reminder-head">
    <span class="setup-reminder-icon"><Icon name="settings" size={16} /></span>
    <div>
      <b>{error ? translate("setupReminder.archiveSetupStatusUnavailable") : translate("setupReminder.chooseYourFilingTree")}</b>
      <span>{error ? translate("setupReminder.suchiCouldNotVerifyWhetherSetupIsComplete") : translate("setupReminder.chooseHowDocumentsAreOrganizedToFinishArchive")}</span>
    </div>
  </div>
  <div class="setup-reminder-actions">
    {#if error}<button class="btn sm" onclick={() => onRetry?.()}>{translate("common.retry")}</button>{/if}
    <a role="button" class="btn primary sm" href={filingHref("#/settings?tab=archive&section=filing-tree")} onclick={onContinue}>{translate("setupReminder.continueSetup")}</a>
  </div>
</aside>
