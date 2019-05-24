import fetch from 'dva/fetch'
import { message } from 'antd'

interface HttpError {
    response?: any
}

function checkStatus(response: { status: number; statusText: string | undefined }) {
    if (response.status >= 200 && response.status < 300) {
        return response
    }

    const error = <HttpError>(new Error(response.statusText))
    error.response = response
    throw error
}

function handleResponseData(data: any, toast: boolean) {
    const ret = {
        data,
        headers: {},
    }

    // 如果错误码 !== 0 默认弹出 toast
    if (toast && ret.data && ret.data.code !== 0) {
        message.error(ret.data.msg)
        throw ret.data
    }
    return ret.data
}

/**
 * Requests a URL, returning a promise.
 *
 * @param  {string} url       The URL we want to request
 * @param  {object} [options] The options we want to pass to "fetch"
 * @return {object}           An object containing either "data" or "err"
 */
const request = {
    postJSON,
}

async function postJSON (
    url: string,
    body: { [key: string]: any },
    options: any = {},
    toast: boolean = true,
) {
    const response = await fetch(url, {
        ...options,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
    })

    checkStatus(response)

    const data = await response.json()

    handleResponseData(data, toast)
}

export default request
