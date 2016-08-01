import webpack from "webpack";
import common from "./webpack.common";

module.exports = {
  ...common,
  entry: {
    app: [
      `webpack-dev-server/client?http://localhost:3000`,
      'webpack/hot/only-dev-server',
      './src/index.js'
    ]
  },
  devtool: 'eval',
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    ...common.plugins,
  ]
};
