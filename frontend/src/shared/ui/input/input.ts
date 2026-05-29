import type { Snippet } from 'svelte';
import type { HTMLInputAttributes } from 'svelte/elements';

export type InputIDef = Pick<HTMLInputAttributes, 'type' | 'value' | 'placeholder' | 'class'> & {
	type?: HTMLInputAttributes['type'];
	label?: string;
};

export interface InputI {
	email: InputIDef & { type?: 'email' };
	password: InputIDef & { type?: 'password' };
	username: InputIDef & { type?: 'username' };
	file: InputIDef & { type?: 'file'; label?: string; children?: Snippet; id?: string };
	select: InputIDef & { type?: 'select'; options: Option[]; multiple?: boolean };
	html: InputIDef & { type?: 'html' };
	textarea: InputIDef & { type?: 'textarea' };
}

interface Option {
	title: string;
	value?: string | number;
	callback?: (value: any) => void;
}
