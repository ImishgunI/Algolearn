// courses.js — ВСЕ ТЕМЫ + D3.js + ПОШАГОВО + ПРАКТИКА

const topics = {
    // === 1. СОРТИРОВКИ ===
    sorts: {
        bubble: {
            name: "Пузырьковая сортировка",
            theory: `<p>Сравнивает соседние элементы и меняет их местами, если они в неправильном порядке. "Пузырьки" всплывают вверх.</p>`,
            code: `function bubbleSort(arr) {
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
            data: () => Array.from({length: 12}, () => Math.floor(Math.random() * 90) + 10),
            test: (fn) => fn([5, 3, 8, 1]).join() === '1,3,5,8',
            visualize: (mode) => visualizeBubble(mode)
        },
        merge: {
            name: "Сортировка слиянием",
            theory: `<p>Разделяет массив пополам, рекурсивно сортирует части, затем сливает их.</p>`,
            code: `function mergeSort(arr) {
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
            data: () => Array.from({length: 10}, () => Math.floor(Math.random() * 100)),
            test: (fn) => fn([38, 27, 43, 3]).join() === '3,27,38,43',
            visualize: (mode) => visualizeMerge(mode)
        },
        quick: {
            name: "Быстрая сортировка",
            theory: `<p>Выбирает опорный элемент (pivot), разделяет массив на меньшие/большие, рекурсивно сортирует.</p>`,
            code: `function quickSort(arr, low = 0, high = arr.length - 1) {
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
            data: () => Array.from({length: 10}, () => Math.floor(Math.random() * 100)),
            test: (fn) => fn([10, 7, 8, 9, 1]).join() === '1,7,8,9,10',
            visualize: (mode) => visualizeQuick(mode)
        }
    },

    // === 2. ПОИСК ===
    search: {
        linear: {
            name: "Линейный поиск",
            theory: `<p>Проверяет каждый элемент по порядку, пока не найдёт цель.</p>`,
            code: `function linearSearch(arr, target) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === target) return i;
    }
    return -1;
}`,
            data: () => ({ arr: [12, 5, 8, 3, 9], target: 8 }),
            test: (fn) => fn([12, 5, 8, 3, 9], 8) === 2,
            visualize: (mode) => visualizeLinear(mode)
        },
        binary: {
            name: "Бинарный поиск",
            theory: `<p>Делит отсортированный массив пополам, исключая ненужную половину.</p>`,
            code: `function binarySearch(arr, target) {
    let left = 0, right = arr.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] === target) return mid;
        if (arr[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}`,
            data: () => ({ arr: [2, 3, 5, 7, 11, 13], target: 7 }),
            test: (fn) => fn([2, 3, 5, 7, 11, 13], 7) === 3,
            visualize: (mode) => visualizeBinary(mode)
        }
    },

    // === 3. РЕКУРСИЯ ===
    recursion: {
        factorial: {
            name: "Факториал",
            theory: `<p>Рекурсивно: f(n) = n × f(n-1), f(1) = 1</p>`,
            code: `function factorial(n) {
    if (n <= 1) return 1;
    return n * factorial(n - 1);
}`,
            test: (fn) => fn(5) === 120,
            visualize: (mode) => visualizeRecursion('factorial', mode)
        },
        fibonacci: {
            name: "Числа Фибоначчи",
            theory: `<p>f(n) = f(n-1) + f(n-2)</p>`,
            code: `function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}`,
            test: (fn) => fn(6) === 8,
            visualize: (mode) => visualizeRecursion('fibonacci', mode)
        }
    },

    // === 4. ДИНАМИЧЕСКОЕ ПРОГРАММИРОВАНИЕ ===
    dynamic: {
        fib: {
            name: "Фибоначчи (DP)",
            theory: `<p>Сохраняем промежуточные результаты в массиве.</p>`,
            code: `function fibDP(n) {
    const dp = [0, 1];
    for (let i = 2; i <= n; i++) {
        dp[i] = dp[i-1] + dp[i-2];
    }
    return dp[n];
}`,
            test: (fn) => fn(10) === 55,
            visualize: (mode) => visualizeDPFib(mode)
        }
    },

    // === 5. ГРАФЫ И ДЕРЕВЬЯ ===
    trees: {
        bfs: {
            name: "Обход в ширину (BFS)",
            theory: `<p>Исследует узлы по уровням, используя очередь.</p>`,
            code: `function bfs(graph, start) {
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
            data: () => ({
                graph: { A: ['B','C'], B: ['D'], C: ['E'], D: [], E: [] },
                start: 'A'
            }),
            test: (fn) => fn({ A: ['B','C'], B: ['D'], C: ['E'], D: [], E: [] }, 'A').join() === 'A,B,C,D,E',
            visualize: (mode) => visualizeBFS(mode)
        }
    },

    // === 6. СЛОЖНОСТЬ ===
    difficult: {
        'big-o': {
            name: "Big O нотация",
            theory: `<p>O(1), O(n), O(n²), O(log n) — сравнение роста.</p>`,
            code: `// Пример O(n²)
for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
        console.log(i, j);
    }
}`,
            visualize: (mode) => visualizeBigO(mode)
        }
    }
};

