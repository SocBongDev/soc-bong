<script lang="ts">
	import { CloseNotification, notifications } from '$lib/store'
	import Alert from './Alert.svelte'

	const autoClose = (id: string) => () => setTimeout(handleClose(id), 3000)

	function handleClose(id: string) {
		return () => CloseNotification(id)
	}
</script>

<div class="absolute bottom-10 right-1/2 z-50 translate-x-1/2">
	<div class="stack">
		{#each $notifications as notification (notification.id)}
			<Alert autoClose={autoClose(notification.id)} on:click={handleClose(notification.id)}>
				{notification.description}
			</Alert>
		{/each}
	</div>
</div>
