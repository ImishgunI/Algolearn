let currentLessons = {};
let currentTopic = null;
let interval = null;

const topicMapping = {
  sorts: {
    bubble: "Пузырьковая сортировка",
    merge: "Сортировка слиянием",
    quick: "Быстрая сортировка",
  },
  search: {
    linear: "Линейный поиск",
    binary: "Бинарный поиск",
  },
  recursion: {
    factorial: "Факториал",
    fibonacci: "Числа Фибоначчи",
  },
  dynamic: {
    fib: "Фибоначчи (DP)",
  },
  trees: {
    bfs: "Обход в ширину (BFS)",
  },
  difficult: {
    "big-o": "Big O нотация",
  },
};

async function loadLessonsData() {
  try {
    const response = await fetch("http://localhost:8080/lessons");
    if (response.ok) {
      const data = await response.json();
      const lessonsArray = data.lessons || data;
      console.log("Loaded lessons:", lessonsArray);

      currentLessons = transformLessonsData(lessonsArray);
      console.log("Transformed lessons:", currentLessons);
      return currentLessons;
    } else {
      console.error("Ошибка загрузки уроков");
      return {};
    }
  } catch (err) {
    console.error("Нет связи с сервером:", err);
    return {};
  }
}

function transformLessonsData(lessons) {
  const topics = {};

  for (const [category, lessonsMap] of Object.entries(topicMapping)) {
    topics[category] = {};
    for (const [key, title] of Object.entries(lessonsMap)) {
      topics[category][key] = {
        name: title,
        theory: "",
        code: getDefaultCode(title),
        data: getDefaultData(title),
        test: getDefaultTest(title),
        visualize: getVisualizationFunction(title),
      };
    }
  }

  lessons.forEach((lesson) => {
    const category = lesson.category;

    if (topics[category]) {
      for (const [key, topic] of Object.entries(topics[category])) {
        if (topic.name === lesson.title) {
          topic.theory = lesson.content;
          topic.lesson_id = lesson.id;
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
    "Пузырьковая сортировка": `function bubbleSort(arr) {
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
    "Сортировка слиянием": `function mergeSort(arr) {
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
    "Быстрая сортировка": `function quickSort(arr, low = 0, high = arr.length - 1) {
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
    "Линейный поиск": `function linearSearch(arr, target) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === target) return i;
    }
    return -1;
}`,
    "Бинарный поиск": `function binarySearch(arr, target) {
    let left = 0, right = arr.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] === target) return mid;
        if (arr[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}`,
    Факториал: `function factorial(n) {
    if (n <= 1) return 1;
    return n * factorial(n - 1);
}`,
    "Числа Фибоначчи": `function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}`,
    "Фибоначчи (DP)": `function fibDP(n) {
    const dp = [0, 1];
    for (let i = 2; i <= n; i++) {
        dp[i] = dp[i-1] + dp[i-2];
    }
    return dp[n];
}`,
    "Обход в ширину (BFS)": `function bfs(graph, start) {
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
    "Big O нотация": `// Пример O(n²)
for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
        console.log(i, j);
    }
}`,
  };

  return codeMap[lessonTitle] || "// Код для этого урока будет добавлен позже";
}

function getDefaultData(lessonTitle) {
  const dataMap = {
    "Пузырьковая сортировка": () =>
      Array.from({ length: 12 }, () => Math.floor(Math.random() * 90) + 10),
    "Сортировка слиянием": () =>
      Array.from({ length: 10 }, () => Math.floor(Math.random() * 100)),
    "Быстрая сортировка": () =>
      Array.from({ length: 10 }, () => Math.floor(Math.random() * 100)),
    "Линейный поиск": () => ({ arr: [12, 5, 8, 3, 9], target: 8 }),
    "Бинарный поиск": () => ({ arr: [2, 3, 5, 7, 11, 13], target: 7 }),
    Факториал: () => 5,
    "Числа Фибоначчи": () => 6,
    "Фибоначчи (DP)": () => 10,
    "Обход в ширину (BFS)": () => ({
      graph: { A: ["B", "C"], B: ["D"], C: ["E"], D: [], E: [] },
      start: "A",
    }),
    "Big O нотация": () => ({}),
  };

  return dataMap[lessonTitle] || (() => []);
}

function getDefaultTest(lessonTitle) {
  const testMap = {
    "Пузырьковая сортировка": (fn) => fn([5, 3, 8, 1]).join() === "1,3,5,8",
    "Сортировка слиянием": (fn) => fn([38, 27, 43, 3]).join() === "3,27,38,43",
    "Быстрая сортировка": (fn) => fn([10, 7, 8, 9, 1]).join() === "1,7,8,9,10",
    "Линейный поиск": (fn) => fn([12, 5, 8, 3, 9], 8) === 2,
    "Бинарный поиск": (fn) => fn([2, 3, 5, 7, 11, 13], 7) === 3,
    Факториал: (fn) => fn(5) === 120,
    "Числа Фибоначчи": (fn) => fn(6) === 8,
    "Фибоначчи (DP)": (fn) => fn(10) === 55,
    "Обход в ширину (BFS)": (fn) =>
      fn({ A: ["B", "C"], B: ["D"], C: ["E"], D: [], E: [] }, "A").join() ===
      "A,B,C,D,E",
    "Big O нотация": (fn) => true,
  };

  return testMap[lessonTitle] || (() => true);
}

function getVisualizationFunction(lessonTitle) {
  const visMap = {
    "Пузырьковая сортировка": (mode) => visualizeBubble(mode),
    "Сортировка слиянием": (mode) => visualizeMerge(mode),
    "Быстрая сортировка": (mode) => visualizeQuick(mode),
    "Линейный поиск": (mode) => visualizeLinear(mode),
    "Бинарный поиск": (mode) => visualizeBinary(mode),
    Факториал: (mode) => visualizeRecursion("factorial", mode),
    "Числа Фибоначчи": (mode) => visualizeRecursion("fibonacci", mode),
    "Фибоначчи (DP)": (mode) => visualizeDPFib(mode),
    "Обход в ширину (BFS)": (mode) => visualizeBFS(mode),
    "Big O нотация": (mode) => visualizeBigO(mode),
  };

  return (
    visMap[lessonTitle] ||
    (() => {
      d3.select("#vis").html(
        "<p>Визуализация для этого урока пока не доступна</p>"
      );
    })
  );
}

// === КОММЕНТАРИИ ===
async function addComment(name) {
  const text = document.getElementById("new-comment").value.trim();
  const email = localStorage.getItem("email");
  if (!text) return;

  try {
    const response = await fetch("http://localhost:8080/newComment", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: email,
        lesson_title: name,
        content: text,
      }),
    });

    if (response.ok) {
      document.getElementById("new-comment").value = "";
      loadComments(name);
      alert("Комментарий успешно добавлен");
    } else {
      const error = await response.text();
      alert("Ошибка: " + error);
    }
  } catch (err) {
    alert("Нет связи с сервером");
  }
}

