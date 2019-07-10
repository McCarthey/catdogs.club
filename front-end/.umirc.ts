import { IConfig } from 'umi-types'
const routes = require('./src/router')

// ref: https://umijs.org/config/
const config: IConfig = {
    outputPath: './catdogs',
    treeShaking: true,
    routes,
    urlLoaderExcludes: [/.scss$/],
    theme: {
        'primary-color': '#e91e63',
    },
    chainWebpack(config) {
        // 使用scss-resources-loader为项目注入全局变量
        config.module
            .rule('sass-resources')
            .test(/.scss$/)
            .use(['style-loader', 'css-loader', 'postcss-loader', 'sass-loader'])
            .loader('sass-resources-loader')
            .options({ resources: './src/assets/resources.scss' })
    },
    proxy: {
        '/api': {
            target: 'http://118.24.146.34:8888/api',
            changeOrigin: true,
            pathRewrite: { '^/api': '/' },
        },
    },
    plugins: [
        // ref: https://umijs.org/plugin/umi-plugin-react.html
        [
            'umi-plugin-react',
            {
                antd: true,
                dva: {
                    immer: true,
                },
                title: 'front-end',
                dll: false,

                routes: {
                    exclude: [
                        /models\//,
                        /services\//,
                        /model\.(t|j)sx?$/,
                        /service\.(t|j)sx?$/,
                        /components\//,
                    ],
                },

                dynamicImport: {
                    loadingComponent: './components/PageLoading/index',
                    webpackChunkName: true,
                    level: 3,
                },
            },
        ],
    ],
}

export default config
