import React from 'react'
import { connect } from 'dva'
import { Form, Icon, Input, Button, Checkbox } from 'antd'
import Link from 'umi/link'
import styles from './style.scss'

import { SignUpReq } from '@/types/sign'

class SignUp extends React.Component<any, any> {
    constructor(props: any) {
        super(props)
        this.state = {}
    }

    handleSubmit = (e: any) => {
        e.preventDefault()
        this.props.form.validateFields((err: any, values: SignUpReq) => {
            this.props.dispatch({
                type: 'sign/login',
                payload: values,
            })
        })
    }

    render() {
        const { getFieldDecorator } = this.props.form
        return (
            <Form onSubmit={this.handleSubmit} className={styles['login-form']}>
                <Form.Item>
                    {getFieldDecorator('email', {
                        rules: [{ required: true, message: '请输入邮箱' }],
                    })(
                        <Input
                            prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            placeholder="用户名"
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
                    {getFieldDecorator('remember', {
                        valuePropName: 'checked',
                        initialValue: true,
                    })(<Checkbox>记住密码</Checkbox>)}
                    <Link className={styles['login-form-forgot']} to="">
                        忘记密码
                    </Link>
                    <Button
                        type="primary"
                        htmlType="submit"
                        className={styles['login-form-button']}
                    >
                        登录
                    </Button>
                    <Link to="/sign/signup">立即注册</Link>
                </Form.Item>
            </Form>
        )
    }
}

const WrappedSignUpForm = Form.create({ name: 'normal_login' })(SignUp)

function mapStateToProps(state: any) {
    const { isLoggedIn } = state
    return { isLoggedIn }
}

export default connect(mapStateToProps)(WrappedSignUpForm)
