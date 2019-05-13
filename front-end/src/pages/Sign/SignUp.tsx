import React from 'react'
import { Form, Icon, Input, Button, Modal } from 'antd'
import Link from 'umi/link'
import router from 'umi/router'
import styles from './style.scss'

import api_sign from '@/services/sign'

import { SignUpReq } from '@/types/sign'

interface signUpState {
    email: string
    registerDone: boolean
}

class SignUp extends React.Component<any, signUpState> {
    constructor(props: any) {
        super(props)
        this.state = {
            email: '',
            registerDone: false,
        }
    }

    handleSignUp = (e: any) => {
        e.preventDefault()
        this.props.form.validateFields(async (err: any, values: SignUpReq) => {
            if (!err) {
                try {
					const res = await api_sign.signUpByEmail(values)
                    this.setState(
                        {
                            email: this.props.form.getFieldValue('email'),
                            registerDone: true,
                        },
                        () => {
                            Modal.success({
                                title: '注册成功',
                                content: (
                                    <div>
                                        请前往
                                        <span style={{ fontWeight: 'bold' }}>
                                            {this.state.email}
                                        </span>
                                        查看收件箱，激活账号
                                    </div>
								),
								onOk() { router.push('/sign/signin')}
                            })
                        },
                    )
                } catch (e) {
                    console.log('error code:', e.code)
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
                        rules: [{ required: true, message: '请输入邮箱' }],
                    })(
                        <Input
                            prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            placeholder="邮箱"
                        />,
                    )}
                </Form.Item>
                <Form.Item>
                    {getFieldDecorator('password', {
                        rules: [{ required: true, message: '请输入密码' }],
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
                    已有帐号？ <Link to="/sign/signin">即刻登录</Link>
                </Form.Item>
            </Form>
        )
    }
}

const WrappedSignUpForm = Form.create({ name: 'normal_login' })(SignUp)

export default WrappedSignUpForm
