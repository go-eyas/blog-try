import webpack from 'webpack';
import path from 'path';

const commonConfig = {
  resolve: {
    root: path.join(__dirname, 'src/vendor'),
    extensions: ['', '.js', '.jsx', '.css', '.scss'],
    alias: {
      root: path.join(__dirname, '../'),
      src: path.join(__dirname, '../src'),
      vendor: path.join(__dirname, '../src/vendor'),
      common: path.join(__dirname, '../src/common'),
      cc: path.join(__dirname, '../src/common/components'),
      modules: path.join(__dirname, '../src/modules'),
      utils: path.join(__dirname, '../src/utils'),


      framework: path.join(__dirname, '../src/modules/framework'),
      blog: path.join(__dirname, '../src/modules/blog'),
      user: path.join(__dirname, '../src/modules/user'),
      admin: path.join(__dirname, '../src/modules/admin')
    }
  },
  output: {
    path: path.join(__dirname, '../../prod'),
    filename: '[name].js',
    chunkFilename: '[name].[chunkhash:5].chunk.js',
    publicPath: '/static/'
  },
  module: {
    noParse:[
      path.join(__dirname, '../node_modules/highlight.js/lib/languages/autoit.js')
    ],
    loaders: [
      {
        test: /\.(js|jsx)$/,
        loader: 'babel',
        include: [path.join(__dirname, '../src'), path.join(__dirname, './config')],
        exclude: [path.join(__dirname, '../node_modules'), path.join(__dirname, '../src/vendor')]
      }, {
        test: /\.css$/,
        loader: 'style-loader!css-loader?modules&importLoaders=2&localIdentName=[local]'
      }, {
        test: /\.(sass|scss)$/,
        loader: 'style-loader!css-loader?modules&importLoaders=2&localIdentName=[local]___[hash:base64:5]!postcss-loader',
        exclude: [path.join(__dirname, '../src/common/style')]
      }, {
        test: /\.(sass|scss)$/,
        loader: 'style-loader!css-loader!postcss-loader',
        include: [path.join(__dirname, '../src/common/style')]
      }, {
        test: /\.(png|jpg|jpeg|gif)$/,
        loader: 'url-loader?limit=8192&name=resource/img/[hash].[ext]'
      }, {
        test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
        loader: 'url-loader?limit=10000&name=resource/font/[hash].[ext]&minetype=application/font-woff'
      }, {
        test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
        loader: 'file-loader?name=resource/font/[hash].[ext]'
      }
    ]
  },
  postcss: () => [
    // require('postcss-inline-comment'),
    require('precss')
  ],
  plugins: [
    new webpack.NoErrorsPlugin(),

    new webpack.DefinePlugin({
        'process.env.NODE_ENV': '"'+process.env.NODE_ENV+'"'
    }),
    // 全局变量
    // new webpack.ProvidePlugin({
    //   cx: 'classname',
    //   config: path.join(__dirname, './config.client')
    // })
  // css 文件单独打包
  // new ExtractTextPlugin('style.css', {
  //     allChunks: true
  // }),
  ]
};
export default commonConfig;