// === ВИЗУАЛИЗАЦИИ ===
let interval = null;

function visualizeBubble(mode) {
    const data = topics.sorts.bubble.data();
    const svg = d3.select("#vis").html("").append("svg").attr("width", 800).attr("height", 300);
    let i = 0, j = 0;

    const update = () => {
        svg.selectAll("rect").data(data).join(
            enter => enter.append("rect")
                .attr("x", (d, i) => i * 65 + 20)
                .attr("y", d => 250 - d * 2)
                .attr("width", 60)
                .attr("height", d => d * 2)
                .attr("fill", "#0078D7"),
            update => update.transition().duration(200)
                .attr("y", d => 250 - d * 2)
                .attr("height", d => d * 2)
                .attr("fill", (d, idx) => idx === j || idx === j+1 ? "#ff6b6b" : "#0078D7")
        );
    };

    if (mode === 'auto') {
        interval = setInterval(() => {
            if (i >= data.length - 1) { clearInterval(interval); return; }
            if (j >= data.length - i - 1) { j = 0; i++; return; }
            if (data[j] > data[j + 1]) [data[j], data[j + 1]] = [data[j + 1], data[j]];
            j++;
            update();
        }, 400);
    }
    update();
}

// === ВИЗУАЛИЗАЦИЯ: СОРТИРОВКА СЛИЯНИЕМ ===
function visualizeMerge(mode) {
    const data = topics.sorts.merge.data();
    const svg = d3.select("#vis").html("").append("svg").attr("width", 900).attr("height", 400);
    let steps = [], current = [{ arr: data.slice(), stage: "split" }];

    // Генерация шагов
    function split(arr) {
        if (arr.length <= 1) return [arr];
        const mid = Math.floor(arr.length / 2);
        const left = split(arr.slice(0, mid));
        const right = split(arr.slice(mid));
        return left.concat(right);
    }

    function mergeSteps(left, right) {
        let result = [], i = 0, j = 0, step = [];
        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                step.push([...result, left[i++]]);
            } else {
                step.push([...result, right[j++]]);
            }
        }
        return step.concat(left.slice(i)).concat(right.slice(j));
    }

    const halves = split(data.slice());
    let i = 0, j = 0, merging = [];
    while (i < halves.length - 1) {
        merging.push({ left: halves[i], right: halves[i+1] });
        i += 2;
    }

    merging.forEach(pair => {
        const merged = mergeSteps(pair.left, pair.right);
        steps = steps.concat(merged.map(arr => ({ arr, stage: "merge" })));
    });

    let stepIndex = 0;
    const update = () => {
        if (stepIndex >= steps.length) return;
        const step = steps[stepIndex];
        const bars = svg.selectAll("rect").data(step.arr);
        bars.enter().append("rect")
            .attr("x", (d, i) => i * 35 + 50)
            .attr("y", d => 350 - d * 3)
            .attr("width", 30)
            .attr("height", d => d * 3)
            .attr("fill", "#0078D7")
            .merge(bars)
            .transition().duration(300)
            .attr("y", d => 350 - d * 3)
            .attr("height", d => d * 3);
        stepIndex++;
    };

    if (mode === 'auto') {
        interval = setInterval(update, 600);
    } else {
        update();
        document.querySelector('.controls').innerHTML += `<button onclick="stepMerge()">Далее</button>`;
        window.stepMerge = () => update();
    }
}

