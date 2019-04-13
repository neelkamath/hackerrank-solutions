let button = document.getElementById('btn');

button.addEventListener('click', () => {
    let counter = parseInt(button.innerHTML);
    button.innerHTML = ++counter;
});