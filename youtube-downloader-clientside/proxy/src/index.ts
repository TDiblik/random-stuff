import jwt from '@tsndr/cloudflare-worker-jwt';

interface Env {
	JWT_SECRET: string;
}

// We want to return this response on every single possible fail to confuse possible abusers :).
function EmptyResponse() {
	return new Response('<h1> Nothing to see here :) </h1>', {
		headers: {
			'Content-Type': 'text/html; charset=utf-8',
		},
	});
}

export default {
	async fetch(request, env): Promise<Response> {
		if (request.method === 'OPTIONS') {
			return new Response(null, {
				headers: {
					'Access-Control-Allow-Origin': '*',
					'Access-Control-Allow-Methods': 'GET, OPTIONS',
					'Access-Control-Max-Age': '86400',
					'Access-Control-Allow-Headers': '*',
				},
			});
		}

		const req_url = new URL(request.url);
		const target_format_index = parseInt(req_url.searchParams.get('hidden_parameter_for_the_format_index_to_fuck_with_bots') ?? '{}');
		let supplied_body;
		try {
			supplied_body = (await request.json()) as any;
		} catch {
			return EmptyResponse();
		}

		console.log(supplied_body.hidden_parameter_for_the_token_to_fuck_with_bots);
		if (
			supplied_body == null ||
			supplied_body.hidden_parameter_for_the_token_to_fuck_with_bots == null ||
			target_format_index == null ||
			Number.isNaN(target_format_index)
		) {
			return EmptyResponse();
		}

		const supplied_token = supplied_body.hidden_parameter_for_the_token_to_fuck_with_bots;
		try {
			if (!(await jwt.verify(supplied_token, env.JWT_SECRET))) {
				return EmptyResponse();
			}
		} catch {
			return EmptyResponse();
		}

		const { payload } = jwt.decode(supplied_token);
		if (payload == null) {
			return EmptyResponse();
		}

		const video_metadata = (payload as any).video_metadata;
		if (video_metadata == null || video_metadata.formats == null || video_metadata.formats.length == 0) {
			return EmptyResponse();
		}

		const desired_format = video_metadata.formats[target_format_index];
		if (desired_format == null || desired_format.url == null) {
			return EmptyResponse();
		}

		let response = await fetch(desired_format.url, { ...request, redirect: 'follow' });
		response = new Response(response.body, response);
		response.headers.set('Access-Control-Allow-Origin', '*');
		response.headers.set('Access-Control-Allow-Methods', 'GET, OPTIONS');
		response.headers.set('Access-Control-Max-Age', '86400');
		response.headers.set('Access-Control-Allow-Headers', '*');
		return response;
	},
} satisfies ExportedHandler<Env>;
