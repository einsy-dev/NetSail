export async function fileToByte(f: File): Promise<string> {
	const arrayBuffer = await f.arrayBuffer();
	const blob = new Blob([arrayBuffer]);
	const reader = new FileReader();

	return new Promise((res, rej) => {
		reader.onload = () => {
			const result = reader.result as string;
			res(result.substring(result.indexOf(',') + 1));
		};
		reader.onerror = () => rej(reader.error);
		reader.readAsDataURL(blob);
	});
}
