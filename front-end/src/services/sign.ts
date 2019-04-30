import httpRequester from '@/utils/request'
import { SignUpReq } from '@/types/sign'

function signUpByEmail({ email, password }: SignUpReq) {
    return httpRequester('/api/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            email,
            password,
        }),
    })
}

export default {
    signUpByEmail,
}
