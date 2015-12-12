<template lang="jade">
div
    h1.title(:style="config&&config.title_style") {{config.title}}
    .comment(v-show="!hidding" transition="fade")
        h3(:style="config&&config.text_style")
            img(:src="comment.avatar")
            span.name {{comment.name}}
        p.comment {{comment.quote}}
    p.info(:style="config&&config.info_style") {{data.info}}
</template>

<script>
var Widget = require('./widget.js');
export default Widget.extend({
    data() {
        return {
            hidding: false,
            index: 0
        }
    }, computed: {
        comment: function() {
            return this.data.comments[this.index];
        }
    }, ready() {
        setInterval(() => {
            this.hidding = true;
            setTimeout( () => {
                if (this.index + 1 >= this.data.comments.length ) {
                    this.index = 0;
                } else {
                    this.index += 1;
                };
                this.hidding = false;
            }, 300);
        }, 5000);
    }
});
</script>

<style lang="stylus" scoped>
.title
    color rgba(255, 255, 255, 0.7)
    margin-bottom 15px

.name
    padding-left 5px

.info
    color rgba(255, 255, 255, 0.7)

.fade-transition
    transition all 0.3s ease
    visibility visible
    opacity 1

.fade-enter, .fade-leave
    opacity 0
    visibility hidden
</style>
