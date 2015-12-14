<template lang="jade">
.gridster
    ul
        li(v-for="(index, widget) in dashboard.widgets", data-row="{{rowAtIndex(index)}}", data-col="{{index%dashboard.columns+1}}", data-sizex="{{widget.sizex}}", data-sizey="{{widget.sizey}}", :style="widget.style")
            component(:is="registerComponent(widget.view)",
                      :id="widget.id",
                      :data="widget.data || {}",
                      :config="widget.config",
                      class="widget")
            i(class="fa icon-background",
              :class="[widget.icon]",
              :style="{ 'margin-top': dashboard.size.height/3.4 + 'px' }",
              v-if="widget['icon']",)
</template>

<script>

var Vue = require('vue');

var shortenedNumber = function (value) {
    if (isNaN(value)) {
        return value;
    };
    if (value >=1000000000 ) {
        return (value / 1000000000).toFixed(1) + 'B';
    };
    if (value >=1000000 ) {
        return (value / 1000000).toFixed(1) + 'M';
    };
    if (value >=1000 ) {
        return (value / 1000).toFixed(1) + 'K';
    };
    return value.toFixed(0);
};
Vue.filter('shortenedNumber', shortenedNumber);


Vue.directive('animate-number', {
    update: function(newVal, oldVal) {
        this.stopTimer();

        if (isNaN(newVal) || isNaN(oldVal)) {
            this.el.innerHTML = newVal;
            return ;
        };

        var step = (newVal - oldVal)/50;
        var current = oldVal;
        this.timer = setInterval(() => {
            if ((current > newVal && step > 0) ||
                (current < newVal && step < 0)) {
                this.stopTimer();
                current = newVal;
            } else {
                current += step;
            };
            this.el.innerHTML = shortenedNumber(current);
        }, 10);
    }, stopTimer: function() {
        if (this.timer) {
            clearInterval(this.timer);
        };
    }, unbind: function() {
        this.stopTimer();
    }
});

Vue.config.debug = true;

export default {
    computed: {
        dashboardName: function() {
            var path = window.location.pathname;
            return path.substr(path.lastIndexOf('/') + 1);
        },
        dashboard: function() {
            var dashboard = require('../dashboards/' + this.dashboardName + '.json');
            for (var i = 0; i < dashboard.widgets.length; i++) {
                var widget = dashboard.widgets[i];
                if (!widget.config) {
                    widget.config = {};
                };
                if (!widget.config.width) {
                    widget.config.width = widget.sizex*dashboard.size.width - 12*2;
                };
                if (!widget.config.height) {
                    widget.config.height = widget.sizey*dashboard.size.height - 12*2 - 42 - 12 - 21*2 - 42;
                };
            };
            return dashboard;
        }
    }, ready () {
        var ul = this.$el.firstElementChild||this.$el.firstChild;
        $(ul).gridster({
            widget_margins: [this.dashboard.margin.x, this.dashboard.margin.y],
            widget_base_dimensions: [this.dashboard.size.width, this.dashboard.size.height],
            max_cols: this.dashboard.columns
        });

        var closeTimer;
        this.listenEvents();
        window.onfocus = ()=> {
            if (closeTimer) {
                clearTimeout(closeTimer);
                closeTimer = null;
            };
            if (!this.source) {
                this.listenEvents();
            };
        };
        window.onblur = ()=> {
            closeTimer = setTimeout(() => {
                closeTimer = null;
                this.source.close();
                this.source = null;
            }, 3000);
        };
    }, methods: {
        registerComponent: function(widgetView) {
            if (Vue.component(widgetView)) {
                return widgetView;
            }
            var component = require('./widgets/' + widgetView + '.vue');
            Vue.component(widgetView, component);
            return widgetView;
        }, rowAtIndex: function(index) {
            var x = 0;
            for (var i = 0; i < index; i++) {
                var widget = this.dashboard.widgets[i];
                x += parseInt(widget.sizex);
            };
            return Math.floor(x/this.dashboard.columns) + 1;
        }, listenEvents: function() {
            this.source = new EventSource(this.dashboardName + '/events');
            this.source.addEventListener('open', e => {
                console.log("[Dashing]Connected", e);
            });

            this.source.addEventListener('error', e => {
                console.log("[Dashing]Connection error", e);
                if (e.currentTarget.readyState === EventSource.CLOSED) {
                    console.log("[Dashing]Connection closed");
                };
                setTimeout(() => {
                    window.location.reload();
                }, 5*60*1000);
            });

            this.source.addEventListener('widget', (e) => {
                var timestamp = new Date();
                var hours = timestamp.getHours()
                var minutes = ("0" + timestamp.getMinutes()).slice(-2);
                var seconds = ("0" + timestamp.getSeconds()).slice(-2);
                var updated = `Last update: ${hours}:${minutes}:${seconds}`;

                var data = JSON.parse(e.data);
                var dataId = data.id;
                delete data['id'];

                for (var i = 0; i < this.$children.length; i++) {
                    var widget = this.$children[i];
                    if (widget.id === dataId) {
                        for (let key in data) {
                            widget.$set('data.' + key, data[key]);
                            widget.$set('updated', updated);
                        };
                    };
                };
            });

            this.source.addEventListener('dashboard', (e) => {
                if (window.location.pathname !== data.id) {
                    return ;
                };
                var data = JSON.parse(e.data);
                if (data.event === 'reload') {
                    window.location.reload(true);
                };
            });
        }
    }
}
</script>

