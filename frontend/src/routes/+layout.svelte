<script lang="ts">
	import Sidebar from '$widgets/sidebar/index.svelte';
	import './layout.css';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={'#'} /></svelte:head>
<div class="relative h-screen flex flex-col bg-[rgba(43,43,43,1)]">
	<Sidebar />
	<div class="relative flex flex-col grow overflow-y-scroll">
		{@render children?.()}
	</div>
</div>