function formatDateReadable(dateString) {
  const date = new Date(dateString);
  return date.toLocaleString("ru-RU", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
}

function displayComments(comments) {
  const commentList = document.getElementById("comment-list");

  if (comments.length === 0) {
    commentList.innerHTML = "<p>Пока нет комментариев. Будьте первым!</p>";
    return;
  }
  console.log(comments.created_at);
  commentList.innerHTML = comments
    .map(
      (comment) => `
        <div class="comment">
            <strong class="comment-author">${comment.first_name} ${
        comment.last_name
      }</strong>
            <span class="date">${formatDateReadable(comment.created_at)}</span>
            <p>${comment.content}</p>
        </div>
    `
    )
    .join("");
}

async function loadComments(lessonTitle) {
  try {
    const response = await fetch(
      `http://localhost:8080/comments?lesson_title=${encodeURIComponent(
        lessonTitle
      )}`
    );
    if (response.ok) {
      const data = await response.json();
      displayComments(data.comments);
    } else {
      console.error("Ошибка загрузки комментариев");
    }
  } catch (err) {
    console.error("Ошибка сети при загрузке комментариев", err);
  }
}

async function loadPage() {
  console.log("=== START loadPage ===");

  if (Object.keys(currentLessons).length === 0) {
    console.log("Loading lessons data...");
    await loadLessonsData();
    console.log("Lessons loaded:", currentLessons);
  }

  const page = location.pathname.split("/").pop().replace(".html", "");
  const hash = location.hash.substring(1);

  console.log("Page:", page, "Hash:", hash);
  console.log("Available pages:", Object.keys(currentLessons));

  let topicKey = hash;
  if (!topicKey && currentLessons[page]) {
    topicKey = Object.keys(currentLessons[page])[0];
    console.log("Auto-selected topic:", topicKey);
  }

  const topic = currentLessons[page]?.[topicKey];
  console.log("Found topic:", topic);

  if (!topic) {
    console.log("Topic not found!");
    document.getElementById("content").innerHTML = `
            <div class="error">
                <h2>Урок не найден</h2>
                <p>Доступные темы в ${page}: ${
      currentLessons[page]
        ? Object.keys(currentLessons[page]).join(", ")
        : "нет"
    }</p>
                <p>Debug: page="${page}", topicKey="${topicKey}"</p>
            </div>
        `;
    return;
  }

  currentTopic = topic;

  document.getElementById("content").innerHTML = `
        <h1>${topic.name}</h1>
        <button class="favorite-btn" onclick="toggleFavorite(
            '${topic.name}'
        )">⭐ В избранное
        </button>
        <button id="done-lesson-btn" class="done-lesson" onclick="setDoneLesson('${topic.name}')">Отметить пройденным</button>

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
            <form id="comment-form">
                <textarea id="new-comment" placeholder="Поделись мыслями..." required></textarea>
                <button type="submit">Отправить</button>
            </form>
            <div id="comment-list"></div>
        </div>
    `;

  document.getElementById("comment-form").addEventListener("submit", (e) => {
    e.preventDefault();
    addComment(topic.name);
  });

  await loadComments(topic.name);

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

document.addEventListener("DOMContentLoaded", function () {
  console.log("DOM loaded, initializing...");
  loadPage();
});

window.addEventListener("hashchange", function () {
  console.log("Hash changed, reloading page...");
  loadPage();
});

async function setDoneLesson(lessonTitle) {
    const email = localStorage.getItem("email");

    if (!email) {
        alert("Войдите в аккаунт, чтобы отмечать уроки пройденными.");
        return;
    }

    const btn = document.getElementById("done-lesson-btn");

    btn.disabled = true;
    btn.textContent = "Сохраняю...";

    try {
        const response = await fetch(`http://localhost:8080/lessons/complete?email=${email}`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                lesson_title: lessonTitle
            })
        });

        if (response.ok) {
            btn.textContent = "Урок пройден";
            btn.classList.add("lesson-done");
        } else {
            btn.textContent = "Ошибка! Попробовать снова";
            btn.disabled = false;
        }
    } catch (err) {
        console.error("Ошибка сети:", err);
        btn.textContent = "Ошибка сети!";
        btn.disabled = false;
    }
}

