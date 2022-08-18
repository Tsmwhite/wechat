'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    WEB_SOCKET_SERVER: '"ws://127.0.0.1:8011/ws"',
})
