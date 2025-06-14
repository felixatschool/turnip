<script>
import { onMount } from 'svelte';

let message = '...';

onMount(async () => {
    try {
    const res = await fetch('https://staging-turnip-backend.fly.dev/api/health');
    if (!res.ok) throw new Error('Network error');
    const data = await res.json();
    message = data.status || 'No message received';
    } catch (e) {
    message = 'Failed to load message';
    }
    });
</script>

<style>
main {
height: 100vh;
display: flex;
         justify-content: center;
         align-items: center;
background: linear-gradient(to bottom, #6a0dad, #ffffff);
            font-family: system-ui, sans-serif;
color: #4b0082;
       font-size: 2rem;
       text-align: center;
padding: 1rem;
}
</style>

<main>
<p>{message}</p>
</main>

