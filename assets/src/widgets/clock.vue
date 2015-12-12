<template lang="jade">
div
    h1 {{date}}
    h2 {{time}}
</template>

<script>
var Widget = require('./widget.js');
export default Widget.extend({
    data() {
        return {
            date: 0,
            time: 0
        }
    }, ready () {
        function formatTime(i) {
            if (i < 10) {
                return "0" + i;
            } else {
                return i;
            }
        };
        setInterval( () => {
            var today = new Date();
            var h = today.getHours();
            var m = today.getMinutes();
            var s = today.getSeconds();
            var m = formatTime(m);
            var s = formatTime(s);
            this.date = today.toLocaleDateString();
            this.time = h + ":" + m + ":" + s;
        }, 500);
    }
});
</script>

<style lang="stylus" scoped>
</style>
