// var path = require('path');

// function resolve(dir) {
//   return path.join(__dirname, dir)
// }
// const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');

module.exports = {
  assetsDir: 'webPage',
  devServer: {
    proxy: {
      '/api': {
        target: 'http://10.8.2.28:10088/',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        },
      },
    }
  },

  chainWebpack:  config => {
    config.module
      .rule('vue')
      .test(/\.vue$/)
      .use('babel')
        .loader('vue-loader')
        .loader('iview-loader')
        .options({
          prefix: false
        })
        .end();
    // if (process.env.NODE_ENV === 'production') {
    //   return {
    //     plugins: [
    //       new BundleAnalyzerPlugin()
    //     ]
    //   }
    // }
    

  },



  pwa: {
    iconPaths: {
      favicon32: 'favicon.png',
      favicon16: 'favicon.png',
      appleTouchIcon: 'favicon.png',
      maskIcon: 'favicon.png',
      msTileImage: 'favicon.png'
    }
  },


  outputDir: 'gateway'


}
