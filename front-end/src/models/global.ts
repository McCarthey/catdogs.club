import api_sign from '../services/sign'
import { SignUpReq } from '@/types/sign'
import { DvaEffect } from '@/types/common'

interface GlobalModelState {
    isLoggedIn: boolean
}

interface LoginAction {
    payload: {
        isLoggedIn: boolean
    }
}

export default {
    state: {
        isLoggedIn: false,
    } as GlobalModelState,
    reducers: {
        changeLogState(state: GlobalModelState, { payload: { isLoggedIn } }: LoginAction) {
            return { ...state, isLoggedIn }
        },
    },
    effects: {
        *login({ payload: { email, password } }: { payload: SignUpReq }, { call, put }: DvaEffect) {
            yield call(api_sign.signInByEmail, { email, password })
            yield put({ type: 'changeLogState', payload: { isLoggedIn: true } })
        },
    },
}
