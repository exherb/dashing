<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{getConfig('title', '')}}
    svg(:width="width", :height="height")
    p.info(:style="config&&config.info_style") {{{legend}}}
    p.updated {{updated}}
</template>

<script>
var d3 = require('d3');
var D3Widget = require('./d3widget.js');
export default D3Widget.extend({
    ready() {
        this.drawMap('/maps/china.geo.json', this.getConfig('center', [105, 38]), this.getConfig('scale', 500), '#d8d8d8', '#fefefe', 1, (path) => {
            var datas = [];
            var g = this.svg.append("g");
            var bubules;
            this.$watch('data.value', function(newValue) {
                datas.push({
                    coordinate: this.data.coordinate,
                    value: newValue
                });
                var bubuleGroups = g.selectAll("g").data(datas)
                var appended = bubuleGroups.enter().append("g")
                    .attr("transform", function(d) { return "translate(" + path.projection()(d.coordinate) + ")"; });
                appended.append("circle")
                    .style('fill', 'brown')
                    .style('fill-opacity', '0.5')
                    .attr("r", 0)
                    .transition().attr("r", function(d) { return d.value; });
                appended.append('text')
                    .attr('dx', -6)
                    .attr('dy', 6)
                    .text(function(d) { return d.value; })
                    .style('fill', 'white')
                    .style('font-size', '12px');
                bubules = bubuleGroups.selectAll('circle');
                if (datas.length === 1) {
                    var repeat = (zoom) => {
                        bubules.transition().attr("r", function(d) { 
                            if (zoom) {
                                return d.value * 1.1;
                            } else {
                                return d.value * 0.9;
                            };
                        }).each("end", () => {
                            repeat(!zoom);
                        });
                    };
                    repeat(true);
                };
            });
        });
        
    }, methods: {
        drawMap: function(mapUrl, center, scale, fill, stroke, strokeWidth, callback) {
            var svg = this.svg;
            var projection = d3.geo.mercator()
                .translate([this.width / 2, this.height/2])
                .center(center).scale(scale);
            var path = d3.geo.path().projection(projection);
            d3.json(mapUrl, function(json) {
                var map = svg.append('g')
                    .attr('class', 'map');
                map.selectAll('path')
                        .data(json.features)
                        .enter()
                        .append('path')
                        .attr('d', path)
                        .attr('fill', fill)
                        .attr('stroke', stroke)
                        .attr('stroke-width', strokeWidth);
                callback(path);
            });
        }
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