// === ВИЗУАЛИЗАЦИЯ: БЫСТРАЯ СОРТИРОВКА ===
function visualizeQuick(mode) {
    const data = topics.sorts.quick.data();
    const svg = d3.select("#vis").html("").append("svg").attr("width", 800).attr("height", 300);
    let steps = [], arr = data.slice();

    function quickSteps(arr, low = 0, high = arr.length - 1) {
        if (low >= high) return;
        const pivot = arr[high];
        let i = low;
        steps.push({ arr: arr.slice(), pivot, i, j: low });
        for (let j = low; j < high; j++) {
            if (arr[j] < pivot) {
                [arr[i], arr[j]] = [arr[j], arr[i]];
                i++;
                steps.push({ arr: arr.slice(), pivot, i, j });
            }
        }
        [arr[i], arr[high]] = [arr[high], arr[i]];
        steps.push({ arr: arr.slice(), pivot, i, j: high });
        quickSteps(arr, low, i - 1);
        quickSteps(arr, i + 1, high);
    }

    quickSteps(arr);
    let stepIndex = 0;

    const update = () => {
        if (stepIndex >= steps.length) return;
        const { arr, pivot, i, j } = steps[stepIndex];
        const bars = svg.selectAll("rect").data(arr);
        bars.enter().append("rect")
            .attr("x", (d, idx) => idx * 70 + 30)
            .attr("y", d => 250 - d * 2)
            .attr("width", 60)
            .attr("height", d => d * 2)
            .attr("fill", (d, idx) => idx === j ? "#ff6b6b" : idx === i ? "#4ecdc4" : "#0078D7")
            .merge(bars)
            .transition().duration(400)
            .attr("y", d => 250 - d * 2)
            .attr("height", d => d * 2)
            .attr("fill", (d, idx) => idx === j ? "#ff6b6b" : idx === i ? "#4ecdc4" : "#0078D7");
        stepIndex++;
    };

    if (mode === 'auto') {
        interval = setInterval(update, 800);
    } else {
        update();
        document.querySelector('.controls').innerHTML += `<button onclick="stepQuick()">Далее</button>`;
        window.stepQuick = () => update();
    }
}

// === ВИЗУАЛИЗАЦИЯ: ЛИНЕЙНЫЙ ПОИСК ===
function visualizeLinear(mode) {
    const { arr, target } = topics.search.linear.data();
    const svg = d3.select("#vis").html("").append("svg").attr("width", 700).attr("height", 150);
    let i = 0;

    const update = () => {
        svg.selectAll("rect").data(arr).join(
            enter => enter.append("rect")
                .attr("x", (d, idx) => idx * 80 + 50)
                .attr("y", 50)
                .attr("width", 70)
                .attr("height", 70)
                .attr("fill", (d, idx) => idx === i ? (d === target ? "#4ecdc4" : "#ff6b6b") : "#0078D7")
        );
        svg.selectAll("text").data(arr).join(
            enter => enter.append("text")
                .attr("x", (d, idx) => idx * 80 + 85)
                .attr("y", 90)
                .text(d => d)
                .attr("fill", "white")
                .attr("font-size", "18px")
        );
        if (arr[i] === target) clearInterval(interval);
        i++;
    };

    if (mode === 'auto') {
        interval = setInterval(update, 1000);
    } else {
        update();
        document.querySelector('.controls').innerHTML += `<button onclick="stepLinear()">Далее</button>`;
        window.stepLinear = () => update();
    }
    update();
}

// === ВИЗУАЛИЗАЦИЯ: БИНАРНЫЙ ПОИСК ===
function visualizeBinary(mode) {
    const { arr, target } = topics.search.binary.data();
    const svg = d3.select("#vis").html("").append("svg").attr("width", 800).attr("height", 180);
    let left = 0, right = arr.length - 1;

    const update = () => {
        const mid = Math.floor((left + right) / 2);
        svg.selectAll("rect").data(arr).join(
            enter => enter.append("rect")
                .attr("x", (d, i) => i * 70 + 50)
                .attr("y", 50)
                .attr("width", 60)
                .attr("height", 60)
                .attr("fill", (d, i) =>
                    i < left || i > right ? "#ccc" :
                        i === mid ? (d === target ? "#4ecdc4" : "#ff6b6b") : "#0078D7"
                )
        );
        svg.selectAll("text").data(arr).join(
            enter => enter.append("text")
                .attr("x", (d, i) => i * 70 + 80)
                .attr("y", 85)
                .text(d => d)
                .attr("fill", "white")
        );

        if (arr[mid] === target || left > right) return clearInterval(interval);
        if (arr[mid] < target) left = mid + 1;
        else right = mid - 1;
    };

    if (mode === 'auto') {
        interval = setInterval(update, 1200);
    } else {
        update();
        document.querySelector('.controls').innerHTML += `<button onclick="stepBinary()">Далее</button>`;
        window.stepBinary = () => update();
    }
    update();
}