document.getElementById("done-lesson").addEventListener()

// === ВИЗУАЛИЗАЦИИ ===
function visualizeBubble(mode) {
  if (!currentTopic) return;

  const data = currentTopic.data();
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 800)
    .attr("height", 300);

  let i = 0,
    j = 0;

  const update = () => {
    svg
      .selectAll("rect")
      .data(data)
      .join(
        (enter) =>
          enter
            .append("rect")
            .attr("x", (d, i) => i * 65 + 20)
            .attr("y", (d) => 250 - d * 2)
            .attr("width", 60)
            .attr("height", (d) => d * 2)
            .attr("fill", "#0078D7"),
        (update) =>
          update
            .transition()
            .duration(200)
            .attr("y", (d) => 250 - d * 2)
            .attr("height", (d) => d * 2)
            .attr("fill", (d, idx) =>
              idx === j || idx === j + 1 ? "#ff6b6b" : "#0078D7"
            )
      );
  };

  if (mode === "auto") {
    interval = setInterval(() => {
      if (i >= data.length - 1) {
        clearInterval(interval);
        return;
      }
      if (j >= data.length - i - 1) {
        j = 0;
        i++;
        return;
      }
      if (data[j] > data[j + 1])
        [data[j], data[j + 1]] = [data[j + 1], data[j]];
      j++;
      update();
    }, 400);
  }
  update();
}

