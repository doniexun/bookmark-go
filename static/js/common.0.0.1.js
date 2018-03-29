(function () {
    function notify(content, type) {
        var classname = type === 'error' ? 'is-danger' : 'is-primary'
        var html = [
            '<div class="notification ' + classname + '">',
            '<button class="delete"></button>',
            content,
            '</div>'
        ];
        $('body').prepend(html.join(''));
        $('.notification .delete').click(function () {
            $('.notification').remove();
        });
    }

    window.App = {
        notify: notify,
    };
})();
