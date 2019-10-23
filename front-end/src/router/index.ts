const routes = [
    {
        path: '/sign',
        component: './sign/_layout',
        routes: [
            { path: '/sign/signup', component: './sign/signUp' }, // 不能放在路由表末尾
            { path: '/sign/signin', component: './sign/signIn' }, // 同上
        ],
    },
    {
        path: '/',
        component: './_layout',
        routes: [
            {
                path: '/gif',
                component: './gif',
            },
            {
                path: '/home',
                component: './home',
            },
            // {
            //     path: '/forum',
            //     component: './forum',
            // },
            {
                path: '/tutorial',
                component: './tutorial',
            },
            {
                path: '/',
                redirect: '/home',
            },
            {
                path: '*',
                component: './page404',
            },
        ],
    },
    {
        path: '*',
        component: './page404',
    },
]

module.exports = routes
