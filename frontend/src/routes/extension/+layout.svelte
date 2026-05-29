<script lang="ts">
	import { Header } from '$widgets';
	import { Controls } from '$widgets/extension';
	import { Events } from '@wailsio/runtime';
	import { onMount, setContext } from 'svelte';

	let { children } = $props();
	let tabs = [
		{ title: 'Metrics', path: '/extension' },
		{ title: 'Accounts', path: '/extension/accounts' },
		{ title: 'Links', path: '/extension/links' }
	];

	onMount(() => {
		Events.On('app:update', (event) => {
			console.log('app:update', event.data);
		});
	});
	setContext('ext', { data: { test: 'test' } });
</script>

<Header {tabs} />
<div class="relative flex flex-col grow ">
	{@render children?.()}
	<Controls />
</div>
