<script lang="ts">
	import type { InputI } from '../input';
	import { FilePlus } from '@lucide/svelte';

	let { id, children, class: className }: InputI['file'] = $props();

	const style = 'cursor-pointer overflow-hidden text-white';
	const styleActive = 'bg-red-500';
</script>

<div
	class="drop-zone {style} data-[file-drop-target-active]:{styleActive}"
	{id}
	data-file-drop-target
>
	<div class="relative flex grow">
		{@render children?.()}
		<div class="absolute inset-0 z-50 flex items-center justify-center bg-black/50">
			<FilePlus class="h-30 w-30" />
		</div>
	</div>
</div>
