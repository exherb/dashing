var Vue = require('vue');
var Widget = Vue.extend({
    props: ["id", "config", "data"],
    computed: {
        width: function() {
            return this.config.width;
        },
        height: function() {
            return this.config.height;
        }
    }, methods: {
        getConfig: function(name, defaultValue) {
            if (this.config) {
                var configValue = this.config[name];
                if (configValue !== undefined) {
                    return configValue;
                };
            };
            return defaultValue;
        }
    }
});
module.exports = Widget;
