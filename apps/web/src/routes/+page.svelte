<script lang="ts">
	import type { PageData } from './$types'
	import ArrowRightIcon from '~icons/fa-solid/arrow-right'
	import FloatingLabel from '$lib/components/FloatingLabel/FloatingLabel.svelte'
	import { goto } from '$app/navigation'
	import { createForm } from 'felte'
	import { validator } from '@felte/validator-zod'
	import { loginSchema as schema } from './validate'
	import type { LoginSchema } from './validate'
	import { isAuthError } from '@supabase/supabase-js'

	export let data: PageData
	let { supabase } = data
	let loading = false
	let loginErr: string | undefined = undefined
	$: ({ supabase } = data)

	const { form, errors } = createForm<LoginSchema>({
		extend: validator({ schema }),
		debounced: { timeout: 500 },
		onSubmit: handleSignIn
	})

	async function handleSignIn({ email, password }: LoginSchema) {
		loading = true
		try {
			const res = await supabase.auth.signInWithPassword({
				email,
				password
			})
			if (res.error !== null && isAuthError(res.error)) {
				loginErr = res.error.message
				return
			}

			return goto('/admin/registrations')
		} catch (e) {
			loginErr = 'Lỗi do server, vui lòng thử lại sau ít phút!'
		} finally {
			loading = false
		}
	}
</script>

<div class="flex h-screen w-full items-center">
	<div class="m-auto w-full max-w-md">
		<h4 class="mb-8 text-center text-lg">Admin sign in</h4>
		{#if loginErr}
			<div class="alert alert-warning mb-8">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-6 w-6 shrink-0 stroke-current"
					fill="none"
					viewBox="0 0 24 24"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
					/></svg
				>
				<span>{loginErr}</span>
			</div>
		{/if}
		<form class="flex w-full flex-col gap-8" use:form>
			<FloatingLabel error={$errors.email} name="email" placeholder="Email" />
			<div class="flex flex-col gap-2">
				<FloatingLabel
					error={$errors.password}
					placeholder="Password"
					name="password"
					type="password"
				/>
				<a class="text-xs text-neutral-400 transition-colors duration-500 hover:text-black" href="/"
					>Forgotten password?</a
				>
			</div>
			<button class="group btn btn-primary btn-block">
				{#if loading}
					<span class="loading loading-spinner" />
				{/if}
				<span class="capitalize">Login</span>
				<ArrowRightIcon class="transition group-hover:translate-x-2" />
			</button>
		</form>
	</div>
</div>
