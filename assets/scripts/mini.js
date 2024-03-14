setTimeout(function () {
    lightGallery(
        document.querySelector('.mini__gallery'), {
            speed: 500,
            plugins: [lgZoom, lgThumbnail],
        }
    );
}, 1000);

document.querySelector('.back').addEventListener('click', function (e) {
    e.preventDefault();
    window.history.back();
});