<script lang="ts">
	import { invalidate } from '$app/navigation'
	import { onMount } from 'svelte'
	import '../app.css'
	import AlertList from './AlertList.svelte'

	export let data

	let { supabase, session } = data
	$: ({ supabase, session } = data)

	onMount(() => {
		const {
			data: { subscription }
		} = supabase.auth.onAuthStateChange((event, currentSession) => {
			if (currentSession?.expires_at !== session?.expires_at) {
				invalidate('supabase:auth')
			}
		})

		return () => subscription.unsubscribe()
	})
</script>

<AlertList />
<slot />