function visualizeQuick(mode) {
  if (!currentTopic) return;

  const data = currentTopic.data();
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 800)
    .attr("height", 300);

  let steps = [];
  let arr = [...data];

  function quickSteps(arr, low = 0, high = arr.length - 1) {
    if (low >= high) return;
    const pivot = arr[high];
    let i = low;
    steps.push({ arr: [...arr], pivot, i, j: low });
    for (let j = low; j < high; j++) {
      if (arr[j] < pivot) {
        [arr[i], arr[j]] = [arr[j], arr[i]];
        i++;
        steps.push({ arr: [...arr], pivot, i, j });
      }
    }
    [arr[i], arr[high]] = [arr[high], arr[i]];
    steps.push({ arr: [...arr], pivot, i, j: high });
    quickSteps(arr, low, i - 1);
    quickSteps(arr, i + 1, high);
  }

  quickSteps(arr);
  let stepIndex = 0;

  const update = () => {
    if (stepIndex >= steps.length) return;
    const { arr, pivot, i, j } = steps[stepIndex];
    const bars = svg.selectAll("rect").data(arr);
    bars
      .enter()
      .append("rect")
      .attr("x", (d, idx) => idx * 70 + 30)
      .attr("y", (d) => 250 - d * 2)
      .attr("width", 60)
      .attr("height", (d) => d * 2)
      .attr("fill", (d, idx) =>
        idx === j ? "#ff6b6b" : idx === i ? "#4ecdc4" : "#0078D7"
      )
      .merge(bars)
      .transition()
      .duration(400)
      .attr("y", (d) => 250 - d * 2)
      .attr("height", (d) => d * 2)
      .attr("fill", (d, idx) =>
        idx === j ? "#ff6b6b" : idx === i ? "#4ecdc4" : "#0078D7"
      );
    stepIndex++;
  };

  if (mode === "auto") {
    interval = setInterval(update, 800);
  } else {
    update();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepQuick()">Далее</button>`;
    window.stepQuick = () => update();
  }
}

function visualizeLinear(mode) {
  if (!currentTopic) return;

  const data = currentTopic.data();
  const { arr, target } = data;
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 700)
    .attr("height", 150);

  let i = 0;

  const update = () => {
    svg
      .selectAll("rect")
      .data(arr)
      .join((enter) =>
        enter
          .append("rect")
          .attr("x", (d, idx) => idx * 80 + 50)
          .attr("y", 50)
          .attr("width", 70)
          .attr("height", 70)
          .attr("fill", (d, idx) =>
            idx === i ? (d === target ? "#4ecdc4" : "#ff6b6b") : "#0078D7"
          )
      );
    svg
      .selectAll("text")
      .data(arr)
      .join((enter) =>
        enter
          .append("text")
          .attr("x", (d, idx) => idx * 80 + 85)
          .attr("y", 90)
          .text((d) => d)
          .attr("fill", "white")
          .attr("font-size", "18px")
      );
    if (arr[i] === target) clearInterval(interval);
    i++;
  };

  if (mode === "auto") {
    interval = setInterval(update, 1000);
  } else {
    update();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepLinear()">Далее</button>`;
    window.stepLinear = () => update();
  }
  update();
}

function visualizeBinary(mode) {
  if (!currentTopic) return;

  const data = currentTopic.data();
  const { arr, target } = data;
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 800)
    .attr("height", 180);

  let left = 0,
    right = arr.length - 1;

  const update = () => {
    const mid = Math.floor((left + right) / 2);
    svg
      .selectAll("rect")
      .data(arr)
      .join((enter) =>
        enter
          .append("rect")
          .attr("x", (d, i) => i * 70 + 50)
          .attr("y", 50)
          .attr("width", 60)
          .attr("height", 60)
          .attr("fill", (d, i) =>
            i < left || i > right
              ? "#ccc"
              : i === mid
              ? d === target
                ? "#4ecdc4"
                : "#ff6b6b"
              : "#0078D7"
          )
      );
    svg
      .selectAll("text")
      .data(arr)
      .join((enter) =>
        enter
          .append("text")
          .attr("x", (d, i) => i * 70 + 80)
          .attr("y", 85)
          .text((d) => d)
          .attr("fill", "white")
      );

    if (arr[mid] === target || left > right) return clearInterval(interval);
    if (arr[mid] < target) left = mid + 1;
    else right = mid - 1;
  };

  if (mode === "auto") {
    interval = setInterval(update, 1200);
  } else {
    update();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepBinary()">Далее</button>`;
    window.stepBinary = () => update();
  }
  update();
}

