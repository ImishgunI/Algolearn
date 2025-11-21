document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
  anchor.addEventListener("click", function (e) {
    e.preventDefault();
    const target = document.querySelector(this.getAttribute("href"));
    if (target) {
      target.scrollIntoView({
        behavior: "smooth",
        block: "start",
      });
    }
  });
});

window.addEventListener("scroll", function () {
  const header = document.querySelector(".header");
  if (window.scrollY > 100) {
    header.style.background = "rgba(255, 255, 255, 0.95)";
    header.style.backdropFilter = "blur(10px)";
    header.style.boxShadow = "0 2px 10px rgba(0,0,0,0.1)";
  } else {
    header.style.background = "#FFFFFF";
    header.style.backdropFilter = "none";
    header.style.boxShadow = "none";
  }
});

// Плавная прокрутка
function scrollToSection(selector) {
  const el = document.querySelector(selector);
  if (el) el.scrollIntoView({ behavior: "smooth", block: "start" });
}

// Переходы по курсам
document.querySelectorAll(".btn-start").forEach((btn) => {
  btn.addEventListener("click", function () {
    if (!localStorage.getItem("isLoggedIn")) {
      alert("Войдите в аккаунт, чтобы начать курс");
      window.location.href = "login.html";
      return;
    }

    const page = this.dataset.page;
    const topic = this.dataset.topic;
    window.location.href = `${page}.html#${topic}`;
  });
});

function updateAuthButtons() {
  const authButtons = document.getElementById("authButtons");
  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";

  if (!authButtons) return;

  if (isLoggedIn) {
    const role = localStorage.getItem("role");
    console.log(role)
    let buttonsHTML = `
      <a href="dashboard.html" class="btn">Личный кабинет</a>
      <button onclick="logout()" class="btn logout">Выйти</button>
    `;
    if (role === "admin") {
      buttonsHTML += `<button onclick="window.location.href='admin-dashboard.html'">Админ-панель</button>`;
    }
    authButtons.innerHTML = buttonsHTML;
  } else {
    authButtons.innerHTML = `
            <a href="login.html" class="btn">Вход</a>
            <a href="registration.html" class="btn">Регистрация</a>
        `;
  }
}

function logout() {
  localStorage.setItem("isLoggedIn", false);
  localStorage.removeItem("role");
  localStorage.removeItem("email");
  alert("Вы вышли из аккаунта");
  updateAuthButtons();
}

window.logout = logout;

document.addEventListener("DOMContentLoaded", function () {
  updateAuthButtons();
});
