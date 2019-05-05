import httpRequester from '@/utils/request'
import { SignUpReq } from '@/types/sign'

function signUpByEmail({ email, password }: SignUpReq) {
    console.log(httpRequester.postJSON)
    return httpRequester.postJSON('/api/register', { email, password })
}

function signInByEmail({ email, password }: SignUpReq) {
    return httpRequester.postJSON('/api/login', { email, password })
}

export default {
    signUpByEmail,
    signInByEmail,
}
