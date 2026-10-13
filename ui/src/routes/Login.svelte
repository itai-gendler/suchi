<!-- SPDX-License-Identifier: AGPL-3.0-or-later -->
<script>
  import { t as translate } from "../lib/i18n.svelte.js";
  import { onMount } from 'svelte'
  import { login } from '../lib/api.js'
  import { getLoginPath, usesExternalLogin } from '../lib/login.js'
  import { session, refreshSession } from '../lib/session.svelte.js'
  import BrandMark from '../lib/BrandMark.svelte'

  let { onSignedIn } = $props()
  let email = $state('')
  let password = $state('')
  let err = $state('')
  let busy = $state(false)
  const externalLogin = usesExternalLogin()

  onMount(() => {
    if (externalLogin) location.assign(getLoginPath())
  })

  async function submit(e) {
    e.preventDefault()
    err = ''; busy = true
    try {
      await login(email.trim(), password)
      await refreshSession()
      if (!session.user) throw new Error(translate("login.signInDidNotStickCheckTheServer"))
      onSignedIn?.()
    } catch (ex) {
      err = ex.status === 401 ? translate("login.emailOrPasswordDidNotMatch") : (ex.message || translate("login.couldNotSignIn"))
    } finally { busy = false }
  }
</script>

{#if !externalLogin}
  <div class="login-wrap">
    <form class="card login-card" onsubmit={submit}>
      <div class="brand">
        <BrandMark />
        <b style="font-size:1.2rem">suchi</b>
      </div>
      {#if err}<div class="err">{err}</div>{/if}
      <div class="field">
        <label for="email">{translate("login.email")}</label>
        <input id="email" class="input" type="email" bind:value={email} autocomplete="username" required />
      </div>
      <div class="field">
        <label for="pw">{translate("login.password")}</label>
        <input id="pw" class="input" type="password" bind:value={password} autocomplete="current-password" required />
      </div>
      <button class="btn primary" style="width:100%;justify-content:center" disabled={busy}>
        {busy ? translate("login.signingIn") : translate("login.signIn")}
      </button>
    </form>
  </div>
{/if}
