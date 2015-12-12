<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{data.title}}
    h3(:style="config&&config.text_style")  {{data.text}}
    p.info(:style="config&&config.info_style") {{data.info}}
    p.updated {{updated}}
</template>

<script>
var Widget = require('./widget.js');
export default Widget.extend({});
</script>

<style lang="stylus" scoped>
.title
    color rgba(255, 255, 255, 0.7)

.info
    color rgba(255, 255, 255, 0.7)

.updated
    color rgba(255, 255, 255, 0.7)
</style>
