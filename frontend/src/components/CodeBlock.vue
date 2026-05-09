<script setup lang="ts">
import hljs from "highlight.js/lib/core";
import go from "highlight.js/lib/languages/go";

hljs.registerLanguage("go", go);

const props = defineProps<{
  codeLines: string[];
  activeLines: number[];
}>();

function highlight(line: string) {
  return hljs.highlight(line, {
    language: "go",
  }).value;
}
</script>

<template>
  <div class="code-block">
    <div
      v-for="(line, index) in codeLines"
      :key="index"
      class="code-line"
      :class="{
        active: activeLines.includes(index + 1)
      }"
    >
      <div class="line-number">
        {{ index + 1 }}
      </div>

      <div class="line-code">
        <code v-html="highlight(line)" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.code-block {
  background: #0f172a;
  border-radius: 14px;
  overflow: hidden;
  border: 1px solid rgba(255,255,255,0.06);
  font-family: "JetBrains Mono", "Fira Code", monospace;
  font-size: 13px;
}

.code-line {
  display: flex;
  align-items: center;
  transition: background 0.2s ease, border-left-color 0.2s ease;
  border-left: 3px solid transparent;
}

.code-line:hover {
  background: rgba(255,255,255,0.03);
}

.code-line.active {
  background: rgba(99, 102, 241, 0.15);
  border-left-color: #6366f1;
}

.line-number {
  width: 48px;
  flex-shrink: 0;
  padding-right: 12px;
  text-align: right;
  color: #64748b;
  background: rgba(255,255,255,0.03);
  border-right: 1px solid rgba(255,255,255,0.05);
  user-select: none;
}

.line-code {
  flex: 1;
  display: flex;
  padding-left: 8px;
  color: #e2e8f0;
  white-space: pre;
}

.line-code code {
  display: block;
  line-height: 1;
  background: transparent !important;
}

/* Highlight.js */
:deep(.hljs-keyword) { color: #c084fc; }
:deep(.hljs-string) { color: #86efac; }
:deep(.hljs-number) { color: #fca5a5; }
:deep(.hljs-title) { color: #7dd3fc; }
:deep(.hljs-function) { color: #7dd3fc; }
:deep(.hljs-comment) { color: #64748b; font-style: italic; }
:deep(.hljs-built_in) { color: #facc15; }
</style>