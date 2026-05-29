<script lang="ts">
	import * as ParseCSV from '$go/services/parse-csv/csv';
	import { Clear, Copy, Input } from '$shared/ui';
	import { Table } from '$widgets/parseCsv';
	import { sidebarState } from '$widgets/sidebar';
	import { copyFormatCSV, filterCSV } from './services';
	import { FilePlus, Menu } from '@lucide/svelte';
	import { Events } from '@wailsio/runtime';

	let value: string[][] = $state([]);
	// let filtered = $derived(filterCSV(value, value[0].length ? value : undefined));

	Events.On('parseCsv', () => {
		alert('ParseEvent');
	});
</script>

<div class="flex flex-col h-screen overflow-hidden">
	<div class="flex justify-between bg-black text-white">
		<Menu
			class="h-full"
			onclick={() => {
				sidebarState.set(true);
			}}
		/>
		<Clear onclick={ParseCSV.Clear} />
		<Input
			type="select"
			bind:value
			options={value[0]?.map((el) => ({ title: el, value: el }))}
			multiple
			class="p-1 bg-amber-700 w-full"
		/>
		<Copy value={copyFormatCSV(value)} />
	</div>

	<div class="relative flex-1 min-h-0 p-4 flex flex-col">
		<div class="h-full p-2 absolute inset-0" onclick={ParseCSV.ParseFiles} role="none">
			<div
				class="h-full rounded border bg-[rgba(60,60,60,1)]"
				id={'parse-csv'}
				data-file-drop-target
			>
				<div
					class="flex flex-col gap-2 border rounded h-full w-full items-center justify-center cursor-pointer text-white"
				>
					<FilePlus class="w-10 h-10 " />
					<span>Choose a file or drag it over here</span>
				</div>
			</div>
		</div>
		<div class="flex-1 overflow-auto">
			<Table data={value} />
		</div>
	</div>
</div>
