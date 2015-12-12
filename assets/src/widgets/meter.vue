<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{getConfig('title', '')}}
    canvas(:width="width", :height="height")
    h2.value {{ data.value | shortenedNumber }}
    p.info(:style="config&&config.info_style") {{getConfig('info', '')}}
    p.updated {{updated}}
</template>

<script>
var Chart = require('chart.js');

var Widget = require('./widget.js');
export default Widget.extend({
    created() {
        if (!this.data.value) {
            this.$set('data.value', 0);
        };
    }, ready() {
        var max = this.getConfig('max', 100);

        var canvas = this.$el.querySelector('canvas');
        var ctx = canvas.getContext("2d");
        var chart = new Chart(ctx).Doughnut([
                {
                    value: this.data.value,
                    color: this.getConfig('color', 'white'),
                    label: this.getConfig('title', "")
                }, {
                    value: max - this.data.value,
                    color: this.getConfig('color', '#662b4c'),
                    label: this.getConfig('title', "max")
                }, {
                    value: max/2.0,
                    color: this.getConfig('color', "transparent"),
                    label: this.getConfig('title', "empty")
                }
            ], {
                showTooltips: false,
                segmentShowStroke: false,
                percentageInnerCutout: 60
            });
        this.$watch('data.value', function(newVal, oldVal) {
            chart.segments[0].value = newVal;
            chart.segments[1].value = max - newVal;
            chart.update();
        });
    }
});
</script>

<style lang="stylus" scoped>
canvas
    transform rotate(240deg)
    -ms-transform rotate(240deg)
    -webkit-transform rotate(240deg)
    -o-transform rotate(240deg)
    -moz-transform rotate(240deg)

.title
    color rgba(255, 255, 255, 0.7)

.value
    font-size 30px
    position absolute
    top 50%
    left 0
    right 0

.info
    color rgba(255, 255, 255, 0.3)

.updated
    color rgba(0, 0, 0, 0.3)
</style>
