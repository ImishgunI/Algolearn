document.addEventListener("DOMContentLoaded", () => {
    loadUserProfile();
    loadStats();
});

// Получение данных профиля
async function loadUserProfile() {
    const email = localStorage.getItem("email");

    const response = await fetch(`http://localhost:8080/user/profile?email=${email}`);
    const data = await response.json();

    document.getElementById("usernameInput").value = data.first_name;
    document.getElementById("lastnameInput").value = data.last_name;
    document.getElementById("emailInput").value = data.email;
    document.getElementById("roleInput").value = data.role;
}

// Сохранение изменений профиля
async function saveProfile() {
    const username = document.getElementById("usernameInput").value;
    const lastname = document.getElementById("lastnameInput").value;
    const password = document.getElementById("passwordInput").value;
    const email = localStorage.getItem("emailInput");

    await fetch("http://localhost:8080/user/profile/update", {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ email, username, lastname, password })
    });

    alert("Данные обновлены");
}

// Загрузка статистики
async function loadStats() {
    const email = localStorage.getItem("email");
    const response = await fetch(`http://localhost:8080/user/progress?email=${email}`);
    const data = await response.json();

    document.getElementById("progressCompleted").textContent = data.completed;
    document.getElementById("favoritesCount").textContent = data.favorites;
}

// Выход
function logout() {
    localStorage.clear();
    window.location.href = "login.html";
}
