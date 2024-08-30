<script lang="ts">
	import { goto } from '$app/navigation'
	import DateField from '$lib/components/DateField.svelte'
	import SelectField from '$lib/components/SelectField.svelte'
	import TextField from '$lib/components/TextField.svelte'
	import { validator } from '@felte/validator-zod'
	import dayjs from 'dayjs'
	import { createForm } from 'felte'
	import ArrowLeftIcon from '~icons/fa-solid/arrow-left'
	import { CreateUserSchema as schema } from './validate'
	import PasswordField from '$lib/components/PasswordField.svelte'
	import {
		PUBLIC_AUTH0_CALLBACK_URL,
		PUBLIC_AUTH0_CLIENT_ID,
		PUBLIC_API_SERVER_URL
	} from '$env/static/public'
	import type { PageData } from './$types'
	import type { UserProps } from '$lib/common/type'
	import { dialogProps, Notify, openDialog } from '$lib/store'

	let scrollClass = ''
	let loading = false
	export let data: PageData

	function handleContentScroll(panel: HTMLElement) {
		const heightDiff = panel.scrollHeight - panel.offsetHeight
		if (heightDiff > 0) {
			scrollClass = 'scrolled'
		}

		if (panel.scrollTop === 0) {
			scrollClass = 'scroll-reached-top'
		}

		if (panel.scrollTop + panel.offsetHeight === panel.scrollHeight) {
			scrollClass = 'scroll-reached-bottom'
		}
	}

	function goBack() {
		resetDefaultForm()
		goto('/')
	}

	const defaultFormValues = {
		last_name: '',
		first_name: '',
		phone_number: '',
		email: '',
		createdAt: dayjs().format('dd/MM/YYYY'),
		updatedAt: dayjs().format('dd/MM/YYYY'),
		dob: dayjs().format('dd/MM/YYYY'),
		password: '',
		confirmPassword: '',
		connection: 'Username-Password-Authentication',
		agencyId: 1,
		is_active: true,
		verify_email: true,
		auth0_user_id: ''
	}

	function resetDefaultForm() {
		setInitialValues(defaultFormValues)
		reset()
	}

	const { form, errors, setInitialValues, reset } = createForm({
		debounced: { timeout: 500 },
		extend: validator({ schema }),
		onSubmit: save
	})

	const userSchema: {
		name: string
		type: 'text' | 'date' | 'select' | 'password'
		required: boolean
		options?: { label: string; value: string }[]
	}[] = [
		{
			name: 'first_name',
			type: 'text',
			required: true
		},
		{
			name: 'last_name',
			type: 'text',
			required: true
		},
		{
			name: 'dob',
			type: 'date',
			required: true
		},
		{
			name: 'phone_number',
			type: 'text',
			required: true
		},
		{
			name: 'email',
			type: 'text',
			required: true
		},
		{
			name: 'password',
			type: 'text' || 'password',
			required: true
		},
		{
			name: 'confirmPassword',
			type: 'text' || 'password',
			required: true
		}
	]

	let formData: Record<string, string> = {}
	$: userSchema.forEach((field) => {
		if (!(field.name in formData)) {
			formData[field.name] = ''
		}
	})

	function handleLogin() {
		const options = {
			clientID: PUBLIC_AUTH0_CLIENT_ID,
			redirectUri: PUBLIC_AUTH0_CALLBACK_URL,
			responseType: 'code'
		}
		data.webAuthClient.authorize(options)
	}

	async function save(req: UserProps) {
		loading = true
		try {
			const body = {
				email: req.email,
				first_name: req.first_name,
				last_name: req.last_name,
				is_active: true,
				verify_email: false,
				dob: req.dob,
				password: req.password,
				phone_number: req.phone_number,
				connection: 'Username-Password-Authentication',
				agencyId: 1,
				auth0_user_id: ''
			}

			const bodyReq = JSON.stringify(body)

			const request = fetch(`${PUBLIC_API_SERVER_URL}/sign-up`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					accept: 'application/json'
				},
				body: bodyReq
			})

			const response = await request

			if (response.status === 422) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'phía server đã tồn tại dữ liệu này!'
				})
				loading = false
				return
			}

			if (response.status === 500) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'phía server đã tồn tại dữ liệu này!'
				})
				loading = false
				return
			}

			if (response.status === 409) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'Máy chủ đã tồn tại tài khoản email này!'
				})
				loading = false
				return
			}

			resetDefaultForm()
			Notify({
				type: 'success',
				id: crypto.randomUUID(),
				description: 'Đã tạo tài khoản thành công vui lòng liên hệ Quản lý để phê duyệt tài khoản!'
			})
			loading = false
			goto('/')
		} catch (e) {
			console.error('Save error: ', e)
			loading = false
			Notify({ type: 'error', id: crypto.randomUUID(), description: 'Lỗi từ phía server' })
		}
	}
</script>

<div class="relative h-full w-full">
	<section
		class="absolute bottom-0 left-0 right-0 top-1/2 mx-auto flex h-fit max-h-full w-full max-w-md -translate-y-1/2 flex-col items-center justify-center gap-4 rounded-lg border border-solid border-gray-400 p-4"
	>
		<header class="flex w-full flex-row items-center justify-start gap-x-16 md:max-w-2xl">
			<button
				class="btn btn-sm btn-circle btn-outline btn-primary btn-square group"
				on:click={goBack}
			>
				<ArrowLeftIcon class="transition group-hover:-translate-x-1" /></button
			>
			<h2 class="text-center font-sans text-lg font-bold">Mẫu Đăng Ký Tài Khoản</h2>
		</header>
		<section
			class="h-full max-h-80 w-full flex-1 overflow-y-scroll px-2 py-3 {scrollClass}"
			on:scroll={(e) => handleContentScroll(e.currentTarget)}
		>
			<form class="grid grid-cols-1 gap-4 text-sm" id="signupForm" use:form>
				{#each userSchema as { type, name, required, options } (name)}
					{#if type === 'text' || type === 'password'}
						{#if name == 'password' || name == 'confirmPassword'}
							<PasswordField error={$errors[name]} {name} {required} bind:value={formData[name]} />
						{:else}
							<TextField error={$errors[name]} {name} {required} />
						{/if}
					{:else if type === 'select'}
						<SelectField error={$errors[name]} {name} {options} {required} />
					{:else if type === 'date'}
						<DateField error={$errors[name]} {name} {required} />
					{/if}
				{/each}
			</form>
		</section>
		<footer class="flex w-full flex-shrink-0 items-center justify-center px-4 py-3">
			<div class="flex w-full flex-col items-center justify-around gap-5 sm:flex-row">
				<button
					class="btn btn-ghost rounded normal-case active:!translate-y-1"
					on:click={handleLogin}>Đã có tài khoản</button
				>
				<button
					class="btn rounded px-10 normal-case active:!translate-y-1"
					type="submit"
					form="signupForm"
				>
					{#if loading}
						<span class="loading loading-spinner" />
					{:else}
						Đăng Ký
					{/if}
				</button>
			</div>
		</footer>
	</section>
</div>
