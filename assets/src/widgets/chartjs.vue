<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{getConfig('title', '')}}
    canvas(:width="width", :height="height")
    p.info(:style="config&&config.info_style") {{{legend}}}
    p.updated {{updated}}
</template>

<script>
var Chart = require('chart.js');

var Widget = require('./widget.js');
export default Widget.extend({
    data() {
        return {
            legend: ""
        }
    }, computed: {
        canvas: function() {
            return this.$el.querySelector('canvas');
        }
    }, ready() {
        var ctx = this.canvas.getContext("2d");
        var options = this.getConfig('options', {showTooltips: false});

        var chartType = this.getConfig('type', 'Line');
        var isSegment = ['PolarArea', 'Pie', 'Doughnut'].indexOf(chartType) != -1;

        var dataConfig = this.getConfig('data_config', {});
        if (isSegment) {
            var notAllowZero = ['Pie', 'Doughnut'].indexOf(chartType) != -1;
            for (var i = 0; i < dataConfig.length; i++) {
                var dateItem = dataConfig[i];
                if (dateItem.value === undefined || (notAllowZero && dateItem.value === 0)) {
                    dateItem.value = notAllowZero?1:0;
                };
            };
        } else {
            for (var i = 0; i < dataConfig.length; i++) {
                var dataset = dataConfig[i];
                if (dataset.data === undefined) {
                    dataset.data = [];
                };
            };
            dataConfig = {
                labels: [],
                datasets: dataConfig
            };
        };
        var chart = new Chart(ctx)[chartType](dataConfig, options);

        this.legend = chart.generateLegend();

        var limit = this.getConfig('limit', 20);
        this.$watch('data.value', function(newValue) {
            if (isSegment) {
                for (var i = 0; i < chart.segments.length; i++) {
                    var segment = chart.segments[i];
                    segment.value = newValue[i];
                };
            } else {
                var dataset = chart.datasets[0];
                var items = dataset.points || dataset.bars;
                if (items.length >= limit) {
                    for (var i = 0; i < chart.datasets.length; i++) {
                        var dataset = chart.datasets[i];
                        var items = dataset.points || dataset.bars;
                        for (var j = 0; j + 1 < items.length; j++) {
                            var item = items[j];
                            var nextItem = items[j+1];
                            item.label = nextItem.label;
                            item.value = nextItem.value;
                        };
                        var lastItem = items[items.length - 1];
                        lastItem.label = this.data.label;
                        lastItem.value = newValue[i];
                    };

                    if (chart.scale.labels) {
                        chart.scale.labels.push(this.data.label);
                        chart.scale.labels.shift();
                        chart.reflow();
                    } else if(chart.scale.xLabels) {
                        chart.scale.xLabels.shift();
                        chart.scale.xLabels.push(this.data.label);
                        chart.scale.fit();
                    };
                } else {
                    chart.addData(newValue, this.data.label);
                };
            };

            chart.update();
        });
    }
});
</script>

<style lang="stylus" scoped>
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
