let currentLessons = {};
let currentTopic = null;

const topicMapping = {
    'sorts': {
        'bubble': 'Пузырьковая сортировка',
        'merge': 'Сортировка слиянием', 
        'quick': 'Быстрая сортировка'
    },
    'search': {
        'linear': 'Линейный поиск',
        'binary': 'Бинарный поиск'
    },
    'recursion': {
        'factorial': 'Факториал',
        'fibonacci': 'Числа Фибоначчи'
    },
    'dynamic': {
        'fib': 'Фибоначчи (DP)'
    },
    'trees': {
        'bfs': 'Обход в ширину (BFS)'
    },
    'difficult': {
        'big-o': 'Big O нотация'
    }
};

async function loadLessonsData() {
    try {
        const response = await fetch('http://localhost:8080/lessons');
        if (response.ok) {
            const data = await response.json();
            const lessonsArray = data.lessons || data;
            console.log('Loaded lessons:', lessonsArray);
            
            currentLessons = transformLessonsData(lessonsArray);
            console.log('Transformed lessons:', currentLessons);
            return currentLessons;
        } else {
            console.error('Ошибка загрузки уроков');
            return {};
        }
    } catch (err) {
        console.error('Нет связи с сервером:', err);
        return {};
    }
}

function transformLessonsData(lessons) {
    const topics = {};
    
    // Сначала создаем структуру из mapping'а
    for (const [category, lessonsMap] of Object.entries(topicMapping)) {
        topics[category] = {};
        for (const [key, title] of Object.entries(lessonsMap)) {
            topics[category][key] = {
                name: title,
                theory: '', // временно пусто
                code: getDefaultCode(title),
                data: getDefaultData(title),
                test: getDefaultTest(title),
                visualize: getVisualizationFunction(title)
            };
        }
    }
    
    // Затем наполняем данными из БД
    lessons.forEach(lesson => {
        const category = lesson.category;
        
        // Ищем урок в mapping'е по названию
        if (topics[category]) {
            for (const [key, topic] of Object.entries(topics[category])) {
                if (topic.name === lesson.title) {
                    // Обновляем теорию из БД
                    topic.theory = lesson.content;
                    topic.lesson_id = lesson.id; // если есть id
                    console.log(`Updated topic: ${topic.name} in category ${category}`);
                    break;
                }
            }
        }
    });
    
    return topics;
}

function getDefaultCode(lessonTitle) {
    const codeMap = {
        'Пузырьковая сортировка': `function bubbleSort(arr) {
    const n = arr.length;
    for (let i = 0; i < n - 1; i++) {
        for (let j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
            }
        }
    }
    return arr;
}`,
        'Сортировка слиянием': `function mergeSort(arr) {
    if (arr.length <= 1) return arr;
    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid));
    const right = mergeSort(arr.slice(mid));
    return merge(left, right);
}

function merge(left, right) {
    let result = [], i = 0, j = 0;
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) result.push(left[i++]);
        else result.push(right[j++]);
    }
    return result.concat(left.slice(i)).concat(right.slice(j));
}`,
        'Быстрая сортировка': `function quickSort(arr, low = 0, high = arr.length - 1) {
    if (low >= high) return arr;
    const pivotIdx = partition(arr, low, high);
    quickSort(arr, low, pivotIdx - 1);
    quickSort(arr, pivotIdx + 1, high);
    return arr;
}

function partition(arr, low, high) {
    const pivot = arr[high];
    let i = low;
    for (let j = low; j < high; j++) {
        if (arr[j] < pivot) {
            [arr[i], arr[j]] = [arr[j], arr[i]];
            i++;
        }
    }
    [arr[i], arr[high]] = [arr[high], arr[i]];
    return i;
}`,
        'Линейный поиск': `function linearSearch(arr, target) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === target) return i;
    }
    return -1;
}`,
        'Бинарный поиск': `function binarySearch(arr, target) {
    let left = 0, right = arr.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] === target) return mid;
        if (arr[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}`,
        'Факториал': `function factorial(n) {
    if (n <= 1) return 1;
    return n * factorial(n - 1);
}`,
        'Числа Фибоначчи': `function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}`,
        'Фибоначчи (DP)': `function fibDP(n) {
    const dp = [0, 1];
    for (let i = 2; i <= n; i++) {
        dp[i] = dp[i-1] + dp[i-2];
    }
    return dp[n];
}`,
        'Обход в ширину (BFS)': `function bfs(graph, start) {
    const visited = new Set();
    const queue = [start];
    const result = [];
    while (queue.length) {
        const node = queue.shift();
        if (!visited.has(node)) {
            visited.add(node);
            result.push(node);
            queue.push(...graph[node]);
        }
    }
    return result;
}`,
        'Big O нотация': `// Пример O(n²)
for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
        console.log(i, j);
    }
}`
    };
    
    return codeMap[lessonTitle] || '// Код для этого урока будет добавлен позже';
}

