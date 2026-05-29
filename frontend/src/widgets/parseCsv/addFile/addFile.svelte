<script lang="ts">
	import { ParseCsv } from '$go/services/server/api';
	import { FilePlus } from '@lucide/svelte';

	let { files = $bindable([]) }: { files?: File[]; 'bind:files'?: File[] } = $props();

	const style = 'h-full rounded border bg-[rgba(60,60,60,1)]';
</script>

<div class="h-full p-2 absolute inset-0" onclick={ParseCsv} role="none">
	<div class={style} id={'parse-csv'} data-file-drop-target>
		<div
			class="flex flex-col gap-2 border rounded h-full w-full items-center justify-center cursor-pointer text-white"
		>
			<FilePlus class="w-10 h-10 " />
			<span>Choose a file or drag it over here</span>
		</div>
	</div>
</div>
