<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{getConfig('title', '')}}
    svg(width="{{getConfig('width', 340)}}", height="{{getConfig('height', 180)}}")
    p.info(:style="config&&config.info_style") {{{legend}}}
    p.updated {{updated}}
</template>

<script>
var d3 = require('d3');
var Widget = require('./widget.js');
export default Widget.extend({
    ready() {
        var width = getConfig('width', 340),
            height = getConfig('height', 180);

        var n = 100,
            m = 12,
            degrees = 180 / Math.PI;

        var spermatozoa = d3.range(n).map(function() {
          var x = Math.random() * width,
              y = Math.random() * height;
          return {
            vx: Math.random() * 2 - 1,
            vy: Math.random() * 2 - 1,
            path: d3.range(m).map(function() { return [x, y]; }),
            count: 0
          };
        });

        var svgElement = this.$el.querySelector('svg');
        var svg = d3.select(svgElement);
        var g = svg.selectAll("g")
            .data(spermatozoa)
            .enter()
            .append("g");

        var head = g.append("ellipse")
            .attr("rx", 6.5)
            .attr("ry", 4);

        g.append("path")
            .datum(function(d) { return d.path.slice(0, 3); })
            .attr("class", "mid");

        g.append("path")
            .datum(function(d) { return d.path; })
            .attr("class", "tail");

        var tail = g.selectAll("path");

        d3.timer(function() {
          for (var i = -1; ++i < n;) {
            var spermatozoon = spermatozoa[i],
                path = spermatozoon.path,
                dx = spermatozoon.vx,
                dy = spermatozoon.vy,
                x = path[0][0] += dx,
                y = path[0][1] += dy,
                speed = Math.sqrt(dx * dx + dy * dy),
                count = speed * 10,
                k1 = -5 - speed / 3;

            // Bounce off the walls.
            if (x < 0 || x > width) spermatozoon.vx *= -1;
            if (y < 0 || y > height) spermatozoon.vy *= -1;

            // Swim!
            for (var j = 0; ++j < m;) {
              var vx = x - path[j][0],
                  vy = y - path[j][1],
                  k2 = Math.sin(((spermatozoon.count += count) + j * 3) / 300) / speed;
              path[j][0] = (x += dx / speed * k1) - dy * k2;
              path[j][1] = (y += dy / speed * k1) + dx * k2;
              speed = Math.sqrt((dx = vx) * dx + (dy = vy) * dy);
            }
          }

          head.attr("transform", headTransform);
          tail.attr("d", tailPath);
        });

        function headTransform(d) {
          return "translate(" + d.path[0] + ")rotate(" + Math.atan2(d.vy, d.vx) * degrees + ")";
        }

        function tailPath(d) {
          return "M" + d.join("L");
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

ellipse
  fill #fff

path
  fill none
  stroke #fff
  stroke-linecap round

.mid
  stroke-width 4px

.tail
  stroke-width 2px
</style>
