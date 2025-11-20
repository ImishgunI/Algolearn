document.getElementById('registerForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const firstName = document.getElementById('firstName').value.trim();
    const lastName = document.getElementById('lastName').value.trim();
    const email = document.getElementById('email').value.trim();
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    if (password !== confirmPassword) {
        alert('Пароли не совпадают');
        return;
    }

    try {
        const response = await fetch('http://localhost:8080/register', {
            method: 'POST',
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                firstName: firstName,
                lastName: lastName,
                email: email,
                passwordHash: password,
            })
        });

        if (response.ok) {
            alert('Регистрация прошла успешно!');
            window.location.href = 'login.html';
        } else {
            alert('Ошибка регистрации: ' + response.status);
        }
    } catch (error) {
        console.error('Ошибка при запросе:', error);
        alert('Не удалось подключиться к серверу');
    }
});
