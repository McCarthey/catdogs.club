import React from 'react'
import { Form, Icon, Input, Button, Checkbox } from 'antd'
import styles from './style.scss'

class SignUp extends React.Component<any, any> {
    constructor(props: any) {
        super(props)
        this.state = {}
    }

    handleSubmit = (e: any) => {
        e.preventDefault()
        this.props.form.validateFields((err: any, values: any) => {
            if (!err) {
                console.log('Received values of form: ', values)
            }
        })
    }

    render() {
        const { getFieldDecorator } = this.props.form
        return (
            <Form onSubmit={this.handleSubmit} className={styles['login-form']}>
                <Form.Item>
                    {getFieldDecorator('userName', {
                        rules: [{ required: true, message: 'Please input your username!' }],
                    })(
                        <Input
                            prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                            placeholder="用户名"
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
                    <Button type="primary" htmlType="submit" className={styles["login-form-button"]}>
						注册
					</Button>
                    已有帐号？ <a href="/signin">即刻登录</a>
                </Form.Item>
            </Form>
        )
    }
}

const WrappedSignUpForm = Form.create({ name: 'normal_login' })(SignUp);

export default WrappedSignUpForm