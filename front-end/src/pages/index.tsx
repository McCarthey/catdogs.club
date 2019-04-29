import React from 'react';
import styles from './index.scss';
import { Layout, Menu, Avatar, Dropdown } from 'antd';
const { Header, Footer, Content } = Layout;

export default class Index extends React.Component<any, any> {
  constructor(props: any) {
	super(props);
	this.state = {
		username: 'Admin',
	};
  }

  render() {
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
	);

	return (
		<Layout className={styles.normal}>
		<Header className={styles.header}>
			<div>
			<div className="logo" />
			<Menu
				mode="horizontal"
				defaultSelectedKeys={['1']}
				style={{ lineHeight: '64px', color: '#fff' }}
				className={'menu-wrap ' + styles['menu-wrap']}
			>
				<Menu.Item key="1">nav 1</Menu.Item>
				<Menu.Item key="2">nav 2</Menu.Item>
				<Menu.Item key="3">nav 3</Menu.Item>
			</Menu>
			</div>
			<div className={styles['avatar-wrap']}>
			<Dropdown overlay={menu} trigger={['hover']}>
				<div>
				<Avatar size="large" icon="user" className={styles.avatar} />
				<span className={styles.username}>{this.state.username}</span>
				</div>
			</Dropdown>
			</div>
		</Header>
		<Content>
			<div style={{ background: '#fff', padding: 24, minHeight: 280 }}>Content</div>
		</Content>
		<Footer>Catdogs.club ©{new Date().getFullYear()} Created by McCarthey, Yoko</Footer>
		</Layout>
	);
  }
}