function visualizeBFS(mode) {
  if (!currentTopic) return;

  const data = currentTopic.data();
  const { graph, start } = data;
  const nodes = Object.keys(graph);
  const links = [];
  nodes.forEach((n) =>
    graph[n].forEach((t) => links.push({ source: n, target: t }))
  );

  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 600)
    .attr("height", 400);

  const simulation = d3
    .forceSimulation(nodes)
    .force("link", d3.forceLink(links).distance(100))
    .force("charge", d3.forceManyBody().strength(-300))
    .force("center", d3.forceCenter(300, 200));

  const link = svg
    .selectAll(".link")
    .data(links)
    .enter()
    .append("line")
    .attr("stroke", "#999");

  const node = svg
    .selectAll(".node")
    .data(nodes)
    .enter()
    .append("circle")
    .attr("r", 20)
    .attr("fill", (d) => (d === start ? "#ff6b6b" : "#0078D7"));

  const label = svg
    .selectAll(".label")
    .data(nodes)
    .enter()
    .append("text")
    .text((d) => d)
    .attr("x", (d) => (d === start ? -8 : -5))
    .attr("y", 5)
    .attr("fill", "white");

  let queue = [start],
    visited = new Set();

  const step = () => {
    if (!queue.length) return clearInterval(interval);
    const node = queue.shift();
    if (visited.has(node)) return;
    visited.add(node);
    d3.selectAll("circle")
      .filter((d) => d === node)
      .attr("fill", "#4ecdc4");
    graph[node].forEach((n) => {
      if (!visited.has(n)) queue.push(n);
    });
  };

  simulation.on("tick", () => {
    link
      .attr("x1", (d) => d.source.x)
      .attr("y1", (d) => d.source.y)
      .attr("x2", (d) => d.target.x)
      .attr("y2", (d) => d.target.y);
    node.attr("cx", (d) => d.x).attr("cy", (d) => d.y);
    label.attr("x", (d) => d.x).attr("y", (d) => d.y);
  });

  if (mode === "auto") {
    interval = setInterval(step, 1500);
  } else {
    step();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepBFS()">Далее</button>`;
    window.stepBFS = step;
  }
}

function visualizeBigO(mode) {
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 700)
    .attr("height", 350);

  const data = d3.range(1, 21).map((x) => ({
    x,
    "O(1)": 1,
    "O(log n)": Math.log2(x),
    "O(n)": x,
    "O(n log n)": x * Math.log2(x),
    "O(n²)": x * x,
  }));

  const xScale = d3.scaleLinear().domain([1, 20]).range([60, 650]);
  const yScale = d3.scaleLinear().domain([0, 400]).range([300, 50]);

  const line = d3
    .line()
    .x((d) => xScale(d.x))
    .y((d) => yScale(d[Object.keys(d)[1]]));

  const colors = {
    "O(1)": "#4ecdc4",
    "O(log n)": "#45b7d1",
    "O(n)": "#f9ca24",
    "O(n log n)": "#f0932b",
    "O(n²)": "#eb4d4b",
  };

  Object.keys(colors).forEach((key, i) => {
    svg
      .append("path")
      .datum(data)
      .attr("d", (d) => line(d.map((p) => ({ x: p.x, [key]: p[key] }))))
      .attr("fill", "none")
      .attr("stroke", colors[key])
      .attr("stroke-width", 3);

    svg
      .append("text")
      .attr("x", 70)
      .attr("y", 30 + i * 20)
      .text(key)
      .attr("fill", colors[key])
      .attr("font-weight", "bold");
  });

  svg
    .selectAll(".dot")
    .data(data)
    .enter()
    .append("circle")
    .attr("cx", (d) => xScale(d.x))
    .attr("cy", (d) => yScale(d["O(n)"]))
    .attr("r", 3)
    .attr("fill", "#f9ca24");
}

