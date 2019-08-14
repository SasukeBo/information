const path = require('path');
const LiveReloadPlugin = require('webpack-livereload-plugin');
// const glob = require('glob');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const AssetPlugin = require('assets-webpack-plugin');
const VueLoaderPlugin = require('vue-loader/lib/plugin');
const isProd = process.env.NODE_ENV === 'production';
const ROOT_PATH = path.resolve(__dirname, './');

module.exports = (env, options) => ({
  optimization: {
    minimizer: [
      new UglifyJsPlugin({ cache: true, parallel: true, sourceMap: true }),
      new OptimizeCSSAssetsPlugin({})
    ]
  },
  resolve: {
    extensions: ['.js', '.vue'],
    alias: {
      '~': ROOT_PATH,
      js: path.resolve(__dirname, 'js'),
      css: path.resolve(__dirname, 'css'),
      images: path.resolve(__dirname, 'images')
    }
  },
  entry: {
      'app': ['./js/app.js'],
      'babel-polyfill': ['babel-polyfill']
  },
  devtool: 'source-map',
  output: {
    publicPath: '/',
    filename: isProd ? 'js/[name].[contenthash].js' : 'js/[name].js',
    path: path.resolve(__dirname, '../static/')
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: {
          extractCSS: true,
          preserveWhitespace: false,
        }
      },
      {
        test: /\.(ttf|woff|eot|svg|woff2)(\?\S*)?$/,
        use: [
          'url-loader'
        ]
      },
      {
        test: /\.(png|svg|jpg|gif)$/,
        use: [
          {
            loader: 'url-loader',
            options: {
              name: '[hash].[ext]',
              outputPath: 'images/'
            }
          }
        ]
      },
      {
        test: /\.(c|sc|sa)ss$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          'css-loader?sourceMap',
          'sass-loader?sourceMap'
        ]
      }
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({ filename: isProd ? 'css/[name].[contenthash].css' : 'css/[name].css' }),
    new VueLoaderPlugin(),
    new AssetPlugin({ filename: '../static/manifest.json' }),
    new LiveReloadPlugin({delay: 200})
  ]
});
