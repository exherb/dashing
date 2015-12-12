<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{getConfig('title', '')}}
    svg(width="{{getConfig('width', 340)}}", height="{{getConfig('height', 180)}}")
    p.info(:style="config&&config.info_style") {{{legend}}}
    p.updated {{updated}}
</template>

<script>
var Widget = require('./widget.js');
export default Widget.extend({
    
});
</script>

<style lang="stylus" scoped>
</style>
