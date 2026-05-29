<script lang="ts">
	import { Input } from '$shared/ui';
	import { AccountsList } from '$widgets/extension';
	import { AddFile } from '$widgets/parseCsv';

	let files: File[] = $state([]);
</script>

<Input type="file" class="flex grow" bind:value={files} multiple>
	<div class="flex grow flex-col text-white px-2 py-1 gap-1" role="none">
		<AccountsList
			accounts={[
				{
					domain: 'google.com',
					username: 'bob',
					password: 'bobik',
					note: 'New accoun to implemnet new forum features'
				},
				{
					domain: 'google.com',
					username: 'bob',
					password: 'bobik',
					note: 'New accoun to implemnet new forum features'
				},
				{
					domain: 'google.com',
					username: 'bob',
					password: 'bobik',
					note: 'New accoun to implemnet new forum features'
				}
			]}
		/>
	</div>
</Input>
