window.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("menu-btn");
    const showMenuSVG = document.getElementById("show-menu");
    const closeMenuSVG = document.getElementById("close-menu");
    const mobileMenu = document.getElementById("mobile-menu");
    let menuState = false;

    btn.addEventListener("click", (e) => {
        e.preventDefault();
        if (menuState) {
            showMenuSVG.classList.remove("hidden");
            showMenuSVG.classList.add("block");

            closeMenuSVG.classList.remove("block");
            closeMenuSVG.classList.add("hidden");

            mobileMenu.classList.add("hidden");
            menuState = false;
        } else {
            showMenuSVG.classList.remove("block");
            showMenuSVG.classList.add("hidden");

            closeMenuSVG.classList.remove("hidden");
            closeMenuSVG.classList.add("block");

            mobileMenu.classList.remove("hidden");
            menuState = true;
        }
    });
});
