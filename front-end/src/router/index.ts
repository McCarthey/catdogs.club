const routes = [
	{ path: '/', redirect: '/home' },
	{ path: '/signup', component: './Sign/SignUp' }, // 不能放在路由表末尾
	{ path: '/signin', component: './Sign/SignIn' }, // 同上  
	{
		path: '/',
		component: './_layout',
		routes: [{ path: '/home', component: './Home' }],
	},
]

module.exports = routes
