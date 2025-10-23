document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify({ email, password })
        });

        if (response.ok) {
            const data = await response.json();

            // Сохраняем, что пользователь вошёл
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('user_role', data.role || 'Студент');
            localStorage.setItem('user_email', data.email || 'user@example.com');

            alert('Вход успешен!');
            window.location.href = 'dashboard.html';
        } else {
            const error = await response.text();
            alert('Ошибка: ' + error);
        }
    } catch (err) {
        alert('Нет связи с сервером');
    }
});