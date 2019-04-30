import React from 'react'
import { Form, Icon, Input, Button, Checkbox } from 'antd'
import styles from './style.scss'

import api_sign from '@/services/sign'

import { SignUpReq } from '@/types/sign'

class SignUp extends React.Component<any, any> {
    constructor(props: any) {
        super(props)
    }

    handleSignUp = (e: any) => {
        e.preventDefault()
        this.props.form.validateFields(async (err: any, values: SignUpReq) => {
            if (!err) {
                try {
                    const res = await api_sign.signUpByEmail(values)
                    console.log('ressss', res)
                } catch (e) {
                    console.log(e)
                }
            }
        })
    }

    render() {
        const { getFieldDecorator } = this.props.form
        return (
            <Form onSubmit={this.handleSignUp} className={styles['login-form']}>
                <Form.Item>
                    {getFieldDecorator('email', {
                        rules: [{ required: true, message: 'Please input your username!' }],
                    })(
                        <Input
                            prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            placeholder="邮箱"
                        />,
                    )}
                </Form.Item>
                <Form.Item>
                    {getFieldDecorator('password', {
                        rules: [{ required: true, message: 'Please input your Password!' }],
                    })(
                        <Input
                            prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            type="password"
                            placeholder="密码"
                        />,
                    )}
                </Form.Item>
                <Form.Item>
                    <Button
                        type="primary"
                        htmlType="submit"
                        className={styles['login-form-button']}
                    >
                        注册
                    </Button>
                    已有帐号？ <a href="/signin">即刻登录</a>
                </Form.Item>
            </Form>
        )
    }
}

const WrappedSignUpForm = Form.create({ name: 'normal_login' })(SignUp)

export default WrappedSignUpForm
