import { API_HOST } from '../config'

export interface AuthResponse {
    token: string
  }
  
export async function signIn(email: string, password: string): Promise<AuthResponse> {
const response = await fetch(`${API_HOST}/api/v1/sign_in`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
})

if (!response.ok) {
    throw new Error('Login failed')
}

return await response.json()
} 