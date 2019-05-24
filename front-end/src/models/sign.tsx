import React from 'react'
import api_sign from '../services/sign'
import { SignUpReq } from '@/types/sign'
import { DvaEffect } from '@/types/common'
import { message, Modal } from 'antd'
import router from 'umi/router'
import { SignModelState } from '../types/common'

interface LoginAction {
    payload: {
        isLoggedIn: boolean
    }
}

export default {
    state: {
        isLoggedIn: false,
    } as SignModelState,
    reducers: {
        changeLogState(state: SignModelState, { payload: { isLoggedIn } }: LoginAction) {
            return { ...state, isLoggedIn }
        },
    },
    effects: {
        *signUp({ payload: { email, password }}: { payload: SignUpReq }, { call, put }: DvaEffect) {
            const res = yield call(api_sign.signUpByEmail, { email, password })
            console.log(res)
            Modal.success({
                title: '注册成功',
                content: (
                    <div>
                        请前往
                        <span style={{ fontWeight: 'bold' }}>{email}</span>
                        查看收件箱，激活账号
                    </div>
                ),
                onOk() {
                    router.push('/sign/signin')
                },
            })
        },
        *login({ payload: { email, password }}: { payload: SignUpReq }, { call, put }: DvaEffect) {
            const res = yield call(api_sign.signInByEmail, { email, password })
            yield put({ type: 'changeLogState', payload: { isLoggedIn: true } })
            message.success(res.msg)
            router.push('/home')
        },
        *logout() {
            yield // TODO:暂无推出登录接口
        }
    },
}