// === ВИЗУАЛИЗАЦИЯ: СОРТИРОВКА СЛИЯНИЕМ ===
function visualizeMerge(mode) {
  if (!currentTopic) return;

  const data = currentTopic.data();
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 900)
    .attr("height", 400);

  let steps = [];
  let currentStep = 0;

  // Генерация шагов сортировки слиянием
  function generateMergeSteps(arr) {
    const steps = [];

    function mergeSortSteps(arr, depth = 0, position = 0) {
      if (arr.length <= 1) {
        steps.push({
          arr: [...arr],
          depth: depth,
          position: position,
          action: "base",
          left: null,
          right: null,
        });
        return [arr];
      }

      const mid = Math.floor(arr.length / 2);
      const left = arr.slice(0, mid);
      const right = arr.slice(mid);

      // Шаг разделения
      steps.push({
        arr: [...arr],
        depth: depth,
        position: position,
        action: "split",
        left: [...left],
        right: [...right],
      });

      const leftSorted = mergeSortSteps(left, depth + 1, position);
      const rightSorted = mergeSortSteps(right, depth + 1, position + mid);

      // Шаг слияния
      const merged = merge(leftSorted[0], rightSorted[0]);
      steps.push({
        arr: [...merged],
        depth: depth,
        position: position,
        action: "merge",
        left: [...leftSorted[0]],
        right: [...rightSorted[0]],
      });

      return [merged];
    }

    function merge(left, right) {
      let result = [];
      let i = 0,
        j = 0;

      while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
          result.push(left[i]);
          i++;
        } else {
          result.push(right[j]);
          j++;
        }
      }

      return result.concat(left.slice(i)).concat(right.slice(j));
    }

    mergeSortSteps([...arr]);
    return steps;
  }

  steps = generateMergeSteps(data);

  const update = () => {
    if (currentStep >= steps.length) {
      clearInterval(interval);
      return;
    }

    const step = steps[currentStep];
    svg.selectAll("*").remove();

    // Отрисовка текущего состояния
    const barWidth = 25;
    const baseY = 50;

    step.arr.forEach((value, index) => {
      const x = index * (barWidth + 5) + 50;
      const y = baseY + step.depth * 60;
      const height = value * 3;

      svg
        .append("rect")
        .attr("x", x)
        .attr("y", y)
        .attr("width", barWidth)
        .attr("height", height)
        .attr(
          "fill",
          step.action === "merge"
            ? "#4ecdc4"
            : step.action === "split"
            ? "#ff6b6b"
            : "#0078D7"
        );

      svg
        .append("text")
        .attr("x", x + barWidth / 2)
        .attr("y", y + height + 15)
        .text(value)
        .attr("text-anchor", "middle")
        .attr("fill", "black");
    });

    // Подпись действия
    svg
      .append("text")
      .attr("x", 50)
      .attr("y", 30)
      .text(`Шаг ${currentStep + 1}: ${getActionText(step.action)}`)
      .attr("fill", "black")
      .attr("font-weight", "bold");

    currentStep++;
  };

  function getActionText(action) {
    switch (action) {
      case "split":
        return "Разделение";
      case "merge":
        return "Слияние";
      case "base":
        return "Базовый случай";
      default:
        return "";
    }
  }

  if (mode === "auto") {
    interval = setInterval(update, 1000);
  } else {
    update();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepMerge()">Далее</button>`;
    window.stepMerge = () => update();
  }
}

