import { IConfig } from 'umi-types';

// ref: https://umijs.org/config/
const config: IConfig = {
  treeShaking: true,
  urlLoaderExcludes: [/.scss$/],
  chainWebpack(config) {
	// 使用scss-resources-loader为项目注入全局变量
	config.module
		.rule('sass-resources')
		.test(/.scss$/)
		.use(['style-loader', 'css-loader', 'postcss-loader', 'sass-loader'])
		.loader('sass-resources-loader')
		.options({ resources: './src/assets/resources.scss' });
  },
  proxy: {
	'/api': {
		target: 'http://129.204.46.253:8888/api',
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
		dva: true,
		dynamicImport: { webpackChunkName: true },
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
		},
	],
  ],
};

export default config;
