require('gridster');
require('font-awesome-webpack');

var Vue = require('vue');
var App = require('./app.vue');

new Vue({
  el: 'body',
  components: {
    app: App
  }
});