// === ВИЗУАЛИЗАЦИЯ: РЕКУРСИЯ (ФАКТОРИАЛ И ФИБОНАЧЧИ) ===
function visualizeRecursion(type, mode) {
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 800)
    .attr("height", 500);

  const stack = [];
  const maxDepth = type === "factorial" ? 5 : 6;
  let currentStep = 0;
  let callStack = [];

  // Генерация шагов рекурсии
  function generateRecursionSteps() {
    const steps = [];

    function generateFactorialSteps(n, depth = 0) {
      steps.push({
        type: "call",
        function: "factorial",
        argument: n,
        depth: depth,
        result: null,
      });

      if (n <= 1) {
        steps.push({
          type: "return",
          function: "factorial",
          argument: n,
          depth: depth,
          result: 1,
        });
        return 1;
      }

      const recursiveResult = generateFactorialSteps(n - 1, depth + 1);
      const result = n * recursiveResult;

      steps.push({
        type: "return",
        function: "factorial",
        argument: n,
        depth: depth,
        result: result,
      });

      return result;
    }

    function generateFibonacciSteps(n, depth = 0) {
      steps.push({
        type: "call",
        function: "fibonacci",
        argument: n,
        depth: depth,
        result: null,
      });

      if (n <= 1) {
        steps.push({
          type: "return",
          function: "fibonacci",
          argument: n,
          depth: depth,
          result: n,
        });
        return n;
      }

      const fib1 = generateFibonacciSteps(n - 1, depth + 1);
      const fib2 = generateFibonacciSteps(n - 2, depth + 1);
      const result = fib1 + fib2;

      steps.push({
        type: "return",
        function: "fibonacci",
        argument: n,
        depth: depth,
        result: result,
      });

      return result;
    }

    if (type === "factorial") {
      generateFactorialSteps(maxDepth);
    } else {
      generateFibonacciSteps(maxDepth);
    }

    return steps;
  }

  const steps = generateRecursionSteps();

  const update = () => {
    if (currentStep >= steps.length) {
      clearInterval(interval);
      return;
    }

    const step = steps[currentStep];
    svg.selectAll("*").remove();

    // Обновляем стек вызовов
    if (step.type === "call") {
      callStack.push(step);
    } else if (step.type === "return") {
      callStack = callStack.filter((call) => call.depth < step.depth);
    }

    // Отрисовка стека вызовов
    callStack.forEach((call, index) => {
      const y = 100 + index * 60;

      svg
        .append("rect")
        .attr("x", 100)
        .attr("y", y)
        .attr("width", 200)
        .attr("height", 50)
        .attr("fill", "#0078D7")
        .attr("rx", 5);

      svg
        .append("text")
        .attr("x", 200)
        .attr("y", y + 30)
        .text(`${call.function}(${call.argument})`)
        .attr("text-anchor", "middle")
        .attr("fill", "white")
        .attr("font-weight", "bold");
    });

    // Отрисовка текущего шага
    const currentY = 400;
    svg
      .append("text")
      .attr("x", 400)
      .attr("y", currentY - 20)
      .text("Текущая операция:")
      .attr("text-anchor", "middle")
      .attr("fill", "black");

    svg
      .append("rect")
      .attr("x", 300)
      .attr("y", currentY)
      .attr("width", 200)
      .attr("height", 50)
      .attr("fill", step.type === "call" ? "#4ecdc4" : "#ff6b6b")
      .attr("rx", 5);

    svg
      .append("text")
      .attr("x", 400)
      .attr("y", currentY + 30)
      .text(
        `${step.type === "call" ? "Вызов" : "Возврат"} ${step.function}(${
          step.argument
        })${step.result ? ` = ${step.result}` : ""}`
      )
      .attr("text-anchor", "middle")
      .attr("fill", "white")
      .attr("font-weight", "bold");

    // Информация о стеке
    svg
      .append("text")
      .attr("x", 50)
      .attr("y", 50)
      .text(`Глубина стека: ${callStack.length}`)
      .attr("fill", "black")
      .attr("font-weight", "bold");

    currentStep++;
  };

  if (mode === "auto") {
    interval = setInterval(update, 1500);
  } else {
    update();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepRecursion()">Далее</button>`;
    window.stepRecursion = () => update();
  }
}