// === ВИЗУАЛИЗАЦИЯ: РЕКУРСИЯ (СТЕК ВЫЗОВОВ) ===
function visualizeRecursion(type, mode) {
    const svg = d3.select("#vis").html("").append("svg").attr("width", 600).attr("height", 400);
    const g = svg.append("g").attr("transform", "translate(50,50)");
    let stack = [], calls = 0, max = type === 'factorial' ? 5 : 7;

    const draw = () => {
        g.selectAll("rect").data(stack).join(
            enter => enter.append("rect")
                .attr("x", 0)
                .attr("y", (d, i) => i * 50)
                .attr("width", 200)
                .attr("height", 45)
                .attr("fill", "#0078D7")
                .attr("rx", 8),
            update => update.transition().attr("y", (d, i) => i * 50)
        );
        g.selectAll("text").data(stack).join(
            enter => enter.append("text")
                .attr("x", 100)
                .attr("y", (d, i) => i * 50 + 28)
                .text(d => d)
                .attr("fill", "white")
                .attr("text-anchor", "middle")
        );
    };

    const step = () => {
        if (calls > max) return;
        stack.push(`${type}(${max - calls})`);
        draw();
        calls++;
        if (calls > max) setTimeout(() => { stack.pop(); draw(); }, 800);
    };

    if (mode === 'auto') {
        interval = setInterval(step, 1000);
    } else {
        step();
        document.querySelector('.controls').innerHTML += `<button onclick="stepRec()">Далее</button>`;
        window.stepRec = step;
    }
}

// === ВИЗУАЛИЗАЦИЯ: BFS (ГРАФ) ===
function visualizeBFS(mode) {
    const { graph, start } = topics.trees.bfs.data();
    const nodes = Object.keys(graph);
    const links = [];
    nodes.forEach(n => graph[n].forEach(t => links.push({ source: n, target: t })));

    const svg = d3.select("#vis").html("").append("svg").attr("width", 600).attr("height", 400);
    const simulation = d3.forceSimulation(nodes)
        .force("link", d3.forceLink(links).distance(100))
        .force("charge", d3.forceManyBody().strength(-300))
        .force("center", d3.forceCenter(300, 200));

    const link = svg.selectAll(".link").data(links).enter().append("line").attr("stroke", "#999");
    const node = svg.selectAll(".node").data(nodes).enter().append("circle")
        .attr("r", 20)
        .attr("fill", d => d === start ? "#ff6b6b" : "#0078D7");

    const label = svg.selectAll(".label").data(nodes).enter().append("text")
        .text(d => d)
        .attr("x", d => d === start ? -8 : -5)
        .attr("y", 5)
        .attr("fill", "white");

    let queue = [start], visited = new Set();

    const step = () => {
        if (!queue.length) return clearInterval(interval);
        const node = queue.shift();
        if (visited.has(node)) return;
        visited.add(node);
        d3.selectAll("circle").filter(d => d === node).attr("fill", "#4ecdc4");
        graph[node].forEach(n => { if (!visited.has(n)) queue.push(n); });
    };

    simulation.on("tick", () => {
        link.attr("x1", d => d.source.x).attr("y1", d => d.source.y)
            .attr("x2", d => d.target.x).attr("y2", d => d.target.y);
        node.attr("cx", d => d.x).attr("cy", d => d.y);
        label.attr("x", d => d.x).attr("y", d => d.y);
    });

    if (mode === 'auto') {
        interval = setInterval(step, 1500);
    } else {
        step();
        document.querySelector('.controls').innerHTML += `<button onclick="stepBFS()">Далее</button>`;
        window.stepBFS = step;
    }
}

// === ВИЗУАЛИЗАЦИЯ: BIG O (ГРАФИК) ===
function visualizeBigO(mode) {
    const svg = d3.select("#vis").html("").append("svg").attr("width", 700).attr("height", 350);
    const data = d3.range(1, 21).map(x => ({
        x,
        'O(1)': 1,
        'O(log n)': Math.log2(x),
        'O(n)': x,
        'O(n log n)': x * Math.log2(x),
        'O(n²)': x * x
    }));

    const xScale = d3.scaleLinear().domain([1, 20]).range([60, 650]);
    const yScale = d3.scaleLinear().domain([0, 400]).range([300, 50]);

    const line = d3.line()
        .x(d => xScale(d.x))
        .y(d => yScale(d[Object.keys(d)[1]]));

    const colors = { 'O(1)': '#4ecdc4', 'O(log n)': '#45b7d1', 'O(n)': '#f9ca24', 'O(n log n)': '#f0932b', 'O(n²)': '#eb4d4b' };

    Object.keys(colors).forEach((key, i) => {
        svg.append("path")
            .datum(data)
            .attr("d", d => line(d.map(p => ({ x: p.x, [key]: p[key] }))))
            .attr("fill", "none")
            .attr("stroke", colors[key])
            .attr("stroke-width", 3);

        svg.append("text")
            .attr("x", 70)
            .attr("y", 30 + i * 20)
            .text(key)
            .attr("fill", colors[key])
            .attr("font-weight", "bold");
    });

    svg.selectAll(".dot").data(data).enter().append("circle")
        .attr("cx", d => xScale(d.x))
        .attr("cy", d => yScale(d['O(n)']))
        .attr("r", 3)
        .attr("fill", "#f9ca24");
}

