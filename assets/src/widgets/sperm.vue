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
        var pathLength = 12;
        var degrees = 180 / Math.PI;
        var head;
        var tail;
        var redraw = (data) => {
            var g = this.svg.selectAll('g').data(data);
            var appenedG = g.enter().append('g');
            appenedG.append('ellipse')
                .attr('rx', 6.5)
                .attr('ry', 4)
                .style('fill', function(d) {
                    return d.color;
                });
            appenedG.append('path')
                .datum(function(d) {
                    return d.path.slice(0, 3);
                })
                .style('fill', 'none')
                .style('stroke', function(d, i) {
                    var sperm = data[i];
                    return sperm.color;
                })
                .style('stroke-linecap', 'round')
                .style('stroke-width', '4px');
            appenedG.append('path')
                .datum(function(d) {
                    return d.path;
                })
                .style('fill', 'none')
                .style('stroke', function(d, i) {
                    var sperm = data[i];
                    return sperm.color;
                })
                .style('stroke-linecap', 'round')
                .style('stroke-width', '2px');

            head = g.selectAll('ellipse');
            tail = g.selectAll('path');
            if (data.length === 1) {
                d3.timer(() => {
                    for (var i = -1; ++i < data.length;) {
                        var sperm = data[i],
                            path = sperm.path,
                            dx = sperm.vx,
                            dy = sperm.vy,
                            x = path[0][0] += dx,
                            y = path[0][1] += dy,
                            speed = Math.sqrt(dx * dx + dy * dy),
                            count = speed * 10,
                            k1 = -5 - speed / 3;

                        // bounce off the walls.
                        if (x < 0 || x > this.width) sperm.vx *= -1;
                        if (y < 0 || y > this.height) sperm.vy *= -1;

                        // swim!
                        for (var j = 0; ++j < pathLength;) {
                            var vx = x - path[j][0],
                                vy = y - path[j][1],
                                k2 = Math.sin(((sperm.count += count) + j * 3) / 300) / speed;
                            path[j][0] = (x += dx / speed * k1) - dy * k2;
                            path[j][1] = (y += dy / speed * k1) + dx * k2;
                            speed = Math.sqrt((dx = vx) * dx + (dy = vy) * dy);
                        }
                    };

                    head.attr('transform', (d)=> {
                        var degree = Math.atan2(d.vy, d.vx) * degrees;
                        return `translate(${d.path[0]})rotate(${degree})`;
                    });
                    tail.attr('d', function(d) {
                        return 'M' + d.join('L');
                    });
                });
            };
        }

        var sperms = [];
        this.$watch('data.color', function(newColor) {
            var x = Math.random() * this.width,
                y = Math.random() * this.height;
            var sperm =  {
                vx: Math.random() * 2 - 1,
                vy: Math.random() * 2 - 1,
                path: d3.range(pathLength).map(function() { return [x, y]; }),
                count: 0,
                color: newColor
            };
            sperms.push(sperm);
            redraw(sperms);
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
