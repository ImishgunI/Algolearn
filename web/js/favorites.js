document.addEventListener('DOMContentLoaded', function () {
    const header = document.querySelector('.header');
    const favoritesContainer = document.querySelector('.favorites-container');
        
    if (header && favoritesContainer) {
        const headerHeight = header.offsetHeight;
        favoritesContainer.style.paddingTop = (headerHeight + 20) + 'px';
    }
    showAllFavorites();
});

async function showAllFavorites() {
    try {
        const token = localStorage.getItem("isLoggedIn");
        const email = localStorage.getItem("email");
        if (!token) {
            alert("Войдите в аккаунт, чтобы просматривать избранные курсы.");
            return;
        }

        const response = await fetch(`http://localhost:8080/favorites?email=${email}`);
        if (response.ok) {
            const data = await response.json();
            displayFavorites(data.favorites);
        } else {
            console.error("Ошибка загрузки избранных курсов");
        }
    } catch (err) {
        console.error("Ошибка сети:", err);
    }
}

function displayFavorites(favorites) {
    const favoritesList = document.getElementById("favoritesList");

    if (!favorites || favorites.length === 0) {
        favoritesList.innerHTML = "<p>У вас нет избранных курсов.</p>";
        return;
    }

    favoritesList.innerHTML = favorites.map(favorite => `
        <div class="favorite-item">
            <h3>${favorite.lesson_title}</h3>
            <button onclick="removeFromFavorites('${favorite.lesson_id}')">Удалить из избранного</button>
        </div>
    `).join("");
}
async function removeFromFavorites(lessonID) {
    try {
        const response = await fetch(`http://localhost:8080/favorites/remove/${lessonID}`, {
            method: 'DELETE',
        });

        if (response.ok) {
            await showAllFavorites();
        } else {
            console.error('Ошибка удаления курса из избранного');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
}