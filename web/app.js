document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }
    });
});

window.addEventListener('scroll', function() {
    const header = document.querySelector('.header');
    if (window.scrollY > 100) {
        header.style.background = 'rgba(255, 255, 255, 0.95)';
        header.style.backdropFilter = 'blur(10px)';
        header.style.boxShadow = '0 2px 10px rgba(0,0,0,0.1)';
    } else {
        header.style.background = '#FFFFFF';
        header.style.backdropFilter = 'none';
        header.style.boxShadow = 'none';
    }
});

document.querySelectorAll('.btn-start').forEach(button => {
    button.addEventListener('click', function() {
        const courseTitle = this.closest('.course-card').querySelector('h3').textContent;
        const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';

        if (!isLoggedIn) {
            alert('Войдите в аккаунт, чтобы начать курс');
            window.location.href = 'login.html';
            return;
        }

        alert(`Начинаем курс: ${courseTitle}\nФункционал будет доступен после реализации бэкенда`);
    });
});

function updateAuthButtons() {
    const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
    const authContainer = document.getElementById('authButtons');

    if (!authContainer) return;

    if (isLoggedIn) {
        authContainer.innerHTML = `
            <a href="dashboard.html" class="btn">Личный кабинет</a>
            <a href="#" class="btn logout" onclick="logout()">Выйти</a>
        `;
    } else {
        authContainer.innerHTML = `
            <a href="login.html" class="btn">Вход</a>
            <a href="registration.html" class="btn">Регистрация</a>
        `;
    }
}

function logout() {
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('user_role');
    localStorage.removeItem('user_email');
    alert('Вы вышли из аккаунта');
    updateAuthButtons();
}

window.logout = logout;

document.addEventListener('DOMContentLoaded', function() {
    updateAuthButtons();
});