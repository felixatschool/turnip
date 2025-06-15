<script>
	import { onMount } from 'svelte';
	import { config } from '$lib/config';

	let message = '...';

	onMount(async () => {
		try {
			const res = await fetch(`${config.apiBaseUrl}/health`);
			if (!res.ok) throw new Error('Network error');
			const data = await res.json();
			message = data.status || 'No message received';
		} catch (e) {
			message = 'Failed to load message';
		}
	});
</script>

<main>
	<p>{message}</p>
</main>