<style lang="stylus">
background-color = #222
widget-background-color = #dc5945
text-color = #fff

html
    font-size 100%
    -webkit-text-size-adjust 100%
    -ms-text-size-adjust 100%

body
    margin 0
    background-color background-color
    color text-color
    font-size 20px
    font-family 'Open Sans', "Helvetica Neue", Helvetica, Arial, sans-serif

h1, h2, h3, h4, h5, p, ul, ol
    margin 0
    padding 0
  
h1
    margin-bottom 12px
    text-align center
    font-size 30px
    font-weight 400

h2
    text-transform uppercase
    font-size 50px
    font-weight: 700

h3
    font-size 25px
    font-weight 600

a
    text-decoration none
    color inherit

b, strong
    font-weight bold

img
    border 0
    -ms-interpolation-mode bicubic
    vertical-align middle

img, object, iframe
    max-width 100%

table
    border-collapse collapse;
    border-spacing 0
    width 100%

td
    vertical-align middle

.icon-background
    width 100%
    height 100%
    position absolute
    left 0
    top 0
    opacity 0.1
    font-size 180px!important
    text-align center

.info
    font-size 15px
    position absolute
    bottom 32px
    left 0
    right 0

.updated
    font-size 15px
    position absolute
    bottom 12px
    left 0
    right 0

.gridster
    width 1900px
    margin 0 auto

.gridster ul
    list-style none

.gridster .gs-w
    background widget-background-color
    display table
    cursor pointer

.gridster .gs-w .widget
    padding 12px 12px
    text-align center
    width 100%
    display table-cell
    vertical-align middle

.gridster .player
    -webkit-box-shadow 3px 3px 5px rgba(0,0,0,0.3)
    box-shadow 3px 3px 5px rgba(0,0,0,0.3)

.gridster .preview-holder
    border none!important
    border-radius 0!important
    background rgba(255,255,255,.2)!important


.chart-legend, .bar-legend, .line-legend, .pie-legend, .radar-legend, .polararea-legend, .doughnut-legend
    list-style-type none
    text-align center
    -webkit-padding-start 0
    -moz-padding-start 0
    padding-left 0
    li
        display inline-block
        white-space nowrap
        position relative
        border-radius 5px
        padding 2px 8px 2px 20px
        font-size smaller
        cursor default
    span
        display block
        position absolute
        left 0
        top 0
        width 15px
        height 15px
        border-radius 5px
</style>