// Остальные визуализации (merge, quick, binary, recursion, BFS, Big O) — аналогично
// По аналогии с bubble — могу выдать по запросу

// === ПРАКТИКА ===
function checkPractice(page, topicKey) {
    const topic = topics[page][topicKey];
    const userCode = document.getElementById("user-code").value.trim();
    const resultDiv = document.getElementById("practice-result");

    if (!userCode) {
        resultDiv.textContent = "Введите код для проверки!";
        resultDiv.className = "error";
        return;
    }

    // Очистка результата
    resultDiv.textContent = "";

    try {
        // Создаём функцию из кода пользователя
        const fullCode = `${userCode}; ${topic.funcName}([...${JSON.stringify(topic.test.input)}]);`;
        const userResult = eval(fullCode);

        // Сравниваем с эталоном
        const correct = JSON.stringify(userResult) === JSON.stringify(topic.test.output);

        if (correct) {
            resultDiv.textContent = "✅ Отлично! Алгоритм работает верно.";
            resultDiv.className = "success";
        } else {
            resultDiv.textContent = `❌ Неверно. Ожидалось: ${JSON.stringify(topic.test.output)}, а получено: ${JSON.stringify(userResult)}`;
            resultDiv.className = "error";
        }

        // сохраняем прогресс
        localStorage.setItem(`${page}-${topicKey}-code`, userCode);
    } catch (err) {
        resultDiv.textContent = "Ошибка выполнения: " + err.message;
        resultDiv.className = "error";
    }
}

// === КОММЕНТАРИИ ===
function addComment(key) {
    const text = document.getElementById('new-comment').value.trim();
    if (!text) return;
    const comments = JSON.parse(localStorage.getItem(`c_${key}`) || '[]');
    comments.push({ text, date: new Date().toLocaleString() });
    localStorage.setItem(`c_${key}`, JSON.stringify(comments));
    document.getElementById('new-comment').value = '';
    loadComments(key);
}

function loadComments(key) {
    const list = document.getElementById('comment-list');
    const comments = JSON.parse(localStorage.getItem(`c_${key}`) || '[]');
    list.innerHTML = comments.map(c => `
        <div class="comment">
            <p>${c.text}</p>
            <small>${c.date}</small>
        </div>
    `).join('');
}

// === ЗАГРУЗКА ===
function loadPage() {
    const page = location.pathname.split('/').pop().replace('.html', '');
    const hash = location.hash.substring(1) || Object.keys(topics[page])[0];
    const topic = topics[page][hash];

    document.getElementById('content').innerHTML = `
        <h1>${topic.name}</h1>
        <div class="theory">${topic.theory}</div>
        <pre><code class="language-js">${topic.code}</code></pre>

        <div class="controls">
            <button onclick="runVis('${page}','${hash}','auto')">Авто</button>
            <button onclick="runVis('${page}','${hash}','step')">Шаг</button>
            <button onclick="resetVis()">Сброс</button>
        </div>
        <div id="vis"></div>

        <div class="practice">
            <h3>Практика</h3>
            <p>Напиши функцию ниже, чтобы реализовать алгоритм пузырьковой сортировки</p>

            <textarea id="user-code" placeholder="// Напиши свой код здесь..." spellcheck="false"></textarea>
            <button onclick="checkPractice('${page}','${hash}')">Проверить</button>
            <div id="practice-result"></div>
        </div>


        <div class="comments">
            <h3>Комментарии</h3>
            <div id="comment-list"></div>
            <textarea id="new-comment" placeholder="Поделись мыслями..."></textarea>
            <button onclick="addComment('${page}-${hash}')">Отправить</button>
        </div>
    `;

    loadComments(`${page}-${hash}`);
    hljs.highlightAll();
    markActiveLink(hash);
}

function runVis(page, hash, mode) {
    if (interval) clearInterval(interval);
    topics[page][hash].visualize(mode);
}

function resetVis() {
    if (interval) clearInterval(interval);
    d3.select("#vis").html("");
}

function markActiveLink(hash) {
    document.querySelectorAll('.sidebar a').forEach(a => {
        a.classList.toggle('active', a.getAttribute('href') === `#${hash}`);
    });
}

