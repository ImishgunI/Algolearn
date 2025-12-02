document.addEventListener("DOMContentLoaded", () => {
    loadUserProfile();
    loadStats();
});

// Получение данных профиля
async function loadUserProfile() {
    const email = localStorage.getItem("email");

    const response = await fetch(`http://localhost:8080/user/profile?email=${email}`);
    const data = await response.json();

    document.getElementById("usernameInput").value = data.username || "";
    document.getElementById("emailInput").value = data.email;
}

// Сохранение изменений профиля
async function saveProfile() {
    const username = document.getElementById("usernameInput").value;
    const email = localStorage.getItem("email");

    await fetch("http://localhost:8080/user/profile/update", {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ email, username })
    });

    alert("Данные обновлены");
}

// Загрузка статистики
async function loadStats() {
    const email = localStorage.getItem("email");
    const response = await fetch(`http://localhost:8080/progress?email=${email}`);
    const data = await response.json();

    document.getElementById("progressCompleted").textContent = data.completed;
    document.getElementById("progressTotal").textContent = data.total;
    document.getElementById("favoritesCount").textContent = data.favorites;
    document.getElementById("topicsViewed").textContent = data.viewed_topics;
}

// Выход
function logout() {
    localStorage.clear();
    window.location.href = "login.html";
}
