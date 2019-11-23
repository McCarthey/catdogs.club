import React from 'react'
import { connect } from 'dva'
import { ModelState } from '../types/common'
import styles from './layout.scss'
import { Layout, Menu, Avatar, Dropdown } from 'antd'
import Link from 'umi/link'
import { TransitionGroup, CSSTransition } from 'react-transition-group'

const { Header, Footer, Content } = Layout
const routes = require('../router')

class Index extends React.Component<any, any> {
    constructor(props: any) {
        super(props)
        this.state = {
            username: 'Admin',
            routes: routes.slice(-2, -1)[0].routes,
        }
    }

    // TODO: 将登录状态存储到ls中，发送请求时如遇到cookie-session失效的code，则清除ls中的登录状态，并跳转到登录页

    render() {
        console.log(this.props)
        const menu = (
            <Menu>
                <Menu.Item>
                    <a target="_blank" rel="noopener noreferrer" href="#">
                        个人中心
                    </a>
                </Menu.Item>
                <Menu.Divider />
                <Menu.Item className={styles.logout}>
                    <a target="_blank" rel="noopener noreferrer" href="#">
                        退出登录
                    </a>
                </Menu.Item>
            </Menu>
        )

        // 动态修改 menu 的 defaultSelectedKeys
        const { routes } = this.state
        let defaultSelectedKeys
        try {
            defaultSelectedKeys =
                routes[
                    routes.findIndex((r: { [key: string]: string }) =>
                        r.path.includes(window.location.pathname),
                    )
                ].path
        } catch (e) {
            defaultSelectedKeys = ''
        }

        return (
            <Layout className={styles.normal}>
                <Header className={styles.header}>
                    <div className={styles['menu-container']}>
                        <div className={styles.logo}>Logo</div>
                        <Menu
                            mode="horizontal"
                            defaultSelectedKeys={[defaultSelectedKeys]}
                            style={{ lineHeight: '64px', color: '#fff' }}
                            className={'menu-wrap ' + styles['menu-wrap']}
                        >
                            <Menu.Item key="/home">
                                <Link to="/home">首页</Link>
                            </Menu.Item>
                            <Menu.Item key="/tutorial">
                                <Link to="/tutorial">教程</Link>
                            </Menu.Item>
                            <Menu.Item key="/gif">
                                <Link to="/gif">趣图</Link>
                            </Menu.Item>
                            {/* <Menu.Item key="/forum">
                                <Link to="/forum">论坛</Link>
                            </Menu.Item> */}
                        </Menu>
                    </div>
                    <div className={styles['avatar-wrap']}>
                        {this.props.isLoggedIn ? (
                            <Dropdown overlay={menu} trigger={['hover']}>
                                <div>
                                    <Avatar size="large" icon="user" className={styles.avatar} />
                                    <span className={styles.username}>{this.state.username}</span>
                                </div>
                            </Dropdown>
                        ) : (
                            <div>
                                {/* <Link to="/sign/signin">登录 </Link> /{' '}
                                <Link to="/sign/signup">注册</Link> */}
                            </div>
                        )}
                    </div>
                </Header>
                <TransitionGroup exit={false}>
                    <CSSTransition key={location.pathname} classNames="fade" timeout={300}>
                        <Content className={styles['main-content']}>
                            {this.props.children}
                        </Content>
                    </CSSTransition>
                </TransitionGroup>
                <Footer>
                    Catdogs.club ©{new Date().getFullYear()} Created by McCarthey, Yoko
                </Footer>
            </Layout>
        )
    }
}

function mapStateToProps(state: ModelState) {
    const { isLoggedIn } = state.sign
    return { isLoggedIn }
}

export default connect(mapStateToProps)(Index)