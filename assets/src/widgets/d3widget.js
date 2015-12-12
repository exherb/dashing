var d3 = require('d3');
var Widget = require('./widget.js');
var D3Widget = Widget.extend({
    computed: {
        svg: function() {
            var svgElement = this.$el.querySelector('svg');
            return d3.select(svgElement);
        }
    }
});
module.exports = D3Widget;
