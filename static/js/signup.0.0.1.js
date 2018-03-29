(function () {
    var $form = $('#signupForm');
    $form.on('submit', function (e) {
        e.preventDefault();

        var $email = $('.email');
        var $pwd = $('.pwd');
        var $repeat = $('.repeat');
        var emailVal = $email.val().trim();
        var isCorrectEmail = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/.test(emailVal);

        if (!isCorrectEmail) {
            App.notify('邮箱输入有误', 'error');
            return;
        }

        if ($pwd.val() !== $repeat.val()) {
            App.notify('两次密码不匹配', 'error');
            return;
        }

        $.ajax({
            url: '/api/v1/signup',
            type: 'POST',
            success: function (res) {
                console.log(res);
            }
        });
    });
})();
