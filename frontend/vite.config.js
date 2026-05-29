import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import wails from '@wailsio/runtime/plugins/vite';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
	server: {
		host: '127.0.0.1',
		port: 9245,
		strictPort: true,
	},
	plugins: [wails('./src/lib'), tailwindcss(), sveltekit()]
});
