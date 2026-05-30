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
  background: #ffffff; 
  border-radius: 14px; 
  overflow: hidden;
  border: 1px solid #d3d3d3;
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
  background: #f7f7f7;
}

.code-line.active {
    background: #e8f2ff;
    border-left-color: #90caf9;
}

.line-number {
  width: 48px;
  flex-shrink: 0;
  padding-right: 12px;
  text-align: right;
  color: #237893;
  background: #f7f7f7;
  border-right: 1px solid #d3d3d3;
  user-select: none;
}

.line-code {
  flex: 1;
  display: flex;
  padding-left: 8px;
  color: #000000;
  white-space: pre;
}

.line-code code {
  display: block;
  line-height: 1;
  background: transparent !important;
}

/* Highlight.js */
:deep(.hljs-keyword) { color: #7f0055; font-weight: bold; }
:deep(.hljs-string) { color: #2a00ff; }
:deep(.hljs-number) { color: #164; }
:deep(.hljs-title) { color: #0000ff; font-weight: bold;}
:deep(.hljs-function) { color: #0000c0; font-weight: normal;}
:deep(.hljs-comment) { color: #3f7f5f; font-style: italic; }
:deep(.hljs-built_in) { color: #30a; font-weight: bold;}
:deep(.hljs-type) { color: #0000c0; }
:deep(.hljs-literal) { color: #221199; }
:deep(.hljs-operator) { color: #000000; }
:deep(.hljs-variable) { color: #000000; }
</style>