function getDefaultData(lessonTitle) {
    const dataMap = {
        'Пузырьковая сортировка': () => Array.from({ length: 12 }, () => Math.floor(Math.random() * 90) + 10),
        'Сортировка слиянием': () => Array.from({ length: 10 }, () => Math.floor(Math.random() * 100)),
        'Быстрая сортировка': () => Array.from({ length: 10 }, () => Math.floor(Math.random() * 100)),
        'Линейный поиск': () => ({ arr: [12, 5, 8, 3, 9], target: 8 }),
        'Бинарный поиск': () => ({ arr: [2, 3, 5, 7, 11, 13], target: 7 }),
        'Факториал': () => 5,
        'Числа Фибоначчи': () => 6,
        'Фибоначчи (DP)': () => 10,
        'Обход в ширину (BFS)': () => ({
            graph: { A: ["B", "C"], B: ["D"], C: ["E"], D: [], E: [] },
            start: "A"
        }),
        'Big O нотация': () => ({})
    };
    
    return dataMap[lessonTitle] || (() => []);
}


function getDefaultTest(lessonTitle) {
    const testMap = {
        'Пузырьковая сортировка': (fn) => fn([5, 3, 8, 1]).join() === "1,3,5,8",
        'Сортировка слиянием': (fn) => fn([38, 27, 43, 3]).join() === "3,27,38,43",
        'Быстрая сортировка': (fn) => fn([10, 7, 8, 9, 1]).join() === "1,7,8,9,10",
        'Линейный поиск': (fn) => fn([12, 5, 8, 3, 9], 8) === 2,
        'Бинарный поиск': (fn) => fn([2, 3, 5, 7, 11, 13], 7) === 3,
        'Факториал': (fn) => fn(5) === 120,
        'Числа Фибоначчи': (fn) => fn(6) === 8,
        'Фибоначчи (DP)': (fn) => fn(10) === 55,
        'Обход в ширину (BFS)': (fn) => 
            fn({ A: ["B", "C"], B: ["D"], C: ["E"], D: [], E: [] }, "A").join() === "A,B,C,D,E",
        'Big O нотация': (fn) => true
    };
    
    return testMap[lessonTitle] || (() => true);
}

function getVisualizationFunction(lessonTitle) {
    const visMap = {
        'Пузырьковая сортировка': (mode) => visualizeBubble(mode),
        'Сортировка слиянием': (mode) => visualizeMerge(mode),
        'Быстрая сортировка': (mode) => visualizeQuick(mode),
        'Линейный поиск': (mode) => visualizeLinear(mode),
        'Бинарный поиск': (mode) => visualizeBinary(mode),
        'Факториал': (mode) => visualizeRecursion("factorial", mode),
        'Числа Фибоначчи': (mode) => visualizeRecursion("fibonacci", mode),
        'Фибоначчи (DP)': (mode) => visualizeDPFib(mode),
        'Обход в ширину (BFS)': (mode) => visualizeBFS(mode),
        'Big O нотация': (mode) => visualizeBigO(mode)
    };
    
    return visMap[lessonTitle] || (() => {});
}


// === КОММЕНТАРИИ ===
async function addComment(name) {
    const text = document.getElementById("new-comment").value.trim();
    const email = localStorage.getItem("email");
    if (!text) return;

    try {
        const response = await fetch("http://localhost:8080/newComment", {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                email: email,
                lesson_title: name,   
                content: text,
            })
        });

        if (response.ok) {
            document.getElementById("new-comment").value = "";
            const comments = await response.json();
            displayComments(comments);
            alert("Комментарий успешно добавлен");
        } else {
            const error = await response.text();
            alert("Ошибка: " + error);
        }
    } catch (err) {
        alert('Нет связи с сервером');
    }
}

function displayComments(comments) {
    const commentList = document.getElementById("comment-list");
    
    if (comments.length === 0) {
        commentList.innerHTML = "<p>Пока нет комментариев. Будьте первым!</p>";
        return;
    }

    commentList.innerHTML = comments.map(comment => `
        <div class="comment">
            <span class="date">${new Date(comment.created_at).toLocaleString()}</span>
            <p>${comment.content}</p>
        </div>
    `).join('');
}

