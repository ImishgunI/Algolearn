// Проверка прав администратора
function checkAdminAccess() {
    const role = localStorage.getItem('role');
    const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
    
    if (!isLoggedIn || role !== 'admin') {
        alert('Доступ запрещен. Требуются права администратора.');
        window.location.href = 'index.html';
        return false;
    }
    
    document.getElementById('adminName').textContent = 
        localStorage.getItem('firstName') + ' ' + localStorage.getItem('lastName');
    
    return true;
}

// Инициализация админ-панели
document.addEventListener('DOMContentLoaded', function() {
    if (!checkAdminAccess()) return;
    
    updateAuthButtons();
    loadStats();
    showAllUsers();
    loadAllLessons();
    loadPendingComments();
});

// ==================== СТАТИСТИКА ====================
async function loadStats() {
    try {
        const response = await fetch('http://localhost:8080/admin/stats');
        if (response.ok) {
            const stats = await response.json();
            displayStats(stats);
        } else {
            console.error('Ошибка загрузки статистики');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
        // Заглушка для демонстрации
        displayStats({
            usersCount: 150,
            lessonsCount: 25,
            commentsCount: 430,
            tasksCount: 75
        });
    }
}

function displayStats(stats) {
    document.getElementById('usersCount').textContent = stats.usersCount;
    document.getElementById('lessonsCount').textContent = stats.lessonsCount;
    document.getElementById('commentsCount').textContent = stats.commentsCount;
    document.getElementById('tasksCount').textContent = stats.tasksCount;
}

// ==================== УПРАВЛЕНИЕ ПОЛЬЗОВАТЕЛЯМИ ====================
async function showAllUsers() {
    try {
        const response = await fetch('http://localhost:8080/admin/users');
        if (response.ok) {
            const data = await response.json();
            displayUsers(data.users);
        } else {
            console.error('Ошибка загрузки пользователей');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
        // Заглушка для демонстрации
        const demoUsers = [
            { id: 1, first_name: 'Иван', last_name: 'Петров', email: 'ivan@example.com', role: 'student' },
            { id: 2, first_name: 'Мария', last_name: 'Сидорова', email: 'maria@example.com', role: 'manager' },
            { id: 3, first_name: 'Админ', last_name: 'Админов', email: 'admin@example.com', role: 'admin' }
        ];
        displayUsers(demoUsers);
    }
}

function displayUsers(users) {
    const usersList = document.getElementById('usersList');
    
    if (users.length === 0) {
        usersList.innerHTML = '<p>Пользователи не найдены</p>';
        return;
    }

    usersList.innerHTML = users.map(user => `
        <div class="user-item">
            <div class="user-info">
                <strong>${user.first_name} ${user.last_name}</strong>
                <span class="user-email">${user.email}</span>
                <span class="user-role role-${user.role}">${getRoleText(user.role)}</span>
            </div>
            <div class="action-buttons">
                <button onclick="editUser(${user.id})" class="btn-primary">✏️</button>
                <button onclick="deleteUser(${user.id})" class="btn-danger">🗑️</button>
            </div>
        </div>
    `).join('');
}

function searchUsers() {
    const searchTerm = document.getElementById('userSearch').value.toLowerCase();
    // В реальном приложении здесь будет запрос к серверу
    console.log('Поиск пользователей:', searchTerm);
}

function showAddUserForm() {
    document.getElementById('userModalTitle').textContent = 'Добавить пользователя';
    document.getElementById('userForm').reset();
    document.getElementById('userId').value = '';
    document.getElementById('userPassword').required = true;
    openModal('userModal');
}

async function editUser(userId) {
    try {
        const response = await fetch(`http://localhost:8080/admin/users/${userId}`);
        if (response.ok) {
            const user = await response.json();
            document.getElementById('userModalTitle').textContent = 'Редактировать пользователя';
            document.getElementById('userId').value = user.id;
            document.getElementById('userFirstName').value = user.first_name;
            document.getElementById('userLastName').value = user.last_name;
            document.getElementById('userEmail').value = user.email;
            document.getElementById('userRole').value = user.role;
            document.getElementById('userPassword').required = false;
            document.getElementById('userPassword').placeholder = 'Оставьте пустым, если не меняется';
            openModal('userModal');
        }
    } catch (err) {
        console.error('Ошибка загрузки пользователя:', err);
    }
}

async function deleteUser(userId) {
    if (!confirm('Вы уверены, что хотите удалить пользователя?')) return;
    
    try {
        const response = await fetch(`http://localhost:8080/admin/users/${userId}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            alert('Пользователь удален');
            showAllUsers();
        } else {
            alert('Ошибка удаления пользователя');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
}

// ==================== УПРАВЛЕНИЕ УРОКАМИ ====================
async function loadAllLessons() {
    try {
        const response = await fetch('http://localhost:8080/admin/lessons');
        if (response.ok) {
            const data = await response.json();
            displayLessons(data.lessons);
        } else {
            console.error('Ошибка загрузки уроков');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
}

function displayLessons(lessons) {
    const lessonsList = document.getElementById('lessonsList');
    
    if (!lessons || lessons.length === 0) {
        lessonsList.innerHTML = '<p>Уроки не найдены</p>';
        return;
    }

    lessonsList.innerHTML = lessons.map(lesson => `
        <div class="lesson-item">
            <div class="lesson-info">
                <strong>${lesson.title}</strong>
                <div>Категория: ${lesson.category} • Сложность: ${getDifficultyText(lesson.difficulty)}</div>
            </div>
            <div class="action-buttons">
                <button onclick="editLesson(${lesson.id})" class="btn-primary">✏️</button>
                <button onclick="deleteLesson(${lesson.id})" class="btn-danger">🗑️</button>
            </div>
        </div>
    `).join('');
}

function showAddLessonForm() {
    document.getElementById('lessonModalTitle').textContent = 'Добавить урок';
    document.getElementById('lessonForm').reset();
    document.getElementById('lessonId').value = '';
    openModal('lessonModal');
}

async function editLesson(lessonId) {
    try {
        const response = await fetch(`http://localhost:8080/admin/lessons/${lessonId}`);
        if (response.ok) {
            const lesson = await response.json();
            document.getElementById('lessonModalTitle').textContent = 'Редактировать урок';
            document.getElementById('lessonId').value = lesson.id;
            document.getElementById('lessonTitle').value = lesson.title;
            document.getElementById('lessonCategory').value = lesson.category;
            document.getElementById('lessonDifficulty').value = lesson.difficulty;
            document.getElementById('lessonContent').value = lesson.content;
            openModal('lessonModal');
        }
    } catch (err) {
        console.error('Ошибка загрузки урока:', err);
    }
}

async function deleteLesson(lessonId) {
    if (!confirm('Вы уверены, что хотите удалить урок?')) return;
    
    try {
        const response = await fetch(`http://localhost:8080/admin/lessons/${lessonId}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            alert('Урок удален');
            loadAllLessons();
        } else {
            alert('Ошибка удаления урока');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
}

// ==================== МОДЕРАЦИЯ КОММЕНТАРИЕВ ====================
async function loadAllComments() {
    try {
        const response = await fetch('http://localhost:8080/admin/comments');
        if (response.ok) {
            const data = await response.json();
            displayComments(data.comments);
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
}

function displayComments(comments) {
    const commentsList = document.getElementById('commentsList');
    
    if (!comments || comments.length === 0) {
        commentsList.innerHTML = '<p>Комментарии не найдены</p>';
        return;
    }

    commentsList.innerHTML = comments.map(comment => `
        <div class="comment-item">
            <div class="comment-info">
                <strong>${comment.first_name} ${comment.last_name}</strong>
                <div>Урок: ${comment.lesson_title}</div>
                <p>${comment.content}</p>
                <small>${formatDateReadable(comment.created_at)}</small>
            </div>
            <div class="action-buttons">
                <button onclick="deleteComment(${comment.id})" class="btn-danger">🗑️</button>
            </div>
        </div>
    `).join('');
}

function formatDateReadable(dateString) {
    const date = new Date(dateString);
    return date.toLocaleString('ru-RU', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
    });
}

async function deleteComment(commentId) {
    if (!confirm('Удалить комментарий?')) return;
    
    try {
        const response = await fetch(`http://localhost:8080/admin/comments/${commentId}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            alert('Комментарий удален');
            loadAllComments();
        } else {
            alert('Ошибка удаления комментария');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
}

// ==================== ОБРАБОТКА ФОРМ ====================
document.getElementById('userForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const userId = document.getElementById('userId').value;
    const userData = {
        first_name: document.getElementById('userFirstName').value,
        last_name: document.getElementById('userLastName').value,
        email: document.getElementById('userEmail').value,
        role: document.getElementById('userRole').value
    };
    
    const password = document.getElementById('userPassword').value;
    if (password) {
        userData.password = password;
    }
    
    try {
        const url = userId ? `http://localhost:8080/admin/update/users/${userId}` : 'http://localhost:8080/admin/add/users';
        const method = userId ? 'PUT' : 'POST';
        
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        });
        
        if (response.ok) {
            alert(userId ? 'Пользователь обновлен' : 'Пользователь создан');
            closeModal('userModal');
            showAllUsers();
        } else {
            alert('Ошибка сохранения пользователя');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
});

document.getElementById('lessonForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const lessonId = document.getElementById('lessonId').value;
    const lessonData = {
        title: document.getElementById('lessonTitle').value,
        category: document.getElementById('lessonCategory').value,
        difficulty: document.getElementById('lessonDifficulty').value,
        content: document.getElementById('lessonContent').value,
    };
    
    try {
        const url = lessonId ? `http://localhost:8080/admin/update/lessons/${lessonId}` : 'http://localhost:8080/admin/add/lessons';
        const method = lessonId ? 'PUT' : 'POST';
        
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(lessonData)
        });
        
        if (response.ok) {
            alert(lessonId ? 'Урок обновлен' : 'Урок создан');
            closeModal('lessonModal');
            loadAllLessons();
        } else {
            alert('Ошибка сохранения урока');
        }
    } catch (err) {
        console.error('Ошибка сети:', err);
    }
});

// ==================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ====================
function getRoleText(role) {
    const roles = {
        'student': 'Студент',
        'manager': 'Менеджер',
    };
    return roles[role] || role;
}

function getDifficultyText(difficulty) {
    const difficulties = {
        'beginner': 'Начальный',
        'intermediate': 'Средний',
        'advanced': 'Продвинутый'
    };
    return difficulties[difficulty] || difficulty;
}

function openModal(modalId) {
    document.getElementById(modalId).style.display = 'block';
}

function closeModal(modalId) {
    document.getElementById(modalId).style.display = 'none';
}

// Закрытие модальных окон при клике вне контента
window.addEventListener('click', function(event) {
    const modals = document.querySelectorAll('.modal');
    modals.forEach(modal => {
        if (event.target === modal) {
            modal.style.display = 'none';
        }
    });
});

// Закрытие модальных окон на Escape
document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        const modals = document.querySelectorAll('.modal');
        modals.forEach(modal => {
            modal.style.display = 'none';
        });
    }
});