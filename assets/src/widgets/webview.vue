<template lang="jade">
div
    iframe(:src="data.url", frameborder=0)
</template>

<script>
var Widget = require('./widget.js');
export default Widget.extend({});
</script>

<style lang="stylus" scoped>
div
    padding 0!important

iframe
    margin 0
    padding 0
    width 100%
    height 100%
</style>
