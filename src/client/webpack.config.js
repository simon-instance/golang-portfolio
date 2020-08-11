const path = require("path");
const hwp = require("html-webpack-plugin");
const terser = require("terser-webpack-plugin");
const extracter = require("mini-css-extract-plugin");
const cssminifier = require("optimize-css-assets-webpack-plugin");

const isDev = process.env.NODE_ENV === "dev";

module.exports = {
  mode: isDev ? "development" : "production",
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
        test: /\.scss$/,
        use: [
          // Creates `style` nodes from JS strings
          isDev ? "style-loader" : extracter.loader,
          // Translates CSS into CommonJS
          {
            loader: "css-loader",
            options: {
              sourceMap: isDev,
            },
          },
          // Compile sass to css
          {
            loader: "sass-loader",
            options: {
              sourceMap: isDev,
            },
          },
        ],
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js", ".scss"],
  },
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "src", "build"),
  },
  plugins: [
    new hwp({
      template: path.resolve(__dirname, "src", "app", "static", "index.html"),
    }),
    new extracter({ filename: "style.css" }),
  ],
  optimization: {
    minimize: true,
    minimizer: [
      new terser({
        extractComments: true,
      }),
      new cssminifier({
        filename: isDev ? "[name].css" : "[name].[hash].css",
        chunkFilename: isDev ? "[id].css" : "[id].[hash].css",
      }),
    ],
  },
};
