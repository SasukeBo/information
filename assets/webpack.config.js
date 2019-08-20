const path = require('path');
const LiveReloadPlugin = require('webpack-livereload-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const AssetPlugin = require('assets-webpack-plugin');
const VueLoaderPlugin = require('vue-loader/lib/plugin');
const isProd = process.env.NODE_ENV === 'production';
const ROOT_PATH = path.resolve(__dirname, './');


var plugins = function () {
  var plugins =
    [
      new MiniCssExtractPlugin({ filename: isProd ? 'css/[name].[hash].css' : 'css/[name].css' }),
      new VueLoaderPlugin(),
      new AssetPlugin({ filename: '../static/manifest.json' }),
    ]

  if (!isProd) {
    plugins.push(new LiveReloadPlugin({ delay: 200 }))
  }

  return plugins
}()

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
    'app': './js/app.js',
    'babel-polyfill': 'babel-polyfill'
  },
  devtool: 'source-map',
  output: {
    publicPath: '/',
    filename: isProd ? 'js/[name].[hash].js' : 'js/[name].js',
    // filename: 'js/[name].[hash].js',
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
  optimization: {
    splitChunks: {
      chunks: 'all',
      minSize: 30000,
      maxSize: 0,
      minChunks: 1,
      maxAsyncRequests: 5,
      maxInitialRequests: 3,
      automaticNameDelimiter: '~',
      name: true,
      cacheGroups: {
        element: {
          test: /[\\/]node_modules[\\/]element-ui[\\/]/,
          name: 'element',
          chunks: 'all'
        },
        vue: {
          test: /[\\/]node_modules[\\/]vue.*[\\/]/,
          name: 'vue',
          chunks: 'all'
        },
        default: {
          minChunks: 2,
          priority: -20,
          reuseExistingChunk: true
        }
      }
    }
  },
  plugins,
});
