document.addEventListener('DOMContentLoaded', function() {
    document.body.style.opacity = '0';
    document.body.style.transition = 'opacity 0.8s ease';

    setTimeout(() => {
        document.body.style.opacity = '1';
    }, 100);

    document.querySelectorAll('nav a').forEach(link => {
        link.addEventListener('click', function(e) {
            if (this.hostname === window.location.hostname) {
                e.preventDefault();
                const targetUrl = this.href;
                document.body.style.opacity = '0';
                setTimeout(() => {
                    window.location.href = targetUrl;
                }, 400);
            }
        });
    });
});
