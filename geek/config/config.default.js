/* eslint valid-jsdoc: "off" */

'use strict';
const path = require("path")
/**
 * @param {Egg.EggAppInfo} appInfo app info
 */
module.exports = appInfo => {
    /**
     * built-in config
     * @type {Egg.EggAppConfig}
     **/
    const config = {};
    config.security = {
        csrf: false,
    };
    config.cluster = {
        listen: {
            path: '',
            port: 7001,
            // hostname: 'geekbang.org',
        }
    };


    // use for cookie sign key, should change to your own and keep security
    config.keys = appInfo.name + '_1554791649832_867';

    // add your middleware config here
    config.middleware = [];

    // add your user config here
    const userConfig = {
        // myAppName: 'egg',
    };
    const assets = {
        publicPath: path.join(appInfo.baseDir, 'app/view/index.html'),
        devServer: {
            debug: true,
            port: 8000,
            env: {
                BROWSER: 'none',
                ESLINT: 'none',
                SOCKET_SERVER: 'http://geekbang.org:8000',
                PUBLIC_PATH: 'http://geekbang.org:8000',
            },
        },
    }
    const view = {
        defaultViewEngine: 'nunjucks',
        mapping: {
            '.html': 'nunjucks',
        }
    }

    return {
        ...config,
        ...userConfig,
        assets,
        view,
    };
};
