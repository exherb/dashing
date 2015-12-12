<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{getConfig('title', '')}}
    h2.value(:style="config&&config.value_style")
        | {{config.prefix}}
        span(v-animate-number="data.value")
        | {{config.suffix}}
    p.diff
        i(class="fa", :class="arrow", v-if="getConfig('show_arrow', true)")
        span(v-if="getConfig('show_diff', true)" v-animate-number="diff")
        span(v-if="getConfig('show_diff', true) && getConfig('percentage', true)") %
    p.info(:style="config&&config.info_style") {{getConfig('info', '')}}
    p.updated {{updated}}
</template>

<script>
var Vue = require('vue');
var Widget = require('./widget.js');
export default Widget.extend({
    created() {
        if (!this.data.value) {
            this.$set('data.value', 0);
        };
    }, data() {
        return {
            arrow: '',
            diff: 0
        };
    }, ready () {
        this.$watch("data.value", function(newVal, oldVal) {
            if (oldVal === undefined) {
                return ;
            };

            if (newVal > oldVal) {
                this.arrow = 'fa-arrow-up';
            } else {
                this.arrow = 'fa-arrow-down';
            };

            var diff = Math.abs(newVal - oldVal);
            if (this.getConfig('percentage', true)) {
                if (oldVal === 0) {
                    diff = 100;
                } else {
                    diff = Math.round(diff / oldVal * 100);
                };
            };
            this.diff = diff;
        });
    }
});
</script>

<style lang="stylus" scoped>
value-color = #fff

.title
    color rgba(255, 255, 255, 0.7)

.info
    color rgba(255, 255, 255, 0.7)

.value
    color value-color

.diff
    font-weight 500
    font-size 30px
    color value-color

.updated
    color rgba(0, 0, 0, 0.3)
</style>
