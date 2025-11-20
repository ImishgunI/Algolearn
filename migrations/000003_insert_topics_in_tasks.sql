-- Задачи для каждого урока
INSERT INTO tasks (lesson_id, title, description, difficulty, answer) VALUES
-- Пузырьковая сортировка
(
    (SELECT id FROM lessons WHERE title = 'Пузырьковая сортировка'),
    'Реализация пузырьковой сортировки',
    'Напишите функцию bubbleSort, которая принимает массив чисел и возвращает отсортированный массив по возрастанию.',
    'beginner',
    'function bubbleSort(arr) {\n    const n = arr.length;\n    for (let i = 0; i < n - 1; i++) {\n        for (let j = 0; j < n - i - 1; j++) {\n            if (arr[j] > arr[j + 1]) {\n                [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];\n            }\n        }\n    }\n    return arr;\n}'
),

-- Сортировка слиянием
(
    (SELECT id FROM lessons WHERE title = 'Сортировка слиянием'),
    'Реализация сортировки слиянием',
    'Напишите функцию mergeSort, которая рекурсивно сортирует массив используя алгоритм слияния.',
    'intermediate',
    'function mergeSort(arr) {\n    if (arr.length <= 1) return arr;\n    const mid = Math.floor(arr.length / 2);\n    const left = mergeSort(arr.slice(0, mid));\n    const right = mergeSort(arr.slice(mid));\n    return merge(left, right);\n}\n\nfunction merge(left, right) {\n    let result = [], i = 0, j = 0;\n    while (i < left.length && j < right.length) {\n        if (left[i] <= right[j]) result.push(left[i++]);\n        else result.push(right[j++]);\n    }\n    return result.concat(left.slice(i)).concat(right.slice(j));\n}'
),

-- Быстрая сортировка
(
    (SELECT id FROM lessons WHERE title = 'Быстрая сортировка'),
    'Реализация быстрой сортировки',
    'Напишите функцию quickSort с использованием partition для разделения массива.',
    'intermediate',
    'function quickSort(arr, low = 0, high = arr.length - 1) {\n    if (low >= high) return arr;\n    const pivotIdx = partition(arr, low, high);\n    quickSort(arr, low, pivotIdx - 1);\n    quickSort(arr, pivotIdx + 1, high);\n    return arr;\n}\n\nfunction partition(arr, low, high) {\n    const pivot = arr[high];\n    let i = low;\n    for (let j = low; j < high; j++) {\n        if (arr[j] < pivot) {\n            [arr[i], arr[j]] = [arr[j], arr[i]];\n            i++;\n        }\n    }\n    [arr[i], arr[high]] = [arr[high], arr[i]];\n    return i;\n}'
),

-- Линейный поиск
(
    (SELECT id FROM lessons WHERE title = 'Линейный поиск'),
    'Реализация линейного поиска',
    'Напишите функцию linearSearch, которая находит индекс target в массиве или возвращает -1.',
    'beginner',
    'function linearSearch(arr, target) {\n    for (let i = 0; i < arr.length; i++) {\n        if (arr[i] === target) return i;\n    }\n    return -1;\n}'
),

-- Бинарный поиск
(
    (SELECT id FROM lessons WHERE title = 'Бинарный поиск'),
    'Реализация бинарного поиска',
    'Напишите функцию binarySearch для поиска в отсортированном массиве.',
    'beginner',
    'function binarySearch(arr, target) {\n    let left = 0, right = arr.length - 1;\n    while (left <= right) {\n        const mid = Math.floor((left + right) / 2);\n        if (arr[mid] === target) return mid;\n        if (arr[mid] < target) left = mid + 1;\n        else right = mid - 1;\n    }\n    return -1;\n}'
),

-- Факториал
(
    (SELECT id FROM lessons WHERE title = 'Факториал'),
    'Рекурсивный факториал',
    'Напишите рекурсивную функцию factorial для вычисления факториала числа.',
    'beginner',
    'function factorial(n) {\n    if (n <= 1) return 1;\n    return n * factorial(n - 1);\n}'
),

-- Числа Фибоначчи
(
    (SELECT id FROM lessons WHERE title = 'Числа Фибоначчи'),
    'Рекурсивные числа Фибоначчи',
    'Напишите рекурсивную функцию fibonacci для вычисления n-го числа Фибоначчи.',
    'beginner',
    'function fibonacci(n) {\n    if (n <= 1) return n;\n    return fibonacci(n - 1) + fibonacci(n - 2);\n}'
),

-- Фибоначчи (DP)
(
    (SELECT id FROM lessons WHERE title = 'Фибоначчи (DP)'),
    'Фибоначчи с динамическим программированием',
    'Реализуйте вычисление чисел Фибоначчи с использованием динамического программирования.',
    'intermediate',
    'function fibDP(n) {\n    const dp = [0, 1];\n    for (let i = 2; i <= n; i++) {\n        dp[i] = dp[i-1] + dp[i-2];\n    }\n    return dp[n];\n}'
),

-- BFS
(
    (SELECT id FROM lessons WHERE title = 'Обход в ширину (BFS)'),
    'Реализация BFS для графа',
    'Напишите функцию bfs для обхода графа в ширину.',
    'intermediate',
    'function bfs(graph, start) {\n    const visited = new Set();\n    const queue = [start];\n    const result = [];\n    while (queue.length) {\n        const node = queue.shift();\n        if (!visited.has(node)) {\n            visited.add(node);\n            result.push(node);\n            queue.push(...graph[node]);\n        }\n    }\n    return result;\n}'
);