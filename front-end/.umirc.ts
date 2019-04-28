import { IConfig } from 'umi-types';

// ref: https://umijs.org/config/
const config: IConfig = {
  treeShaking: true,
  urlLoaderExcludes: [/.scss$/],
  chainWebpack(config) {
    config.module
      .rule('sass-resources')
      .test(/.scss$/)
      .use(['style-loader', 'css-loader', 'postcss-loader', 'sass-loader'])
      .loader('sass-resources-loader')
      .options({resources: './src/assets/resources.scss'})
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