// === ВИЗУАЛИЗАЦИЯ: ФИБОНАЧЧИ С ДИНАМИЧЕСКИМ ПРОГРАММИРОВАНИЕМ ===
function visualizeDPFib(mode) {
  const svg = d3
    .select("#vis")
    .html("")
    .append("svg")
    .attr("width", 800)
    .attr("height", 400);

  const n = 10;
  const dp = [0, 1];
  let currentStep = 0;
  const steps = [];

  // Генерация шагов ДП
  for (let i = 2; i <= n; i++) {
    steps.push({
      index: i,
      dp: [...dp],
      calculation: `dp[${i}] = dp[${i - 1}] + dp[${i - 2}] = ${dp[i - 1]} + ${
        dp[i - 2]
      } = ${dp[i - 1] + dp[i - 2]}`,
      result: dp[i - 1] + dp[i - 2],
    });
    dp.push(dp[i - 1] + dp[i - 2]);
  }

  const update = () => {
    if (currentStep >= steps.length) {
      clearInterval(interval);
      return;
    }

    const step = steps[currentStep];
    svg.selectAll("*").remove();

    // Отрисовка массива dp
    const cellSize = 60;
    const startX = 100;
    const startY = 100;

    // Заголовок
    svg
      .append("text")
      .attr("x", 400)
      .attr("y", 50)
      .text(`Вычисление Fibonacci(${step.index})`)
      .attr("text-anchor", "middle")
      .attr("fill", "black")
      .attr("font-weight", "bold")
      .attr("font-size", "20px");

    // Отрисовка ячеек массива
    step.dp.forEach((value, index) => {
      const x = startX + index * cellSize;
      const isCurrent = index === step.index;
      const isUsed = index === step.index - 1 || index === step.index - 2;

      // Ячейка
      svg
        .append("rect")
        .attr("x", x)
        .attr("y", startY)
        .attr("width", cellSize)
        .attr("height", cellSize)
        .attr("fill", isCurrent ? "#ff6b6b" : isUsed ? "#4ecdc4" : "#0078D7")
        .attr("stroke", "black");

      // Индекс
      svg
        .append("text")
        .attr("x", x + cellSize / 2)
        .attr("y", startY - 10)
        .text(`dp[${index}]`)
        .attr("text-anchor", "middle")
        .attr("fill", "black");

      // Значение
      svg
        .append("text")
        .attr("x", x + cellSize / 2)
        .attr("y", startY + cellSize / 2 + 5)
        .text(value)
        .attr("text-anchor", "middle")
        .attr("fill", "white")
        .attr("font-weight", "bold")
        .attr("font-size", "18px");
    });

    // Формула расчета
    svg
      .append("text")
      .attr("x", 400)
      .attr("y", startY + 100)
      .text(step.calculation)
      .attr("text-anchor", "middle")
      .attr("fill", "black")
      .attr("font-size", "16px");

    // Результат
    svg
      .append("text")
      .attr("x", 400)
      .attr("y", startY + 130)
      .text(`Fibonacci(${step.index}) = ${step.result}`)
      .attr("text-anchor", "middle")
      .attr("fill", "#ff6b6b")
      .attr("font-weight", "bold")
      .attr("font-size", "18px");

    // Визуализация зависимости
    if (step.index >= 2) {
      const arrowStartX = startX + (step.index - 1) * cellSize + cellSize / 2;
      const arrowEndX = startX + step.index * cellSize - 10;

      // Стрелка от dp[i-1]
      svg
        .append("line")
        .attr("x1", arrowStartX)
        .attr("y1", startY + cellSize)
        .attr("x2", arrowEndX)
        .attr("y2", startY + cellSize + 40)
        .attr("stroke", "#4ecdc4")
        .attr("stroke-width", 2);

      svg
        .append("polygon")
        .attr(
          "points",
          `${arrowEndX},${startY + cellSize + 40} ${arrowEndX - 5},${
            startY + cellSize + 35
          } ${arrowEndX + 5},${startY + cellSize + 35}`
        )
        .attr("fill", "#4ecdc4");

      // Стрелка от dp[i-2]
      const arrowStartX2 = startX + (step.index - 2) * cellSize + cellSize / 2;
      svg
        .append("line")
        .attr("x1", arrowStartX2)
        .attr("y1", startY + cellSize)
        .attr("x2", arrowEndX)
        .attr("y2", startY + cellSize + 40)
        .attr("stroke", "#4ecdc4")
        .attr("stroke-width", 2);

      svg
        .append("polygon")
        .attr(
          "points",
          `${arrowEndX},${startY + cellSize + 40} ${arrowEndX - 5},${
            startY + cellSize + 35
          } ${arrowEndX + 5},${startY + cellSize + 35}`
        )
        .attr("fill", "#4ecdc4");

      // Подписи стрелок
      svg
        .append("text")
        .attr("x", (arrowStartX + arrowEndX) / 2)
        .attr("y", startY + cellSize + 60)
        .text("+")
        .attr("text-anchor", "middle")
        .attr("fill", "#4ecdc4")
        .attr("font-weight", "bold");
    }

    currentStep++;
  };

  if (mode === "auto") {
    interval = setInterval(update, 1200);
  } else {
    update();
    document.querySelector(
      ".controls"
    ).innerHTML += `<button onclick="stepDPFib()">Далее</button>`;
    window.stepDPFib = () => update();
  }
}
