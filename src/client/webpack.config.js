const path = require("path");
const hwp = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: "./src/app/index.tsx",
  devtool: "inline-source-map",
  module: {
    rules: [
      {
        test: /\.(tsx|ts)?$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.s[ac]ss$/i,
        use: [
          // Creates `style` nodes from JS strings
          "style-loader",
          // Translates CSS into CommonJS
          "css-loader",
          // Compiles Sass to CSS
          "sass-loader",
        ],
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js"],
  },
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "src", "build"),
  },
  plugins: [
    new hwp({
      template: path.resolve(__dirname, "src", "app", "static", "index.html"),
    }),
  ],
};
