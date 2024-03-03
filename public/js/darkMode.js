
function darkModeMain() {
    const lightSwitch = document.getElementById("light-switch");
    const html = document.querySelector("html");
    lightSwitch.addEventListener("click", (e) => {
        e.preventDefault();
        const path = html.classList.contains("dark") ? "/disable-dark-mode" : "enable-dark-mode";
        fetch(path, { method: "POST" })
            .then(res => res.json())
            .then((res) => {
                if (res.ok === null) {
                    return;
                }
                if (res.ok === true) {
                    return;
                }
            })
            .catch((err) => {
                console.error(err)
            });
        html.classList.toggle("dark");
    });
}

darkModeMain();