async function loadPage() {
    console.log('=== START loadPage ===');
    
    if (Object.keys(currentLessons).length === 0) {
        console.log('Loading lessons data...');
        await loadLessonsData();
        console.log('Lessons loaded:', currentLessons);
    }
    
    const page = location.pathname.split("/").pop().replace(".html", "");
    const hash = location.hash.substring(1);
    
    console.log('Page:', page, 'Hash:', hash);
    console.log('Available pages:', Object.keys(currentLessons));
    
    let topicKey = hash;
    if (!topicKey && currentLessons[page]) {
        topicKey = Object.keys(currentLessons[page])[0];
        console.log('Auto-selected topic:', topicKey);
    }
    
    const topic = currentLessons[page]?.[topicKey];
    console.log('Found topic:', topic);
    
    if (!topic) {
        console.log('Topic not found!');
        document.getElementById("content").innerHTML = `
            <div class="error">
                <h2>Урок не найден</h2>
                <p>Доступные темы в ${page}: ${currentLessons[page] ? Object.keys(currentLessons[page]).join(', ') : 'нет'}</p>
                <p>Debug: page="${page}", topicKey="${topicKey}"</p>
            </div>
        `;
        return;
    }
    
    currentTopic = topic;
    
    document.getElementById("content").innerHTML = `
        <h1>${topic.name}</h1>
        <div class="theory">${topic.theory}</div>
        <pre><code class="language-js">${topic.code}</code></pre>

        <div class="controls">
            <button onclick="runVis('auto')">Авто</button>
            <button onclick="runVis('step')">Шаг</button>
            <button onclick="resetVis()">Сброс</button>
        </div>
        <div id="vis"></div>

        <div class="practice">
            <h3>Практика</h3>
            <p>Напиши свою функцию ниже</p>
            <textarea id="user-code" placeholder="// Напиши свой код здесь..." spellcheck="false"></textarea>
            <button onclick="checkPractice()">Проверить</button>
            <div id="practice-result"></div>
        </div>

        <div class="comments">
            <h3>Комментарии</h3>
            <div id="comment-list"></div>
            <form id="comment-form">
                <textarea id="new-comment" placeholder="Поделись мыслями..." required></textarea>
                <button type="submit">Отправить</button>
            </form>
        </div>
    `;

    document.getElementById("comment-form").addEventListener("submit", (e) => {
        e.preventDefault();
        addComment(topic.name);
    });

    displayComments([]);
    
    hljs.highlightAll();
    markActiveLink(topicKey);
}

function runVis(mode) {
  if (!currentTopic) return;

  if (interval) clearInterval(interval);
  if (currentTopic.visualize) {
    currentTopic.visualize(mode);
  } else {
    d3.select("#vis").html(
      "<p>Визуализация для этого урока пока не доступна</p>"
    );
  }
}

function checkPractice() {
  if (!currentTopic || !currentTopic.test) {
    const resultDiv = document.getElementById("practice-result");
    resultDiv.textContent =
      "Практическое задание для этого урока пока не настроено";
    resultDiv.className = "error";
    return;
  }

  const userCode = document.getElementById("user-code").value.trim();
  const resultDiv = document.getElementById("practice-result");

  if (!userCode) {
    resultDiv.textContent = "Введите код для проверки!";
    resultDiv.className = "error";
    return;
  }

  try {
    const userFunction = new Function(
      "arr",
      userCode +
        "; return bubbleSort ? bubbleSort : quickSort ? quickSort : binarySearch ? binarySearch : null;"
    );
    const testFn = userFunction();

    if (!testFn) {
      resultDiv.textContent =
        "Не найдена функция для тестирования. Убедитесь, что функция имеет правильное имя.";
      resultDiv.className = "error";
      return;
    }

    const isCorrect = currentTopic.test(testFn);

    if (isCorrect) {
      resultDiv.textContent = "✅ Отлично! Алгоритм работает верно.";
      resultDiv.className = "success";
    } else {
      resultDiv.textContent = "❌ Неверно. Проверьте логику алгоритма.";
      resultDiv.className = "error";
    }

    if (currentTopic.lesson_id) {
      localStorage.setItem(`lesson-${currentTopic.lesson_id}-code`, userCode);
    }
  } catch (err) {
    resultDiv.textContent = "Ошибка выполнения: " + err.message;
    resultDiv.className = "error";
  }
}

function formatDate(dateString) {
  const date = new Date(dateString);
  return (
    date.toLocaleDateString("ru-RU") +
    " " +
    date.toLocaleTimeString("ru-RU", {
      hour: "2-digit",
      minute: "2-digit",
    })
  );
}

function showNotification(message, type = "info") {
  const notification = document.createElement("div");
  notification.className = `notification ${type}`;
  notification.textContent = message;

  document.body.appendChild(notification);

  setTimeout(() => {
    notification.remove();
  }, 3000);
}

function resetVis() {
  if (interval) clearInterval(interval);
  d3.select("#vis").html("");
}

function markActiveLink(hash) {
  document.querySelectorAll(".sidebar a").forEach((a) => {
    a.classList.toggle("active", a.getAttribute("href") === `#${hash}`);
  });
}