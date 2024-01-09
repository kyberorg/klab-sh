let titleBox = document.querySelector('title');

window.addEventListener('DOMContentLoaded', (event) => {
    getKonfig();
});

function getKonfig() {
    fetch('/konfig')
        .then((response) => {
            if (!response.ok) {
                throw new Error("error grabbing konfig: response failed");
            }
            return response.text();
        })
        .then((konfigJson) => {
            changeTitle(konfigJson)
        })
        .catch((error) => {
            console.error("failed to get konfig: backend disconnected");
        });
}

function changeTitle(konfigJson) {
    console.debug(konfigJson);
    const konfig = JSON.parse(konfigJson);
    titleBox.innerHTML = konfig.title;
}