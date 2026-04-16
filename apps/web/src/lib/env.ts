// Typed access to Vite-exposed environment variables.
//
// Vite only exposes vars prefixed with VITE_ to the client bundle, and they
// are inlined at build time. In production we want to fail loud if a required
// var is missing rather than send `undefined` to the API.

const isProd = import.meta.env.PROD

function required(name: string, value: string | undefined): string {
  if (value && value.length > 0) return value
  if (isProd) {
    throw new Error(
      `Required env var ${name} is missing. Set it in the Vercel project ` +
        `dashboard (or your local apps/web/.env file).`,
    )
  }
  return ''
}

export const env = {
  apiBaseUrl: required('VITE_API_BASE_URL', import.meta.env.VITE_API_BASE_URL),
} as const
