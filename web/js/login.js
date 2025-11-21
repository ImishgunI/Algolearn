document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {"Content-Type": "application/json"},
            credentials: 'include',
            body: JSON.stringify({
                email: email,
                password: password
            })
        });

        if (response.ok) {
            const data = await response.json();
            alert('Вход успешен!');
            localStorage.setItem("isLoggedIn", "true");
            localStorage.setItem("first_name", data.first_name);
            localStorage.setItem("last_name", data.last_name);
            localStorage.setItem("email", data.email);
            localStorage.setItem("role", data.role);
            window.location.href = 'index.html'
        } else {
            const error = await response.text();
            alert('Ошибка: ' + error);
        }
    } catch (err) {
        alert('Нет связи с сервером');
    }
});